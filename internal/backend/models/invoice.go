package models

type Customer struct {
	Something string
}

type Products []Product
type Product struct {
	Name         string
	Description  string
	Product_code string
	Unit_price   float64
	Category     string
	Active       bool
}
type Order struct {
	Seller   Customer
	Buyer    Customer
	Products map[int]Product
}

type Invoice struct {
	Order Order
	MAPK  string
}
