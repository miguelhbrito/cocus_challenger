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
	LoginHTTP struct {
		loginManager Login
	}
)

func NewLoginHTTP(
	loginManager Login,
) mhttp.HttpHandler {
	return LoginHTTP{
		loginManager: loginManager,
	}
}

func (h LoginHTTP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("receive request to login into system")

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
			log.Error().Err(err).Msg("Error to validate fields from login")
			terrors.Handler(w, 400, err)
			return
		}

		loginEntity := req.GenerateEntity()
		token, err := h.loginManager.Login(loginEntity)
		if err != nil {
			log.Error().Err(err).Msg("Error to login into system")
			terrors.Handler(w, 401, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, token, http.StatusOK); err != nil {
			terrors.Handler(w, http.StatusOK, err)
			return
		}
	}
}
