package models

type Product struct {
	ProductList map[string]ProductRequest
}
type ProductRequest struct {
	Id          string `json:"id"`
	ProductName string `json:"productName" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,min=1" `
	Price       string `json:"price" binding:"required,min=1"`
	Category    string `json:"category"`
}

type AddProductResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	ProductId string `json:"productId"`
}

var Product_List *Product

func InitiateProductList() *Product {
	if Product_List == nil {
		Product_List = new(Product)
	}
	return Product_List
}
