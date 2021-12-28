package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cocus_challenger_refact/app/cocus/mhttp"
	"github.com/cocus_challenger_refact/app/cocus/terrors"
	core "github.com/cocus_challenger_refact/business/core/login"
	"github.com/cocus_challenger_refact/business/data/login"
)

type LoginHandlers struct {
	Log          *log.Logger
	LoginManager core.LoginInt
}

func (h LoginHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)
	h.Log.Printf("receive request to create a new user")

	var req login.NewLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.Log.Printf("Error to decode from json, err %s", err)
		terrors.Handler(w, 500,
			fmt.Errorf("Error to decode from json, err:%s", err.Error()))
		return
	}

	err = req.Validate()
	if err != nil {
		h.Log.Printf("Error to validate fields from login request, err %s", err)
		terrors.Handler(w, 400, err)
		return
	}

	login := req.GenerateEntity()
	err = h.LoginManager.CreateUser(login)
	if err != nil {
		h.Log.Printf("Error to create a new user, err %s", err)
		terrors.Handler(w, 500, err)
		return
	}

	if err := mhttp.WriteJsonResponse(w, nil, http.StatusCreated); err != nil {
		terrors.Handler(w, http.StatusCreated, err)
		return
	}

}

func (h LoginHandlers) Login(w http.ResponseWriter, r *http.Request) {
	h.Log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)
	h.Log.Printf("receive request to login into system")

	var req login.NewLogin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.Log.Printf("Error to decode from json, err %s", err)
		terrors.Handler(w, http.StatusInternalServerError,
			fmt.Errorf("Error to decode from json, err:%s", err.Error()))
		return
	}

	err = req.Validate()
	if err != nil {
		h.Log.Printf("Error to validate fields from login, err %s", err)
		terrors.Handler(w, http.StatusBadRequest, err)
		return
	}

	loginEntity := req.GenerateEntity()
	token, err := h.LoginManager.Login(loginEntity)
	if err != nil {
		h.Log.Printf("Error to login into system, err %s", err)
		terrors.Handler(w, http.StatusUnauthorized, err)
		return
	}

	if err := mhttp.WriteJsonResponse(w, token, http.StatusOK); err != nil {
		terrors.Handler(w, http.StatusOK, err)
		return
	}
}
