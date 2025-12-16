package utils

import (
	"fmt"
	// "log"
	"net/http"
)

// JsonResponse sends a JSON response with the given data and status code
// This utility function standardizes JSON responses across all handlers
func PDFResponse(w http.ResponseWriter, data any, status int) {
	pdfBytes := data.([]byte)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf") // inline = open in browser
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Printf("Failed to marshal JSON response: %v", err)
	// 	// Fallback to a simple error response
	// 	http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	// 	return
	// }
	//
	_, err := w.Write(pdfBytes)
	if err != nil {
		fmt.Println("Failed to write PDF response:", err)
	}
}
