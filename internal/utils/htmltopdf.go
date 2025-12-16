package utils

import (
	"context"
	"encoding/base64"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/skip2/go-qrcode"
)

func HTMLtoPDF(htmlContent string) (pdfinbytes []byte, err error) {
	// create a new browser context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte

	err = chromedp.Run(ctx,
		// Navigate to the HTML using a data URL
		chromedp.Navigate("data:text/html,"+htmlContent),
		chromedp.WaitReady("body"), // wait until the body is loaded
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Use cdproto page.PrintToPDF
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			pdfBuf = buf
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return pdfBuf, nil
}

func GenerateQRcodeBase64(qrURL string) (QRbase64 string, err error) {
	qrpng, err := qrcode.Encode(qrURL, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	QRbase64 = base64.StdEncoding.EncodeToString(qrpng)
	return QRbase64, nil
}
