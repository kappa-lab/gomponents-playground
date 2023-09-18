//go:build go1.18
// +build go1.18

package main

import (
	"log"
	"net/http"

	"github.com/kappa-lab/gomponents-playground/components"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", http.HandlerFunc(handler)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := components.Page(components.Props{
		Title: "Mypage",
		Path:  r.URL.Path,
	}).Render(w)

	if err != nil {
		panic(err)
	}
}
