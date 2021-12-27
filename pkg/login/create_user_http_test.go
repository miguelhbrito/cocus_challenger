package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/request"
	"github.com/cocus_challenger/pkg/terrors"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHTTP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager Login
		h       CreateUserHTTP
		body    []byte
		request request.LoginRequest
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: LoginCustomMock{
				CreateUserMock: func(l entity.LoginEntity) error {
					return nil
				},
			},
			request: request.LoginRequest{
				Username: "username",
				Password: "hashedPassword",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				_ = json.NewEncoder(w).Encode(nil)
			},
		},
		{
			name: "Error to decode json",
			manager: LoginCustomMock{
				CreateUserMock: func(l entity.LoginEntity) error {
					return nil
				},
			},
			body: []byte(""),
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(terrors.HTTPError{Msg: "Error to decode from json, err:EOF"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to validate username and password",
			manager: LoginCustomMock{
				CreateUserMock: func(l entity.LoginEntity) error {
					return nil
				},
			},
			request: request.LoginRequest{
				Username: "",
				Password: "",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(terrors.HTTPError{Msg: "username is required,password is required"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error on create a new user",
			manager: LoginCustomMock{
				CreateUserMock: func(l entity.LoginEntity) error {
					return errors.New("some error")
				},
			},
			request: request.LoginRequest{
				Username: "username",
				Password: "hashedPassword",
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(terrors.HTTPError{Msg: "some error"})
				_, _ = w.Write(data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewCreateUserHTTP(tt.manager)

			body, _ := json.Marshal(tt.request)
			if tt.body != nil {
				body = tt.body
			}
			req, _ := http.NewRequest(http.MethodPost, "/login/create", bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
