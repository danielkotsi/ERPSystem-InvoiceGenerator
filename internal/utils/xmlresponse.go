package utils

import (
	"log"
	"net/http"
)

// JsonResponse sends a JSON response with the given data and status code
// This utility function standardizes JSON responses across all handlers
func XMLResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)

	xml := data.([]byte)
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	log.Printf("Failed to marshal JSON response: %v", err)
	// 	// Fallback to a simple error response
	// 	http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
	// 	return
	// }
	//
	if _, err := w.Write(xml); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}
