package main

import (
	"net/http"
	"text/template"
)

func (*Config) handlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}
