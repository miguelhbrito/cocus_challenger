package triangle

import (
	"net/http"

	"github.com/cocus_challenger/pkg/mhttp"
	"github.com/cocus_challenger/pkg/terrors"
	"github.com/rs/zerolog/log"
)

type (
	ListTrianglesHTTP struct {
		triangleManager Triangle
	}
)

func NewListTrianglesHTTP(
	triangleManager Triangle,
) mhttp.HttpHandler {
	return ListTrianglesHTTP{
		triangleManager: triangleManager,
	}
}

func (h ListTrianglesHTTP) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("receive request to list all triangles")

		ts, err := h.triangleManager.List()
		if err != nil {
			log.Error().Err(err).Msg("Error to list all triangles")
			terrors.Handler(w, 500, err)
			return
		}

		if err := mhttp.WriteJsonResponse(w, ts.Response(), http.StatusOK); err != nil {
			terrors.Handler(w, http.StatusOK, err)
			return
		}
	}
}
