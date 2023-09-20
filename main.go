package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/kappa-lab/gomponents-playground/api"
	"github.com/kappa-lab/gomponents-playground/components"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/api") {
		api.Handle(w, r)
		return
	}

	err := components.Page(components.Props{
		Title: "Mypage",
		Req:   r,
	}).Render(w)

	if err != nil {
		panic(err)
	}
}
