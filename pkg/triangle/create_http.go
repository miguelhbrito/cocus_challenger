package triangle

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
	CreateTriangleHTTP struct {
		triangleManager Triangle
	}
)

func NewCreateTriangleHTTP(
	triangleManager Triangle,
) mhttp.HttpHandler {
	return CreateTriangleHTTP{
		triangleManager: triangleManager,
	}
}

func (h CreateTriangleHTTP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("receive request to create an triangle")

		var req request.Triangle
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Error().Err(err).Msg("Error to decode from json")
			terrors.Handler(w, 500,
				fmt.Errorf("Error to decode from json, err:%s", err.Error()))
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error().Err(err).Msg("Error to validate sides from triangle")
			terrors.Handler(w, 400, err)
			return
		}

		triangle := req.GenerateEntity()
		triangleResult, err := h.triangleManager.Create(triangle)
		if err != nil {
			log.Error().Err(err).Msg("Error to create new triangle")
			terrors.Handler(w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, triangleResult.Response(), http.StatusCreated); err != nil {
			return
		}
	}
}
