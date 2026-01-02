package utils

import (
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
)

func Imageto64(path string) string {
	data, err := os.ReadFile(filepath.Join(path, "static", "images", "logo.png"))
	if err != nil {
		log.Println(err)
	}
	base64image := base64.StdEncoding.EncodeToString(data)
	return base64image
}
