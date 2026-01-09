package models

import (
	"encoding/xml"
)

type InvoicePayload struct {
	XMLName  xml.Name  `xml:"http://www.aade.gr/myDATA/invoice/v1.0 InvoicesDoc"`
	Invoices []Invoice `form:"invoice" xml:"invoice"`
}

type Invoice struct {
	UID            string          `xml:"-"`
	QrURL          string          `xml:"-"`
	QrBase64       string          `xml:"-"`
	LogoImage      string          `xml:"-"`
	MARK           string          `xml:"-"`
	Seller         Company         `json:"issuer" form:"seller" xml:"issuer,omitempty"`
	Byer           Company         `json:"counterpart" form:"buyer" xml:"counterpart"`
	InvoiceHeader  InvoiceHeader   `json:"invoiceHeader" form:"invoiceHeader" xml:"invoiceHeader"`
	PaymentMethods *PaymentMethods `json:"paymentMethods" form:"paymentMethods" xml:"paymentMethods,omitempty"`
	InvoiceDetails []*InvoiceRow   `json:"invoiceDetails" form:"invoiceDetails" xml:"invoiceDetails"`
	InvoiceSummary InvoiceSummary  `json:"invoiceSummary" form:"invoiceSummary" xml:"invoiceSummary"`
	Notices        []string        `json:"notices" form:"notices" xml:"-"`
}

type Company struct {
	CodeNumber    string        `json:"codeNumber" form:"codeNumber" xml:"-"`
	DOI           string        `json:"doi" form:"doi" xml:"-"`
	GEMI          string        `json:"gemi" form:"gemi" xml:"-"`
	Phone         string        `json:"phone" form:"phone" xml:"-"`
	Mobile_Phone  string        `json:"mobile_phone" form:"mobile_phone" xml:"-"`
	Email         string        `json:"email" form:"email" xml:"-"`
	PostalAddress PostalCell    `json:"postalAddress" form:"postalAddress" xml:"-"`
	VatNumber     string        `json:"vatNumber" form:"vatNumber" xml:"vatNumber,omitempty"`
	Country       string        `json:"country" form:"country" xml:"country,omitempty"`
	Branch        int           `json:"branch" form:"branch" xml:"branch"`
	Name          string        `json:"name" form:"name" xml:"name"`
	Address       *AddressType  `json:"address" form:"address" xml:"address,omitempty"`
	Discount      int           `json:"discount" form:"discount" xml:"-"`
	OldBalance    float64       `json:"oldBalance" form:"oldBalance" xml:"-"`
	NewBalance    float64       `json:"newBalance" form:"newBalance" xml:"-"`
	TotalBalance  float64       `json:"totalBalance" form:"totalBalance" xml:"-"`
	BankAccounts  []BankAccount `json:"bankAccounts" form:"bankAccounts" xml:"-"`
}

// BankAccounts for cusotmers and branchcompanies might not be needed
type BranchCompany struct {
	CompanyCode  string        `json:"companyCode" form:"companyCode" xml:"-"`
	BranchCode   string        `json:"branchCode" form:"branchCode" xml:"-"`
	Name         string        `json:"name" form:"name" xml:"-"`
	Address      AddressType   `json:"address" form:"address" xml:"-"`
	Country      string        `json:"country" form:"country" xml:"-"`
	Email        string        `json:"email" form:"email" xml:"-"`
	Phone        string        `json:"phone" form:"phone" xml:"-"`
	Mobile_Phone string        `json:"mobile_phone" form:"mobile_phone" xml:"-"`
	OldBalance   float64       `json:"oldBalance" form:"oldBalance" xml:"-"`
	NewBalance   float64       `json:"newBalance" form:"newBalance" xml:"-"`
	TotalBalance float64       `json:"totalBalance" form:"totalBalance" xml:"-"`
	BankAccounts []BankAccount `json:"bankAccounts" form:"bankAccounts" xml:"-"`
}

type BankAccount struct {
	BankName string `json:"bankName" form:"bankName" xml:"-"`
	IBAN     string `json:"iban" form:"iban" xml:"-"`
}

type PostalCell struct {
	Naming     string `json:"naming" form:"naming" xml:"-"`
	Cellnumber string `json:"cellNumber" form:"cellNumber" xml:"-"`
	PostalCode string `json:"postalcode" form:"postalcode" xml:"-"`
	City       string `json:"city" form:"city" xml:"-"`
}

type AddressType struct {
	Street     string `json:"street,omitempty" form:"street" xml:"street,omitempty"`
	Number     string `json:"number,omitempty" form:"number" xml:"number,omitempty"`
	PostalCode string `json:"postalCode,omitempty" form:"postalCode" xml:"postalCode,omitempty"`
	City       string `json:"city,omitempty" form:"city" xml:"city,omitempty"`
}

type InvoiceHeader struct {
	Time        string `xml:"-"`
	Series      string `json:"series" form:"series" xml:"series"`
	Aa          string `json:"aa" form:"aa" xml:"aa"`
	IssueDate   string `json:"issueDate" form:"issueDate" xml:"issueDate"`
	InvoiceType string `json:"invoiceType" form:"invoiceType" xml:"invoiceType"`
	//Unesseccary values might need an omit empty in the xml naming, we'll see later
	VatPaymentSuspension    bool                    `json:"vatPaymentSuspension" form:"vatPaymentSuspension" xml:"vatPaymentSuspension"`
	Currency                string                  `json:"currency" form:"currency" xml:"currency,omitempty"`
	DispatchDate            string                  `json:"dispatchDate" form:"dispatchDate" xml:"dispatchDate,omitempty"`
	DispatchTime            string                  `json:"dispatchTime" form:"dispatchTime" xml:"dispatchTime,omitempty"`
	VehicleNumber           string                  `json:"vehicleNumber" form:"vehicleNumber" xml:"vehicleNumber,omitempty"`
	MovePurpose             int                     `json:"movePurpose" form:"movePurpose" xml:"movePurpose,omitempty"`
	MovePurposeName         string                  `json:"movePurposeName" form:"movePurposeName" xml:"-"`
	OtherDeliveryNoteHeader OtherDeliveryNoteHeader `json:"otherDeliveryNoteHeader" form:"otherDeliveryNoteHeader" xml:"otherDeliveryNoteHeader"`
	IsDeliveryNote          bool                    `json:"isDeliveryNote" form:"isDeliveryNote" xml:"isDeliveryNote,omitempty"`
}

type OtherDeliveryNoteHeader struct {
	LoadingAddress         AddressType `json:"loadingAddress" form:"loadingAddress" xml:"loadingAddress"`
	DeliveryAddress        AddressType `json:"deliveryAddress" form:"deliveryAddress" xml:"deliveryAddress"`
	StartShippingBranch    int         `json:"startShippingBranch" form:"startShippingBranch" xml:"startShippingBranch"`
	CompleteShippingBranch int         `json:"completeShippingBranch" form:"completeShippingBranch" xml:"completeShippingBranch"`
	DeliveryAddressCode    string      `json:"subcompanycode" form:"subcompanycode" xml:"-"`
	DeliveryAddressName    string      `json:"branchName" form:"branchName" xml:"-"`
}
type PaymentMethods struct {
	Details []PaymentMethodDetail `json:"paymentdatails" form:"paymentdetails" xml:"paymentMethodDetails"`
}
type PaymentMethodDetail struct {
	Type   int     `json:"type" form:"type" xml:"type"`
	Name   string  `json:"name" form:"name" xml:"-"`
	Due    string  `json:"due" form:"due" xml:"-"`
	Amount float64 `json:"amount" form:"amount" xml:"amount"`
}
type InvoiceRow struct {
	CodeNumber             string  `json:"" form:"codeNumber" xml:"-"`
	LineNumber             int     `json:"lineNumber" form:"lineNumber" xml:"lineNumber"`
	ItemDescr              string  `json:"itemDescr,omitempty" form:"itemDescr" xml:"itemDescr,omitempty"`
	Quantity               float64 `json:"quantity,omitempty" form:"quantity" xml:"quantity"`
	MeasurementUnit        int     `json:"measurementUnit" form:"measurementUnit" xml:"measurementUnit"`
	MeasurementUnitName    string  `json:"measurementUnitName" form:"measurementUnitName" xml:"-"`
	UnitNetPrice           float64 `json:"unitPrice,omitempty" form:"unitNetPrice" xml:"-"`
	TotalNetBeforeDiscount float64 `xml:"-"`
	Discount               float64 `xml:"-"`
	DiscountAmount         float64 `xml:"-"`
	// TotalAfterDiscount     float64                     `xml:"-"`
	NetValue               float64                     `json:"netValue" form:"netValue" xml:"netValue"`
	VatCategory            int                         `json:"vatCategory" form:"vatCategory" xml:"vatCategory"`
	VatCategoryName        int                         `json:"vatCategoryName" form:"vatCategoryName" xml:"-"`
	VatAmount              float64                     `json:"vatAmount" form:"vatAmount" xml:"vatAmount"`
	DiscountOption         bool                        `json:"discountOption" form:"discountOption" xml:"discountOption,omitempty"`
	IncomeClassification   *ClassificationItem         `json:"incomeClassification" form:"incomeClassification" xml:"incomeClassification,omitempty"`
	ExpensesClassification *ExpensesClassificationItem `json:"expensesClassification" form:"expensesClassification" xml:"expensesClassification,omitempty"`
}

type InvoiceSummary struct {
	Emptylines             []int                        `xml:"-"`
	TotalNetBeforeDiscount float64                      `xml:"-"`
	TotalDiscount          float64                      `xml:"-"`
	TotalNetValue          float64                      `json:"totalNetValue" form:"totalNetValue" xml:"totalNetValue"`
	TotalVatAmount         float64                      `json:"totalVatAmount" form:"totalVatAmount" xml:"totalVatAmount"`
	TotalWithheldAmount    float64                      `json:"totalWithheldAmount" form:"totalWithheldAmount" xml:"totalWithheldAmount"`
	TotalFeesAmount        float64                      `json:"totalFeesAmount" form:"totalFeesAmount" xml:"totalFeesAmount"`
	TotalStampDutyAmount   float64                      `json:"totalStampDutyAmount" form:"totalStampDutyAmount" xml:"totalStampDutyAmount"`
	TotalOtherTaxesAmount  float64                      `json:"totalOtherTaxesAmount" form:"totalOtherTaxesAmount" xml:"totalOtherTaxesAmount"`
	TotalDeductionsAmount  float64                      `json:"totalDeductionsAmount" form:"totalDeductionsAmount" xml:"totalDeductionsAmount"`
	TotalGrossValue        float64                      `json:"totalGrossValue" form:"totalGrossValue" xml:"totalGrossValue"`
	IncomeClassification   []ClassificationItem         `json:"incomeClassification,omitempty" form:"incomeClassification" xml:"incomeClassification,omitempty"`
	ExpensesClassification []ExpensesClassificationItem `json:"expensesClassification,omitempty" form:"expensesClassification" xml:"expensesClassification,omitempty"`
}

type ExpensesClassificationItem struct {
	ClassificationType     string  `json:"classificationType" form:"classificationType" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 classificationType"`
	ClassificationCategory string  `json:"classificationCategory" form:"classificationCategory" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 classificationCategory"`
	Amount                 float64 `json:"amount" form:"amount" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 amount"`
	VatAmount              float64 `json:"vatAmount" form:"vatAmount" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 vatAmount"`
	VatCategory            int     `json:"vatCategory" form:"vatCategory" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 vatCategory"`
	VatExemptionCategory   int     `json:"vatExemptionCategory" form:"vatExemptionCategory" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 vatExemptionCategory"`
	Id                     int     `json:"id" form:"id" xml:"https://www.aade.gr/myDATA/expensesClassificaton/v1.0 id"`
}

type ClassificationItem struct {
	ClassificationType     string  `json:"classificationType" form:"classificationType" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 classificationType"`
	ClassificationCategory string  `json:"classificationCategory" form:"classificationCategory" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 classificationCategory"`
	Amount                 float64 `json:"amount" form:"amount" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 amount"`
}

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
