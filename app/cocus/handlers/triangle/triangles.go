package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cocus_challenger_refact/app/cocus/mhttp"
	"github.com/cocus_challenger_refact/app/cocus/terrors"
	core "github.com/cocus_challenger_refact/business/core/triangle"
	"github.com/cocus_challenger_refact/business/data/triangle"
)

type TriangleHandlers struct {
	log             *log.Logger
	TriangleManager core.TriangleInt
}

func (h *TriangleHandlers) Create(w http.ResponseWriter, r *http.Request) {
	//h.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)
	//log.Debug().Msg("receive request to create an triangle")
	//h.log.Printf("receive request to create an triangle")

	var req triangle.NewTriangle
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		//log.Error().Err(err).Msg("Error to decode from json")
		h.log.Printf("Error to decode from json, err %s", err)
		terrors.Handler(w, http.StatusInternalServerError,
			fmt.Errorf("Error to decode from json, err:%s", err.Error()))
		return
	}

	err = req.Validate()
	if err != nil {
		//log.Error().Err(err).Msg("Error to validate sides from triangle")
		h.log.Printf("Error to validate sides from triangle, err %s", err)
		terrors.Handler(w, http.StatusBadRequest, err)
		return
	}

	triangle := req.GenerateEntity()
	triangleResult, err := h.TriangleManager.Create(triangle)
	if err != nil {
		//log.Error().Err(err).Msg("Error to create new triangle")
		//h.log.Printf("Error to create new triangle, err %s", err)
		terrors.Handler(w, http.StatusInternalServerError, err)
		return
	}

	if err := mhttp.WriteJsonResponse(w, triangleResult.Response(), http.StatusCreated); err != nil {
		terrors.Handler(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *TriangleHandlers) List(w http.ResponseWriter, r *http.Request) {
	//h.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)
	//log.Debug().Msg("receive request to list all triangles")
	//h.log.Printf("receive request to list all triangles")

	ts, err := h.TriangleManager.List()
	if err != nil {
		//log.Error().Err(err).Msg("Error to list all triangles")
		//h.log.Printf("Error to list all triangles, err %s", err)
		terrors.Handler(w, http.StatusInternalServerError, err)
		return
	}

	if err := mhttp.WriteJsonResponse(w, ts.Response(), http.StatusOK); err != nil {
		terrors.Handler(w, http.StatusInternalServerError, err)
		return
	}
}
