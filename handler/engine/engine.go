package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/geekAshish/DriveDesk/models"
	"github.com/geekAshish/DriveDesk/service"
	"github.com/gorilla/mux"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (h *EngineHandler) GetEngineById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := vars["id"]

	res, err := h.service.GetEngineById(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR: ", err)
		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR: ", err)
		return
	}

	w.Header().Set("Content-Type", "applilcation/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	_, err = w.Write(body)
	if err != nil {
		log.Println("ERROR: ", err)
	}
}

func (h *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("ERROR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdEngine, err := h.service.CreateEngine(ctx, &engineReq)
	if err != nil {
		log.Println("ERROR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdEngine)
	if err != nil {
		log.Println("ERROR: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applilcation/json")
	w.WriteHeader(http.StatusCreated)

	// write the response body
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("ERROR: ", err)
	}
}
