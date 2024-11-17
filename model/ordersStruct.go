package model

type OrderDetails struct {
	OrderID   string
	ProductID string
	Quantity  uint
	Price     uint
	Total     uint
}
type Order struct { //corresponds to the sql table
	UserID  string
	OrderID string
}
type OrdersDetail struct { //corresponds to the whole order
	UserID            string
	OrderID           string
	PurchasedProducts []struct {
		ProductID string
		Quantity  uint
		Price     uint
		Total     uint
	}
}
