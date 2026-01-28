package types

import (
	"-invoice_manager/internal/backend/invoice/payload"
	"fmt"
	"github.com/signintech/gopdf"
	"log"
	"os"
	"strconv"
	"strings"
)

type TextRun struct {
	Font string
	Size float64
	Text string
}

type WordRun struct {
	Font  string
	Size  float64
	Text  string
	Width float64
}

type LineRun struct {
	Words []WordRun
	Width float64
}

func GeneratePDFfromTemp() (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}

	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	data, err := os.ReadFile("../../assets/pdftemplates/invoicetemplate.pdf")
	if err != nil {
		return nil, fmt.Errorf("couldn't read pdf template %w", err)
	}
	err = pdf.ImportPagesFromSource(data, "/MediaBox")
	if err != nil {
		return nil, fmt.Errorf("couldn't load template into pdf %w", err)
	}
	return pdf, nil
}
func Addlineswithnewlines(pdf *gopdf.GoPdf, x, y float64, rect *gopdf.Rect, rowspace float64, lines []string) {

	rightEdge := x + rect.W

	currY := y
	for _, line := range lines {
		textWidth, _ := pdf.MeasureTextWidth(line)
		offsetX := rightEdge - textWidth

		pdf.SetXY(offsetX, currY)
		pdf.Text(line)

		currY += rowspace
	}
}

type Item struct {
	Description string
	Qty         string
	Price       string
}

func SelectVatFields(summary payload.InvoiceSummary) []string {
	return []string{
		strconv.Itoa(int(summary.TotalNetValue)),
		"",
		strconv.Itoa(int(summary.TotalVatAmount)),
		strconv.Itoa(int(summary.TotalGrossValue)),
	}
}
func AddVatRows(pdf *gopdf.GoPdf, x float64, y float64, colWidths []float64, lineHeight float64, summary payload.InvoiceSummary, rowspace float64) float64 {
	fields := SelectVatFields(summary)
	nCols := len(fields)

	wrapped := make([][]string, nCols)
	maxLines := 0
	for c := 0; c < nCols; c++ {
		wrapped[c] = wrapTextLines(pdf, fields[c], colWidths[c])
		if len(wrapped[c]) > maxLines {
			maxLines = len(wrapped[c])
		}
	}
	for i := 0; i < maxLines; i++ {
		currY := y + float64(i)*lineHeight
		currX := x
		for c := 0; c < nCols; c++ {
			if i < len(wrapped[c]) {
				text := wrapped[c][i]
				textWidth, _ := pdf.MeasureTextWidth(text)
				var offsetX float64
				if c == 1 {
					offsetX = colWidths[c] - textWidth - 4
				} else {
					offsetX = (colWidths[c] - textWidth) / 2
				}
				pdf.SetXY(currX+offsetX, currY)
				pdf.Text(text)
			}

			currX += colWidths[c]
		}
	}
	y += float64(maxLines)*lineHeight + rowspace

	return y
}
func AddRows(pdf *gopdf.GoPdf, x float64, y float64, colWidths []float64, lineHeight float64, items []*payload.InvoiceRow, rowspace float64) float64 {
	for _, item := range items {
		fields := SelectInvoiceRowFields(*item)
		nCols := len(fields)
		wrapped := make([][]string, nCols)
		maxLines := 0
		for c := 0; c < nCols; c++ {
			wrapped[c] = wrapTextLines(pdf, fields[c], colWidths[c])
			if len(wrapped[c]) > maxLines {
				maxLines = len(wrapped[c])
			}
		}
		for i := 0; i < maxLines; i++ {
			currY := y + float64(i)*lineHeight
			currX := x
			for c := 0; c < nCols; c++ {
				if i < len(wrapped[c]) {
					text := wrapped[c][i]
					textWidth, _ := pdf.MeasureTextWidth(text)
					var offsetX float64
					switch c {
					case 0, 2:
						offsetX = (colWidths[c] - textWidth) / 2
					case 6:
						offsetX = colWidths[c] - textWidth - 1
					default:
						offsetX = colWidths[c] - textWidth - 4
					}
					pdf.SetXY(currX+offsetX, currY)
					pdf.Text(text)
				}

				currX += colWidths[c]
			}
		}
		y += float64(maxLines)*lineHeight + rowspace
	}
	return y
}
func wrapTextLines(pdf *gopdf.GoPdf, text string, maxWidth float64) []string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}

	var lines []string
	var currLine string

	for _, word := range words {
		testLine := currLine
		if testLine == "" {
			testLine = word
		} else {
			testLine += " " + word
		}

		if num, _ := pdf.MeasureTextWidth(testLine); num <= maxWidth {
			currLine = testLine
		} else {
			lines = append(lines, currLine)
			currLine = word
		}
	}

	if currLine != "" {
		lines = append(lines, currLine)
	}

	return lines
}
func SelectDeliveryFields(invo *payload.Invoice) []string {
	return []string{
		invo.InvoiceHeader.MovePurposeName,
		invo.Seller.Address.Street + " " + invo.Seller.Address.Number,
		invo.InvoiceHeader.OtherDeliveryNoteHeader.DeliveryAddress.Street + " " + invo.InvoiceHeader.OtherDeliveryNoteHeader.DeliveryAddress.Number,
		invo.InvoiceHeader.OtherDeliveryNoteHeader.DeliveryAddress.City,
		"",
		safePaymentMethod(invo),
	}
}
func safePaymentMethod(invo *payload.Invoice) string {
	if invo.PaymentMethods == nil {
		return ""
	}

	if len(invo.PaymentMethods.Details) == 0 {
		return ""
	}

	pm := invo.PaymentMethods.Details[0]
	return pm.Name + pm.Due
}
func SelectBuyerFields(buyer payload.Company) []string {
	return []string{
		buyer.CodeNumber,
		*buyer.Name,
		"hello there",
		buyer.Address.Street + " " + buyer.Address.Number,
		buyer.Address.City + " " + buyer.Address.PostalCode,
		buyer.Phone,
		buyer.VatNumber,
		buyer.DOI,
	}
}
func SelectInvoiceRowFields(row payload.InvoiceRow) []string {
	return []string{
		row.CodeNumber,
		row.ItemDescr,
		row.MeasurementUnitName,
		floatValue(row.Quantity),
		strconv.FormatFloat(row.UnitNetPrice, 'f', 2, 64),
		strconv.FormatFloat(row.TotalNetBeforeDiscount, 'f', 2, 64),
		strconv.FormatFloat(row.Discount, 'f', 2, 64),
		strconv.FormatFloat(row.DiscountAmount, 'f', 2, 64),
		strconv.FormatFloat(row.NetValue, 'f', 2, 64),
		strconv.Itoa(int(row.VatCategoryName)),
		strconv.FormatFloat(row.VatAmount, 'f', 2, 64),
	}
}
func floatValue(s *float64) string {
	if s == nil {
		return ""
	}
	return strconv.FormatFloat(*s, 'f', 2, 64)
}
func MakeBalance(pdf *gopdf.GoPdf, invo *payload.Invoice) {
	err := pdf.SetFont("OpenSans", "", 9)
	if err != nil {
		log.Fatal(err)
	}
	lines := []string{strconv.FormatFloat(invo.Byer.OldBalance, 'f', 2, 64), strconv.FormatFloat(invo.InvoiceSummary.TotalGrossValue, 'f', 2, 64), strconv.FormatFloat(invo.Byer.NewBalance, 'f', 2, 64)}
	Addlineswithnewlines(pdf, 120, 628, &gopdf.Rect{W: 50, H: 250}, 15, lines)
}
func MakeInvoiceHeader(pdf *gopdf.GoPdf, invo *payload.Invoice) {
	err := pdf.SetFont("OpenSans", "", 9)
	if err != nil {
		log.Fatal(err)
	}
	pdf.SetXY(340, 135)
	pdf.Cell(nil, invo.InvoiceHeader.Series+invo.InvoiceHeader.Aa)

	pdf.SetXY(450, 135)
	pdf.Cell(nil, invo.InvoiceHeader.IssueDate)

	pdf.SetXY(542, 135)
	pdf.Cell(nil, invo.InvoiceHeader.Time)
}
func MakeDetails(pdf *gopdf.GoPdf, details []*payload.InvoiceRow) {
	err := pdf.SetFont("OpenSans", "", 7)
	if err != nil {
		log.Fatal(err)
	}

	x := 20.0
	y := 295.0
	colWidths := []float64{38, 153, 52, 39, 58, 52.5, 20.6, 32.5, 51, 23, 37, 15}
	lineHeight := 6.0
	rowspace := 9.0

	AddRows(pdf, x, y, colWidths, lineHeight, details, rowspace)
}

func MakeByer(pdf *gopdf.GoPdf, buyer payload.Company) {
	err := pdf.SetFont("OpenSans", "", 6.7)
	if err != nil {
		log.Fatal(err)
	}
	words := SelectBuyerFields(buyer)
	Addlineswithnewlines(pdf, 120, 174, &gopdf.Rect{W: 50, H: 250}, 12, words)
}
func MakeDelivery(pdf *gopdf.GoPdf, invo *payload.Invoice) {
	err := pdf.SetFont("OpenSans", "", 6.7)
	if err != nil {
		log.Fatal(err)
	}
	words := SelectDeliveryFields(invo)
	Addlineswithnewlines(pdf, 520, 199, &gopdf.Rect{W: 50, H: 250}, 12, words)
}

func MakePrices(pdf *gopdf.GoPdf, summary payload.InvoiceSummary) {
	err := pdf.SetFont("OpenSans", "", 9)
	if err != nil {
		log.Fatal(err)
	}
	words := SelectPricesFields(summary)
	Addlineswithnewlines(pdf, 520, 628, &gopdf.Rect{W: 50, H: 250}, 18, words)
}

func SelectPricesFields(summary payload.InvoiceSummary) []string {
	return []string{
		strconv.FormatFloat(summary.TotalNetBeforeDiscount, 'f', 2, 64),
		strconv.FormatFloat(summary.TotalDiscount, 'f', 2, 64),
		strconv.FormatFloat(summary.TotalNetValue, 'f', 2, 64),
		strconv.FormatFloat(summary.TotalVatAmount, 'f', 2, 64),
		strconv.FormatFloat(summary.TotalGrossValue, 'f', 2, 64),
	}
}
func MakeHeader(pdf *gopdf.GoPdf, invo *payload.Invoice) {

	addressrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: invo.Seller.Address.Street + " " + invo.Seller.Address.Number,
	}
	emptyrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "    ",
	}
	cityrun := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.Address.PostalCode + "  " + invo.Seller.Address.City,
	}
	runs := []TextRun{addressrun, emptyrun, cityrun}
	WrappBoldNormal(pdf, runs, 220, 53, 250, 10)
	postalrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "ΔΙΕΥΘΥΝΣΗ ΑΛΛΗΛΟΓΡΑΦΙΑΣ:  ",
	}

	postalinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.PostalAddress.Naming + " " + invo.Seller.PostalAddress.Cellnumber + " " + invo.Seller.PostalAddress.PostalCode + " " + invo.Seller.PostalAddress.City,
	}
	runs = []TextRun{postalrun, postalinfo}
	WrappBoldNormal(pdf, runs, 220, 61, 160, 7)
	telrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "ΤΗΛ.:  ",
	}

	telinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.Phone,
	}
	runs = []TextRun{telrun, telinfo}
	WrappBoldNormal(pdf, runs, 220, 76, 160, 6)
	whatsApp := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "WhatsApp:  ",
	}

	whatsappinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.Mobile_Phone,
	}
	runs = []TextRun{whatsApp, whatsappinfo}
	WrappBoldNormal(pdf, runs, 285, 76, 160, 6)
	emailrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "E-MAIL:",
	}

	emailinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.Email + "        www.pastry.blogspot.com",
	}
	runs = []TextRun{emailrun, emailinfo}
	WrappBoldNormal(pdf, runs, 220, 84, 210, 6)
	vatrun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "Α.Φ.Μ.: ",
	}

	vatinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.VatNumber,
	}
	doirun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "    Δ.Ο.Υ.:",
	}

	doinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.DOI,
	}
	runs = []TextRun{vatrun, vatinfo, doirun, doinfo}
	WrappBoldNormal(pdf, runs, 220, 92, 210, 6)
	gemirun := TextRun{
		Font: "OpenSansBold",
		Size: 6.5,
		Text: "ΓΕΜΗ:   ",
	}
	gemiinfo := TextRun{
		Font: "OpenSans",
		Size: 6.5,
		Text: invo.Seller.GEMI,
	}
	runs = []TextRun{gemirun, gemiinfo}
	WrappBoldNormal(pdf, runs, 220, 100, 210, 6)
	err := pdf.SetFont("OpenSansBold", "", 10)
	if err != nil {
		log.Fatal(err)
	}
	WrapText(pdf, *invo.Seller.Name, 250, 220, 35, 10)
}
func WrapText(pdf *gopdf.GoPdf, text string, blockwidth, x, y, spacebetween float64) {
	lines := PrepareTextforWrapping(
		pdf,
		text,
		250,
	)

	for _, line := range lines {
		pdf.SetXY(x, y)
		pdf.Text(line)
		y += spacebetween
	}
}
func PrepareTextforWrapping(
	pdf *gopdf.GoPdf,
	text string,
	maxWidth float64,
) []string {
	words := strings.Fields(text) // handles multiple spaces
	if len(words) == 0 {
		return nil
	}

	var lines []string
	currentLine := words[0]

	for _, word := range words[1:] {
		testLine := currentLine + " " + word
		width, _ := pdf.MeasureTextWidth(testLine)

		if width <= maxWidth {
			currentLine = testLine
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	lines = append(lines, currentLine)
	return lines
}
func MakeVatCalculations(pdf *gopdf.GoPdf, summary payload.InvoiceSummary) {
	err := pdf.SetFont("OpenSans", "", 7)
	if err != nil {
		log.Fatal(err)
	}

	x := 203.5
	y := 665.0
	colWidths := []float64{62, 35.7, 62, 62}
	lineHeight := 6.0
	rowspace := 9.0

	AddVatRows(pdf, x, y, colWidths, lineHeight, summary, rowspace)
}

func WrappBoldNormal(pdf *gopdf.GoPdf, runs []TextRun, x, y, maxwidth, lineHeight float64) {
	wordsrun := RunsToWords(pdf, runs)
	wrapruns := WrapRuns(wordsrun, maxwidth)
	DrawWrappedRuns(pdf, x, y, lineHeight, wrapruns)
}
func RunsToWords(pdf *gopdf.GoPdf, runs []TextRun) []WordRun {
	var out []WordRun

	for _, r := range runs {
		pdf.SetFont(r.Font, "", r.Size)
		words := strings.Split(r.Text, " ")

		for _, w := range words {
			word := w + " "
			width, _ := pdf.MeasureTextWidth(word)

			out = append(out, WordRun{
				Font:  r.Font,
				Size:  r.Size,
				Text:  word,
				Width: width,
			})
		}
	}

	return out
}
func WrapRuns(
	words []WordRun,
	maxWidth float64,
) []LineRun {
	var lines []LineRun
	var current LineRun

	for _, w := range words {
		if current.Width+w.Width <= maxWidth {
			current.Words = append(current.Words, w)
			current.Width += w.Width
		} else {
			lines = append(lines, current)
			current = LineRun{
				Words: []WordRun{w},
				Width: w.Width,
			}
		}
	}

	if len(current.Words) > 0 {
		lines = append(lines, current)
	}

	return lines
}
func DrawWrappedRuns(
	pdf *gopdf.GoPdf,
	x, y, lineHeight float64,
	lines []LineRun,
) {
	currY := y

	for _, line := range lines {
		currX := x

		for _, w := range line.Words {
			pdf.SetFont(w.Font, "", w.Size)
			pdf.SetXY(currX, currY)
			pdf.Text(w.Text)

			currX += w.Width
		}

		currY += lineHeight
	}
}
