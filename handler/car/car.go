package car

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/geekAshish/DriveDesk/models"
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

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();

	body, err := io.ReadAll(r.Body);
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest;
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdCar, err := h.service.CreateCar(ctx, &carReq)
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdCar);
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applilcation/json");
	w.WriteHeader(http.StatusCreated)

	// write the response body
	_, err = w.Write(responseBody);
	if err != nil {
		log.Println("ERROR: ", err)
	}
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();
	params := mux.Vars(r);
	id := params["id"];

	body, err := io.ReadAll(r.Body);
	if err != nil {
		log.Println("ERROR READING REQUEST: ", err);
		w.WriteHeader(http.StatusInternalServerError) 
		return
	}

	var carReq models.CarRequest;
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedCar, err := h.service.UpdateCar(ctx, id, &carReq)
	if err != nil {
		log.Println("ERROR WHILE UPDATING THE CART: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(updatedCar);
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applilcation/json");
	w.WriteHeader(http.StatusOK)

	// write the response body
	_, err = w.Write(responseBody);
	if err != nil {
		log.Println("ERROR: ", err)
	}
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context();

	params := mux.Vars(r);
	id := params["id"]

	deleteCar, err := h.service.DeleteCar(ctx, id);
	if err != nil {
		log.Println("ERROR WHILE DELETEING THE CART: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(deleteCar);
	if err != nil {
		log.Println("ERROR: ", err);
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "applilcation/json");
	w.WriteHeader(http.StatusOK)

	// write the response body
	_, err = w.Write(responseBody);
	if err != nil {
		log.Println("ERROR: ", err)
	}

}