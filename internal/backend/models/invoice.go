package models

type Customer struct {
	Name         string
	Address1     string
	NumofAdd1    int
	Address2     string
	NumofAdd2    int
	City         string
	State        string
	Postal_code  string
	Country      string
	Email        string
	Phone        string
	Mobile_Phone string
	VAT          string
}

type Customers []Customer
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
