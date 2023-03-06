package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/liperm/go-api/constants"
	"github.com/liperm/go-api/formatters"
	"github.com/liperm/go-api/models"
	"github.com/liperm/go-api/repository"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	customers := repository.GetItems()
	json.NewEncoder(w).Encode(customers)
}

func GetItemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	item := repository.GetItemById(id)

	if item.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func GetItemByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	item := repository.GetItemByCode(code)

	if item.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(code)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	id, message := repository.CreateItem(item)

	if message != "OK" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error creating item: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.WriteHeader(http.StatusCreated)
	succsessResponse := formatters.FormatSuccessResponse(id)
	json.NewEncoder(w).Encode(succsessResponse)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	item := repository.GetItemById(id)

	if item.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	json.NewDecoder(r.Body).Decode(&item)
	id, message := repository.UpdateItem(&item)

	if message != constants.OK {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error updating item: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := formatters.FormatSuccessResponse(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	item := repository.GetItemById(id)

	if item.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		notFoundResponse := formatters.FormatNotFoundResponse(id)
		json.NewEncoder(w).Encode(notFoundResponse)
		return
	}
	id, message := repository.DeleteItem(&item)

	if message != constants.OK {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := formatters.FormatErrorResponse("Error deleting item: " + message)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := formatters.FormatSuccessResponse(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}
