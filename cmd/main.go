package main

import (
	"exoplanet-service/internals/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
