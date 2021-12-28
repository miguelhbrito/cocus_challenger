package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cocus_challenger_refact/app/cocus/terrors"
	"github.com/cocus_challenger_refact/business/data/login"
	"github.com/golang-jwt/jwt"
)

type Middleware struct {
	Log *log.Logger
}

func NewMiddleware(Log *log.Logger) Middleware {
	return Middleware{
		Log: Log,
	}
}

func (m Middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Log.Printf("Authorization middleware")

		if r.URL.Path != "/login" && r.URL.Path != "/login/create" {
			m.Log.Printf("Authorization middleware checking token auth")

			tokenAuth := r.Header.Get("authorization")
			token, err := jwt.Parse(tokenAuth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return login.JwtKey, nil
			})
			if err != nil {
				m.Log.Printf("Error on get token from request, err: %v", err)
				terrors.Handler(w, 500, err)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				username := claims["username"]
				usernameString := fmt.Sprintf("%s", username)
				m.Log.Printf("Authorization middleware ok, username %s", usernameString)
				next.ServeHTTP(w, r)
			} else {
				m.Log.Printf("Error on decode token, err: %v", err)
				terrors.Handler(w, 401, err)
				return
			}

		} else {
			next.ServeHTTP(w, r)
		}
	})
}
