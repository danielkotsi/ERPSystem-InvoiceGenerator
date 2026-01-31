package models

type Product struct {
	CodeNumber          string  `json:"codeNumber" form:"codeNumber"`
	Name                string  `json:"name" form:"name"`
	Description         string  `json:"description" form:"description"`
	Unit_Net_Price      float64 `json:"unitNetPrice" form:"unitNetPrice"`
	MeasurementUnitCode int     `json:"measurementUnitCode" form:"measurementUnitCode"`
	MeasurementUnit     string  `json:"measurementUnit" form:"measurementUnit"`
	VatCategory         int     `json:"vatCategory" form:"vatCategory"`
	ProductCategory     *string `json:"productCategory,omitempty" form:"productCategory"`
}

type ProductSuggestion struct {
	CodeNumber string `json:"codeNumber" form:"codeNumber"`
	Name       string `json:"name" form:"name"`
}
