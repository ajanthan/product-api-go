package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ajanthan/product-api-go/dai"
	"github.com/ajanthan/product-api-go/model"
	"github.com/gorilla/mux"
)

func AddProductHandler(productDAI dai.ProductDAI) http.HandlerFunc {
	function := func(w http.ResponseWriter, req *http.Request) {
		product := model.Product{}
		if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.InvalidArgsError)
			return
		}
		if err := productDAI.AddProduct(product); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
	return http.HandlerFunc(function)

}
func GetProductByIDHandler(productDAI dai.ProductDAI) http.HandlerFunc {
	function := func(w http.ResponseWriter, req *http.Request) {
		args := mux.Vars(req)
		productID := args["productID"]
		var product model.Product
		if p, err := productDAI.GetProductByID(productID); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		} else {
			product = p
		}

		if err := json.NewEncoder(w).Encode(product); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		}

	}
	return http.HandlerFunc(function)
}
func GetAllProductsHandler(productDAI dai.ProductDAI) http.HandlerFunc {
	function := func(w http.ResponseWriter, req *http.Request) {
		var products []model.Product
		if p, err := productDAI.GetAllProduct(); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		} else {
			products = p
		}

		if err := json.NewEncoder(w).Encode(products); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		}

	}
	return http.HandlerFunc(function)
}
func DeleteProductHandler(productDAI dai.ProductDAI) http.HandlerFunc {
	function := func(w http.ResponseWriter, req *http.Request) {
		args := mux.Vars(req)
		productID := args["productID"]

		if err := productDAI.DeleteProduct(productID); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.InternalServerError)
			return
		}

	}
	return http.HandlerFunc(function)
}
