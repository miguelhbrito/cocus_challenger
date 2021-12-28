package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/cocus_challenger_refact/app/cocus/handlers"
	"github.com/cocus_challenger_refact/platform/config"
	db "github.com/cocus_challenger_refact/platform/db_connect"
	"github.com/cocus_challenger_refact/platform/migrations"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {

	// ==========================================================================================
	// Logging
	log := log.New(os.Stdout, "COCUS : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// ==========================================================================================
	// Configuration

	config, err := config.LoadConfig("./platform")
	if err != nil {
		log.Printf("cannot load config:", err)
	}

	var cfg struct {
		Web struct {
			APIHost         string        `conf:"default:127.0.0.1:5000"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}

	if err := conf.Parse(os.Args[1:], "COCUS", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("COCUS", &cfg)
			if err != nil {
				return errors.Wrap(err, "generation config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	// ==========================================================================================
	// Start DB connection
	dbconnection := db.InitDB(config.DBDriver, config.DBSource)

	// Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	// ==========================================================================================
	// Start API Service
	log.Printf("main : Started : Application initalizing")
	defer log.Println("main : Completed")

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generation config for output")
	}
	log.Printf("main : Config :\n%v\n", out)

	cfg.Web.APIHost = config.ServerAddress
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      handlers.NewMux(log, dbconnection),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// ==========================================================================================
	// Shutdown

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Printf("main : %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", cfg.Web.ShutdownTimeout, err)
			err = api.Close()
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
