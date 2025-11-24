package utils

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"os"
)

func HTMLtoPDF(htmlContent string, outputPath string) error {
	// create a new browser context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var pdfBuf []byte

	err := chromedp.Run(ctx,
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
		return err
	}

	return os.WriteFile(outputPath, pdfBuf, 0644)
}
