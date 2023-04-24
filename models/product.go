package models

type Product struct {
	ProductList map[string]ProductRequest
}
type ProductRequest struct {
	Id          string `json:"id"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
	Category    string `json:"category"`
}

var Product_List *Product

func InitiateProductList() *Product {
	if Product_List == nil {
		Product_List = new(Product)
	}
	return Product_List
}
