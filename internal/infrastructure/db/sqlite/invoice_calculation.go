package sqlite

import (
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
)

// func (r *InvoiceRepo) CalculateAlltheInvoiceLines(invoicetype string, paymentmethods *payload.PaymentMethods, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company) error {
// 	switch invoicetype {
// 	case "1.1":
// 		if err := r.SellingInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
// 			return err
// 		}
// 	case "8.1":
// 		if err := r.RecieptInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
// 			return err
// 		}
// 	case "13.1":
// 		if err := r.BuyingInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
// 			return err
// 		}
// 	case "9.3":
// 		if err := r.DeliveryNoteInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
// 			return err
// 		}
// 	default:
// 	}
// 	return nil
// }
//
// func (r *InvoiceRepo) SellingInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
// 	emptylines := 24
// 	for i, line := range invoicelines {
// 		emptylines--
// 		line.VatCategoryName = utils.VatNames(line.VatCategory)
// 		line.LineNumber = i + 1
// 		if err := r.InvoiceLinePrices(line, buyer.Discount); err != nil {
// 			return err
// 		}
// 		line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */
// 		summary.TotalDiscount += line.DiscountAmount
// 		summary.TotalNetBeforeDiscount += line.TotalNetBeforeDiscount
// 		summary.TotalNetBeforeDiscount = utils.RoundTo2(summary.TotalNetBeforeDiscount)
// 		summary.TotalNetValue += line.NetValue
// 		summary.TotalNetValue = utils.RoundTo2(summary.TotalNetValue)
// 		summary.TotalVatAmount += line.VatAmount
// 		summary.TotalVatAmount = utils.RoundTo2(summary.TotalVatAmount)
// 		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
// 			return err
// 		}
// 	}
// 	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
// 	buyer.NewBalance = buyer.OldBalance
//
// 	if err := r.CompletePaymentMethods(paymentmethods, buyer, summary.TotalGrossValue); err != nil {
// 		return err
// 	}
//
// 	summary.Emptylines = make([]int, emptylines)
// 	return nil
// }
//
// func (r *InvoiceRepo) DeliveryNoteInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
// 	emptylines := 24
// 	for i, line := range invoicelines {
// 		emptylines--
// 		line.LineNumber = i + 1
// 		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
// 			return err
// 		}
// 	}
// 	buyer.NewBalance = buyer.OldBalance
// 	summary.Emptylines = make([]int, emptylines)
// 	return nil
// }
// func (r *InvoiceRepo) BuyingInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
// 	emptylines := 24
// 	for i, line := range invoicelines {
// 		emptylines--
// 		line.VatCategoryName = utils.VatNames(line.VatCategory)
// 		line.LineNumber = i + 1
// 		if err := r.InvoiceLinePrices(line, buyer.Discount); err != nil {
// 			return err
// 		}
// 		line.ExpensesClassification.Amount = line.NetValue /* + line.VatAm unt */
// 		summary.TotalDiscount += line.DiscountAmount
// 		summary.TotalNetBeforeDiscount += line.TotalNetBeforeDiscount
// 		summary.TotalNetBeforeDiscount = utils.RoundTo2(summary.TotalNetBeforeDiscount)
// 		summary.TotalNetValue += line.NetValue
// 		summary.TotalNetValue = utils.RoundTo2(summary.TotalNetValue)
// 		summary.TotalVatAmount += line.VatAmount
// 		summary.TotalVatAmount = utils.RoundTo2(summary.TotalVatAmount)
// 		if err := r.AddExpenseClassificationInSummary(line.ExpensesClassification, summary); err != nil {
// 			return err
// 		}
// 	}
// 	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
// 	summary.Emptylines = make([]int, emptylines)
// 	return nil
// }
//
// func (r *InvoiceRepo) InvoiceLinePrices(line *payload.InvoiceRow, discount int) error {
// 	line.Discount = float64(discount)
// 	floatdiscount := float64(discount) / 100
//
// 	totalNetPriceBeforeDiscount := *line.Quantity * line.UnitNetPrice
// 	line.DiscountAmount = utils.RoundTo2(totalNetPriceBeforeDiscount * floatdiscount)
// 	totalNetPriceAfterDiscount := totalNetPriceBeforeDiscount - line.DiscountAmount
// 	vatAfterDiscount := totalNetPriceAfterDiscount * utils.Vat(line.VatCategory)
//
// 	line.TotalNetBeforeDiscount = utils.RoundTo2(totalNetPriceBeforeDiscount)
// 	line.NetValue = utils.RoundTo2(totalNetPriceAfterDiscount)
// 	line.VatAmount = utils.RoundTo2(vatAfterDiscount)
//
// 	return nil
// }
// func (r *InvoiceRepo) RecieptInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
// 	emptylines := 24
// 	for i, line := range invoicelines {
// 		emptylines--
// 		line.LineNumber = i + 1
// 		line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */
// 		summary.TotalNetValue = line.NetValue
// 		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
// 			return err
// 		}
// 	}
// 	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
// 	buyer.NewBalance = buyer.OldBalance - summary.TotalGrossValue
// 	if err := r.CompletePaymentMethods(paymentmethods, buyer, summary.TotalGrossValue); err != nil {
// 		return err
// 	}
// 	summary.Emptylines = make([]int, emptylines)
// 	return nil
// }
//
// func (r *InvoiceRepo) AddIncomeClassificationInSummary(classificationItem *payload.ClassificationItem, summary *payload.InvoiceSummary) error {
// 	index, exists := r.IncomeCategoryExists(*classificationItem, summary.IncomeClassification)
// 	if exists {
// 		summary.IncomeClassification[index].Amount += classificationItem.Amount
// 	} else {
// 		summary.IncomeClassification = append(summary.IncomeClassification, *classificationItem)
// 	}
// 	return nil
// }
//
// func (r *InvoiceRepo) AddExpenseClassificationInSummary(classificationItem *payload.ExpensesClassificationItem, summary *payload.InvoiceSummary) error {
// 	index, exists := r.ExpenseCategoryExists(*classificationItem, summary.ExpensesClassification)
// 	if exists {
// 		summary.ExpensesClassification[index].Amount += classificationItem.Amount
// 	} else {
// 		summary.ExpensesClassification = append(summary.ExpensesClassification, *classificationItem)
// 	}
// 	return nil
// }
//
// func (r *InvoiceRepo) IncomeCategoryExists(classificationitem payload.ClassificationItem, summary []payload.ClassificationItem) (int, bool) {
// 	for index, category := range summary {
// 		if classificationitem.ClassificationCategory == category.ClassificationCategory && classificationitem.ClassificationType == category.ClassificationType {
// 			return index, true
// 		}
// 	}
//
// 	return 0, false
// }
//
// func (r *InvoiceRepo) ExpenseCategoryExists(classificationitem payload.ExpensesClassificationItem, summary []payload.ExpensesClassificationItem) (int, bool) {
// 	for index, category := range summary {
// 		if classificationitem.ClassificationCategory == category.ClassificationCategory && classificationitem.ClassificationType == category.ClassificationType {
// 			return index, true
// 		}
// 	}
//
// 	return 0, false
// }
