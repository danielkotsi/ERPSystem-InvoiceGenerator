package utils

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"net/url"
	"sync"
	"time"
)

var (
	allocCtx  context.Context
	allocOnce sync.Once
)

// initialize Chrome ONCE
func initChrome() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("disable-default-apps", true),
	)

	allocCtx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
}

func HTMLtoPDF2(htmlContent string) ([]byte, error) {
	allocOnce.Do(initChrome)

	// new TAB, not new browser
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// hard timeout so Chrome can't stall
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var pdfBuf []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+url.PathEscape(htmlContent)),

		// wait until layout is ready (not visible!)
		chromedp.WaitReady("body"),

		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				Do(ctx)
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
