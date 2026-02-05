package models

import (
	"invoice_manager/internal/backend/invoice/payload"
)

type CustomerById struct {
	Customer        payload.Company
	BranchCompanies []payload.BranchCompany
}
