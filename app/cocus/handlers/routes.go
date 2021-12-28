package handlers

import (
	"log"
	"net/http"

	loginHandlers "github.com/cocus_challenger_refact/app/cocus/handlers/login"
	triangleHandlers "github.com/cocus_challenger_refact/app/cocus/handlers/triangle"
	"github.com/cocus_challenger_refact/business/auth"
	loginCore "github.com/cocus_challenger_refact/business/core/login"
	triangleCore "github.com/cocus_challenger_refact/business/core/triangle"
	"github.com/cocus_challenger_refact/business/data/login"
	"github.com/cocus_challenger_refact/business/data/triangle"
	"github.com/cocus_challenger_refact/business/middleware"

	"github.com/gorilla/mux"
)

//NewMux with all routes
func NewMux(log *log.Logger) http.Handler {
	router := mux.NewRouter()

	//Auth manager
	authManager := auth.Authorization{}

	//Starting triangle postgres and manager
	trianglePostgres := triangle.TrianglePostgres{}
	triangleCore := triangleCore.NewCore(trianglePostgres)

	//Starting login postgres and manager
	loginPostgres := login.LoginPostgres{}
	loginCore := loginCore.NewCore(loginPostgres, authManager)

	//HANDLERS ----------------

	//Starting triangle handlers
	handlersTriangle := triangleHandlers.TriangleHandlers{
		TriangleManager: triangleCore,
	}

	//Starting login handlers
	handlersLogin := loginHandlers.LoginHandlers{
		LoginManager: loginCore,
	}

	//ENDPOINTS----------------

	//Login endpoints
	router.HandleFunc("/login/create", handlersLogin.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlersLogin.Login).Methods("POST")

	//Triangle endpoints
	router.HandleFunc("/triangles", handlersTriangle.Create).Methods("POST")
	router.HandleFunc("/triangles", handlersTriangle.List).Methods("GET")

	//Middleware
	router.Use(middleware.Authorization)

	return router
}
