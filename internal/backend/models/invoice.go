package models

type Customer struct {
	Something string
}

type Product struct {
	Something string
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
