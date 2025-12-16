package models

type Customer struct {
	VatNumber    string       `json:"vatNumber" form:"vatNumber"`
	Country      string       `json:"country" form:"country"`
	Branch       int          `json:"branch" form:"branch"`
	EntityType   int          `json:"entityType" form:"entityType"`
	Name         string       `json:"name" form:"name"`
	Address      *AddressType `json:"address,omitempty" form:"address"`
	Email        *string      `form:"email"`
	Phone        *string      `form:"phone"`
	Mobile_Phone *string      `form:"mobile_phone"`
}

//	type AddressType struct {
//		Street     *string `json:"street,omitempty" form:"street"`
//		Number     *string `json:"number,omitempty" form:"number"`
//		City       *string `json:"city,omitempty" form:"city"`
//		PostalCode *string `json:"postalCode,omitempty" form:"postalCode"`
//	}
type Customers []Customer
type Products []Product
type Product struct {
	Name         string  `form:"name"`
	Description  string  `form:"description"`
	Product_code string  `form:"sku"`
	Unit_price   float64 `form:"price"`
	Vat_Category string  `form:"vat_category"`
	Category     string  `form:"category"`
	Active       bool    `form:"active"`
}
type Order struct {
	Seller   Customer
	Buyer    Customer
	Products map[int]Product
}
