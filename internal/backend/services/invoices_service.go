package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"log"
	// "-invoice_manager/internal/utils"
	"encoding/json"
	"net/http"
	"os"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
	MyData  repository.MyData_repo
}

func NewInvoiceService(in repository.Invoice_repo, mydata repository.MyData_repo) *InvoiceService {
	return &InvoiceService{
		Invoice: in,
		MyData:  mydata,
	}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	var invo models.Invoice

	file, err := os.Open("../../finalinvoiceexample.pretty.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&invo)
	if err != nil {
		log.Println(err)
	}

	// err = utils.ParseFormData(r, &invo)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("hello this is the invo from the form", invo)
	//
	// err = s.Invoice.CompleteInvoice(ctx, &invo)
	// if err != nil {
	// 	return nil, err
	// }
	// err = s.MyData.SendInvoice(ctx, &invo)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("\n\n\nthis is the invo QrCodeURL \n\n", invo.QrURL)
	pdf, err = s.Invoice.MakePDF(ctx, &invo)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
