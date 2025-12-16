package models

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
	QRURL              string `xml:"qrURL"`
	Errors             Error  `xml:"errors"`
}

type Error struct {
	Errors []ErrorType `xml:"error"`
}

type ErrorType struct {
	Message string `xml:"message"`
	Code    string `xml:"code"`
}
