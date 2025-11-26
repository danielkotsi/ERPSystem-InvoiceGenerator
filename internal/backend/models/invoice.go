package models

type Customer struct {
	Code         string
	Name         string `form:"name"`
	Address1     string `form:"address_line1"`
	NumofAdd1    int    `form:"address_num1"`
	Address2     string `form:"address_line2"`
	NumofAdd2    int    `form:"address_num2"`
	City         string `form:"city"`
	State        string `form:"state"`
	Postal_code  string `form:"postal_code"`
	Country      string `form:"country"`
	Email        string `form:"email"`
	Phone        string `form:"phone"`
	Mobile_Phone string `form:"mobile_phone"`
	VAT          string `form:"tax_id"`
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

type Invoice struct {
	Order Order
	MAPK  string
}
