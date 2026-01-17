package models

import "fmt"

type ResponseDoc struct {
	Response ResponseType `xml:"response"`
}

type ResponseType struct {
	Index              int    `xml:"index"`
	StatusCode         string `xml:"statusCode"`
	InvoiceUID         string `xml:"invoiceUid"`
	InvoiceMARK        string `xml:"invoiceMark"`
	ClassificationMARK string `xml:"classificationMark"`
	AuthenticationCode string `xml:"authenticationCode"`
	CancellationMARK   string `xml:"cancellationMark"`
	QRURL              string `xml:"qrUrl"`
	Errors             Error  `xml:"errors"`
}

type Error struct {
	Errors []ErrorType `xml:"error"`
}

type ErrorType struct {
	Message string `xml:"message"`
	Code    string `xml:"code"`
}

func (r *ResponseDoc) Error() string {

	err := fmt.Sprintf("Error from MyData API:%v\n", r.Response.StatusCode)
	for _, errortype := range r.Response.Errors.Errors {
		errorToAdd := fmt.Sprintf("Message:%s\nCode:%s\n", errortype.Message, errortype.Code)
		err += errorToAdd
	}
	return err

}
