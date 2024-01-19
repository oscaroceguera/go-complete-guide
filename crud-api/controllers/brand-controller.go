package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/entities"
	"golang-crud-rest-api/repos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


var brandRepo = repos.NewBrandRepo()

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var brand entities.Brand
	json.NewDecoder(r.Body).Decode(&brand)
	brand = brandRepo.Create(brand)
	json.NewEncoder(w).Encode(brand)
	w.WriteHeader((http.StatusCreated))
}

func GetBrands(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brandRepo.GetList())
	w.WriteHeader(http.StatusOK)
}

func GetBrandById(w http.ResponseWriter, r *http.Request) {
	brandIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	brand, err := brandRepo.GetOne(uint(brandIdLong))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Brand not found!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brand)
	w.WriteHeader(http.StatusOK)
}

func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	brandIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	var brand entities.Brand
	json.NewDecoder(r.Body).Decode(&brand)
	_, err = brandRepo.Update(uint(brandIdLong), brand)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Brand not found!")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	brandIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}
	_, err = brandRepo.DeleteOne(uint(brandIdLong))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("BRand not found!")
		return
	}

	json.NewEncoder(w).Encode("Brand Deleted Successfully!")
	w.WriteHeader(http.StatusNoContent)
}