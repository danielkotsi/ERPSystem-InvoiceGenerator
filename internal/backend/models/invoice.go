package models

type Customer struct {
	VatNumber    string       `json:"vatNumber"`
	Country      string       `json:"country"`
	Branch       int          `json:"branch"`
	EntityType   int          `json:"entityType"`
	Name         string       `json:"name"`
	Address      *AddressType `json:"address,omitempty"`
	Email        *string      `form:"email"`
	Phone        *string      `form:"phone"`
	Mobile_Phone *string      `form:"mobile_phone"`
}

type Customers []Customer
type Products []Product
type Product struct {
	Name         string  `form:"name"`
	Description  string  `form:"description"`
	Product_code string  `form:"sku"`
	Unit_price   float64 `form:"price"`
	Currency     string  `form:"currency"`
	Category     string  `form:"category"`
	Active       bool    `form:"active"`
}
type Order struct {
	Seller   Customer
	Buyer    Customer
	Products map[int]Product
}
