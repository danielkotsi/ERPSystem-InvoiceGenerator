package utils

import (
	"log"
	"net/http"

	"github.com/go-playground/form/v4"
)

func ParseFormData(r *http.Request, data any) (err error) {
	decoder := form.NewDecoder()
	if err := r.ParseMultipartForm(20000000); err != nil {
		log.Println(err)
		return err
	}

	if err := decoder.Decode(data, r.Form); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
