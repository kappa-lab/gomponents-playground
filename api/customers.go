package api

import (
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{
    "data": [
        {
            "id": 1,
            "name": "Ben White"
        },
        {
            "id": 2,
            "name": "Pater Rice"
        },
        {
            "id": 3,
            "name": "Bob Marry"
        }
    ]
}`))

}
