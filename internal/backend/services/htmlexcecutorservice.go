package services

import (
	"html/template"
	"log"
	"net/http"
)

type Excecutor struct {
	Tmpl *template.Template
}

func NewHTMLExcecutor(tmpl *template.Template) *Excecutor {
	return &Excecutor{Tmpl: tmpl}
}

func (a *Excecutor) ServeErrorwithHTML(w http.ResponseWriter, err error, status int) error {
	log.Println(err)
	w.WriteHeader(status)
	_ = a.Tmpl.ExecuteTemplate(w, "error.page.html", err)
	return nil
}
