package sqlite

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"fmt"
	"strconv"
	"time"
)

func (r *InvoiceRepo) GetBuyerBalance(ctx context.Context, buyer *payload.Company) error {
	query := "select Balance from customers where CodeNumber=?;"
	rows, err := r.DB.QueryContext(ctx, query, buyer.CodeNumber)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&buyer.OldBalance); err != nil {
			return err
		}
	}

	return nil
}
func (r *InvoiceRepo) GetSellerInfo(ctx context.Context, seller *payload.Company) error {
	query := `select PostalCellName, PostalCellNumber,PostalCellPostalCode, PostalCellCity from users where CodeNumber==?;`

	rows, err := r.DB.QueryContext(ctx, query, seller.CodeNumber)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&seller.PostalAddress.Naming, &seller.PostalAddress.Cellnumber, &seller.PostalAddress.PostalCode, &seller.PostalAddress.City); err != nil {
			return err
		}
	}

	fmt.Println("hello this is the code", seller.CodeNumber)
	fmt.Println("hello this is the postall cell name", seller.PostalAddress.Naming)
	return nil
}
func (r *InvoiceRepo) CompletePaymentMethods(paymentmethods *payload.PaymentMethods, buyer *payload.Company, totalgrossamount float64) error {
	paymenttypes := map[string]int{
		"Επαγ. Λογαριασμός Πληρωμών Ημεδαπής":  1,
		"Επαγ. Λογαριασμός Πληρωμών Αλλοδαπής": 2,
		"Μετρητά":              3,
		"Επιταγή":              4,
		"Επί Πιστώσει":         5,
		"Web Banking":          6,
		"POS / e-POS":          7,
		"Άμεσες Πληρωμές IRIS": 8,
	}
	for i, payment := range paymentmethods.Details {
		paymentmethods.Details[i].Type = paymenttypes[payment.Name]
		paymentmethods.Details[i].Amount = totalgrossamount
		if paymentmethods.Details[i].Type == 5 {
			buyer.NewBalance = utils.RoundTo2(buyer.OldBalance + totalgrossamount)
		}
	}

	return nil
}
func (r *InvoiceRepo) UpdateBalance(ctx context.Context, buyerCodeNumber string, buyerNewBalance float64) error {
	query := "update customers set Balance=? where CodeNumber==?;"
	if _, err := r.DB.ExecContext(ctx, query, buyerNewBalance, buyerCodeNumber); err != nil {
		return err
	}
	return nil
}
func (r *InvoiceRepo) AddToAA(ctx context.Context, invoicetype, aa string) error {
	aaint, err := strconv.Atoi(aa)
	if err != nil {
		return err
	}
	aaint++
	aa = fmt.Sprintf("%05d", aaint)
	fmt.Println(aa)
	query := `update user_invoice_types_series set aa=? where invoice_type==?;`

	if _, err := r.DB.ExecContext(ctx, query, aa, invoicetype); err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) CompleteHTMLinfo(invoiceinfo *models.InvoiceHTMLinfo, invoicetype string) error {
	invoiceinfo.Invoiceinfo.Currency = "EUR"
	invoiceinfo.Invoiceinfo.Invoicetype = invoicetype
	switch invoicetype {
	case "1.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_2"
		invoiceinfo.Invoiceinfo.MovePurpose = "1"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case "8.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_8"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
		invoiceinfo.Invoiceinfo.VatCategory = 8
	case "9.3":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = ""
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category3"
		invoiceinfo.Invoiceinfo.MovePurpose = "3"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case "13.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_201"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category2_2"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
	}
	return nil
}

func (r *InvoiceRepo) CompleteInvoiceHeader(header *payload.InvoiceHeader) error {
	movepurpses := map[int]string{
		1: "Πώληση",
		3: "Δειγματισμός",
		9: "Αγορά",
	}
	header.IssueDate = time.Now().Format("2006-01-02")
	header.Time = time.Now().Format("15:04")
	header.MovePurposeName = movepurpses[header.MovePurpose]
	return nil
}
