package triangle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/request"
	"github.com/cocus_challenger/pkg/api/response"
	"github.com/cocus_challenger/pkg/terrors"
	"github.com/stretchr/testify/assert"
)

func TestCreateTriangleHTTP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager Triangle
		h       CreateTriangleHTTP
		body    []byte
		request request.Triangle
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: TriangleCustomMock{
				CreateMock: func(t entity.Triangle) (entity.Triangle, error) {
					return entity.Triangle{
						Id:    "1",
						Side1: 10,
						Side2: 10,
						Side3: 10,
						Type:  "equilateral",
					}, nil
				},
			},
			request: request.Triangle{
				Side1: 10,
				Side2: 10,
				Side3: 10,
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				_ = json.NewEncoder(w).Encode(response.Triangle{
					Id:    "1",
					Side1: 10,
					Side2: 10,
					Side3: 10,
					Type:  "equilateral",
				})
			},
		},
		{
			name: "Error on get type and save triangle",
			manager: TriangleCustomMock{
				CreateMock: func(t entity.Triangle) (entity.Triangle, error) {
					return entity.Triangle{}, errNotATriangle
				},
			},
			request: request.Triangle{
				Side1: 10,
				Side2: 5,
				Side3: 15,
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(terrors.HTTPError{Msg: errNotATriangle.Error()})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to validate sides",
			manager: TriangleCustomMock{
				CreateMock: func(t entity.Triangle) (entity.Triangle, error) {
					return entity.Triangle{}, nil
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				data, _ := json.Marshal(terrors.HTTPError{Msg: "side1 can't be lower than 0 or equal 0,side2 can't be lower than 0 or equal 0,side3 can't be lower than 0 or equal 0"})
				_, _ = w.Write(data)
			},
		},
		{
			name: "Error to decode json",
			manager: TriangleCustomMock{
				CreateMock: func(t entity.Triangle) (entity.Triangle, error) {
					return entity.Triangle{}, nil
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewCreateTriangleHTTP(tt.manager)

			body, _ := json.Marshal(tt.request)
			if tt.body != nil {
				body = tt.body
			}
			req, _ := http.NewRequest(http.MethodPost, "/triangles", bytes.NewReader(body))

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, req)

			g := httptest.NewRecorder()

			h.Handler()(g, req)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")

		})
	}
}
