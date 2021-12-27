package triangle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cocus_challenger/pkg/api/entity"
	"github.com/cocus_challenger/pkg/api/response"
	"github.com/cocus_challenger/pkg/terrors"
	"github.com/stretchr/testify/assert"
)

func TestListTrianglesHTTP_Handler(t *testing.T) {
	tests := []struct {
		name    string
		manager Triangle
		h       ListTrianglesHTTP
		want    http.HandlerFunc
	}{
		{
			name: "Success",
			manager: TriangleCustomMock{
				ListMock: func() (entity.Triangles, error) {
					return entity.Triangles{{Id: "1"}}, nil
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode([]response.Triangle{{
					Id: "1",
				}})
			},
		},
		{
			name: "Error to list all triangles from system",
			manager: TriangleCustomMock{
				ListMock: func() (entity.Triangles, error) {
					return nil, errors.New("some error")
				},
			},
			want: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				data, _ := json.Marshal(terrors.HTTPError{Msg: errors.New("some error").Error()})
				_, _ = w.Write(data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewListTrianglesHTTP(tt.manager)

			r, _ := http.NewRequest(http.MethodGet, "/triangles", nil)

			w := httptest.NewRecorder()

			tt.want.ServeHTTP(w, r)

			g := httptest.NewRecorder()

			h.Handler()(g, r)

			assert.Equal(t, w.Code, g.Result().StatusCode, fmt.Sprintf("expected status code %v ", w.Code))

			assert.Equal(t, w.Body.String(), g.Body.String(), "body was not equal as expected")
		})
	}
}
