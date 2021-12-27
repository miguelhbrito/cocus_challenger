package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cocus_challenger/pkg/api/request"
	"github.com/cocus_challenger/pkg/mhttp"
	"github.com/cocus_challenger/pkg/terrors"
	"github.com/rs/zerolog/log"
)

type (
	CreateUserHTTP struct {
		loginManager Login
	}
)

func NewCreateUserHTTP(
	loginManager Login,
) mhttp.HttpHandler {
	return CreateUserHTTP{
		loginManager: loginManager,
	}
}

func (h CreateUserHTTP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("receive request to create a new user")

		var req request.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error().Err(err).Msg("Error to decode from json")
			terrors.Handler(w, 500,
				fmt.Errorf("Error to decode from json, err:%s", err.Error()))
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error().Err(err).Msg("Error to validate fields from login request")
			terrors.Handler(w, 400, err)
			return
		}

		login := req.GenerateEntity()
		err = h.loginManager.CreateUser(login)
		if err != nil {
			log.Error().Err(err).Msg("Error to create a new user")
			terrors.Handler(w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, "", http.StatusCreated); err != nil {
			terrors.Handler(w, http.StatusCreated, err)
			return
		}
	}
}
