package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"go-initializr/common"
)

type body struct {
	search string
}

func (app *Config) handlerGetModules(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	param := query.Get("search")

	arr := strings.Split(param, " ")

	modules, err := app.service.GetModules(arr...)

	if err != nil {
		common.ErrorJSON(w, 400, fmt.Sprintf("%v", err))
		return
	}

	common.ResponseJSON(w, 200, modules)
}

func (app *Config) handlerModules(w http.ResponseWriter, r *http.Request) {
	param := r.PostFormValue("search")

	arr := strings.Split(param, " ")

	modules, err := app.service.GetModules(arr...)

	if err != nil {
		common.ErrorJSON(w, 400, fmt.Sprintf("%v", err))
		return
	}

	tmpl, _ := template.New("title").Parse(fmt.Sprintf("<p style='text-align:center'><strong>%s</strong>\n</p>", "New dependencies: "))
	tmpl.Execute(w, nil)

	for _, element := range modules {
		tmpl, _ := template.New("element").Parse(fmt.Sprintf("%s\n", element))
		tmpl.Execute(w, nil)
	}
}
