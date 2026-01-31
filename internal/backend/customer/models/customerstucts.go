package models

import (
	"-invoice_manager/internal/backend/invoice/payload"
)

type CustomerById struct {
	Customer        payload.Company
	BranchCompanies []payload.BranchCompany
}

type CustomerSuggestion struct {
	Name       string `json:"name"`
	CodeNumber string `json:"codeNumber"`
}

type BranchSuggestion struct {
	Name        string `json:"name"`
	CompanyCode string `json:"companyCode"`
	CodeNumber  string `json:"codeNumber"`
}
