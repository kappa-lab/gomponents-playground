package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request, cnn *sql.Conn) {
	data, err := getCustomers(r.Context(), cnn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{Data: data})
}

type Response struct {
	Data []Customer
}

type Customer struct {
	ID   int
	Name string
}

func getCustomers(ctx context.Context, cnn *sql.Conn) ([]Customer, error) {
	r, err := cnn.QueryContext(ctx, "select * from customers")
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := []Customer{}
	for r.Next() {
		var (
			id   int
			name string
		)
		err := r.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		res = append(res, Customer{ID: id, Name: name})
	}
	return res, nil
}
