package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kappa-lab/gomponents-playground/api"
	"github.com/kappa-lab/gomponents-playground/components"
)

func main() {
	DB_PATH := os.Getenv("DB_PATH")
	if DB_PATH == "" {
		DB_PATH = "127.0.0.1:3306"
	}
	db, err := sql.Open("mysql", "root:root@tcp("+DB_PATH+")/sample")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	cnn, err := db.Conn(ctx)
	if err != nil {
		panic(err)
	}
	defer cnn.Close()

	http.HandleFunc("/", handler)

	http.HandleFunc("/api/customers", func(w http.ResponseWriter, r *http.Request) {
		api.Handle(w, r, cnn)
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
