package car

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/geekAshish/DriveDesk/service"
	"github.com/gorilla/mux"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := h.service.GetCarById(ctx, id)
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

	w.Header().Set("Content-Type", "applilcation/json");
	w.WriteHeader(http.StatusOK)

	// write the response body
	_, err = w.Write(body);
	if err != nil {
		log.Println("ERROR: ", err)
	}
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"

	res, err := h.service.GetCarByBrand(ctx, brand, isEngine);
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

	w.Header().Set("Content-Type", "applilcation/json");
	w.WriteHeader(http.StatusOK)

	// write the response body
	_, err = w.Write(body);
	if err != nil {
		log.Println("ERROR: ", err)
	}
}