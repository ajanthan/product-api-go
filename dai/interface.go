package dai

import "github.com/ajanthan/product-api-go/model"

type ProductDAI interface {
	Init() error
	AddProduct(product model.Product) error
	GetProductByID(productID string) (model.Product, error)
	DeleteProduct(productID string) error
	GetAllProduct() ([]model.Product, error)
}
