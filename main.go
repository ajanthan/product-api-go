package main

import (
	"log"
	"net/http"

	"github.com/ajanthan/product-api-go/api"
	"github.com/ajanthan/product-api-go/dai"
	"github.com/gorilla/mux"
)

func main() {
	productDAI := dai.ProductDAIMysql{}
	if err := productDAI.Init(); err != nil {
		log.Panic(err.Error())
	}
	log.Println("Initialized the product backend storage")
	router := mux.NewRouter()
	router.StrictSlash(true)
	addProductHandler := api.AddProductHandler(productDAI)
	getProductHandler := api.GetProductByIDHandler(productDAI)
	deleteProductHandler := api.DeleteProductHandler(productDAI)
	getAllProductsHandler := api.GetAllProductsHandler(productDAI)
	router.HandleFunc("/product", addProductHandler).Methods("POST")
	router.HandleFunc("/product/{productID}", getProductHandler).Methods("GET")
	router.HandleFunc("/product/{productID}", deleteProductHandler).Methods("DELETE")
	router.HandleFunc("/products", getAllProductsHandler).Methods("GET")

	log.Println("Access the server http://localhost:8888/product")
	log.Fatal(http.ListenAndServe(":8888", router))

}
