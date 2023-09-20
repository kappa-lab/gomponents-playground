package main

import (
	"log"
	"net/http"

	"github.com/kappa-lab/gomponents-playground/api"
	"github.com/kappa-lab/gomponents-playground/components"
)

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/api/customers", func(w http.ResponseWriter, r *http.Request) {
		api.Handle(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := components.Page(components.Props{
		Title: "My Shop",
		Req:   r,
	}).Render(w)

	if err != nil {
		panic(err)
	}
}
