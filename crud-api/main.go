package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()


	// Initialize Router
	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)
	RegisterBrandRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	var muxBase = "/api/products"
	
	router.HandleFunc(muxBase, controllers.GetProducts).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.GetProductById).Methods("GET")
	router.HandleFunc(muxBase, controllers.CreateProduct).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.DeleteProduct).Methods("DELETE")
}

func RegisterBrandRoutes(router *mux.Router) {
	var muxBase = "/api/brands"

	router.HandleFunc(muxBase, controllers.GetBrands).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.GetBrandById).Methods("GET")
	router.HandleFunc(muxBase, controllers.CreateBrand).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.UpdateBrand).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.DeleteBrand).Methods("DELETE")
}