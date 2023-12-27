package main

import (
	"net/http"

	"go-initializr/common"
)

type health struct {
	Status string `json:"status"`
}

func (*Config) handlerHealth(w http.ResponseWriter, r *http.Request) {
	common.ResponseJSON(w, 200, health{
		Status: "UP",
	})
}
