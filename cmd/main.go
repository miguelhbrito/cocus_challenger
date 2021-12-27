package main

import (
	"fmt"
	"net/http"

	db "github.com/cocus_challenger/db_connect"
	"github.com/cocus_challenger/migrations"
	"github.com/cocus_challenger/pkg/auth"
	"github.com/cocus_challenger/pkg/login"
	"github.com/cocus_challenger/pkg/middleware"
	"github.com/cocus_challenger/pkg/storage"
	"github.com/cocus_challenger/pkg/triangle"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {

	//Starting db connection
	dbconnection := db.InitDB()
	//Starting migrations
	migrations.InitMigrations(dbconnection)
	defer dbconnection.Close()

	//Starting rounter
	router := mux.NewRouter()
	addrHttp := fmt.Sprintf(":%d", 5000)

	//Auth manager
	authManager := auth.NewManager()

	//Starting triangle postgres and manager
	trianglePostgres := storage.NewTrianglePostgres()
	triangleManager := triangle.NewManager(trianglePostgres)

	//Starting login postgres and manager
	loginPostgres := storage.NewLoginPostgres()
	loginManager := login.NewManager(loginPostgres, authManager)

	//HANDLERS ----------------

	//Starting triangle handlers
	triangleCreate := triangle.NewCreateTriangleHTTP(triangleManager)
	triangleList := triangle.NewListTrianglesHTTP(triangleManager)

	//Starting login handlers
	loginCreateUser := login.NewCreateUserHTTP(loginManager)
	loginSystem := login.NewLoginHTTP(loginManager)

	//ENDPOINTS----------------

	//Login endpoints
	router.HandleFunc("/login/create", middleware.Authorization(loginCreateUser.Handler())).Methods("POST")
	router.HandleFunc("/login", middleware.Authorization(loginSystem.Handler())).Methods("POST")

	//Triangle endpoints
	router.HandleFunc("/triangles", middleware.Authorization(triangleCreate.Handler())).Methods("POST")
	router.HandleFunc("/triangles", middleware.Authorization(triangleList.Handler())).Methods("GET")

	//Starting gateway
	log.Fatal().Err(http.ListenAndServe(addrHttp, router)).Msg("failed to start gateway")
}
