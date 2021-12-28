package middleware

import (
	"fmt"
	"net/http"

	"github.com/cocus_challenger_refact/app/cocus/terrors"
	"github.com/cocus_challenger_refact/business/data/login"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msgf("Authorization middleware")

		if r.URL.Path != "/login" && r.URL.Path != "/login/create" {
			log.Debug().Msgf("Authorization middleware checking token auth")

			tokenAuth := r.Header.Get("authorization")
			token, err := jwt.Parse(tokenAuth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return login.JwtKey, nil
			})
			if err != nil {
				log.Error().Msgf("Error on get token from request, err: %v", err)
				terrors.Handler(w, 500, err)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				username := claims["username"]
				usernameString := fmt.Sprintf("%s", username)
				log.Debug().Msgf("Authorization middleware ok, username %s", usernameString)
				next.ServeHTTP(w, r)
			} else {
				log.Error().Msgf("Error on decode token, err: %v", err)
				terrors.Handler(w, 401, err)
				return
			}

		} else {
			next.ServeHTTP(w, r)
		}
	})
}
