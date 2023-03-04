package models

type Model interface{}

type Order struct {
	ID          int64 `sql:"ID"`
	UserID      int64 `sql:"userID"`
	ProductId   int64 `sql:"productID"`
	OrderStatus int64 `sql:"orderStatus"`
}

type Product struct {
	ID           int64  `json:"id" sql:"ID" seckill:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" seckill:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" seckill:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" seckill:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" seckill:"ProductUrl"`
}
