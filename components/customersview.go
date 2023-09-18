package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	g "github.com/maragudk/gomponents"
	t "github.com/maragudk/gomponents/html"
)

type CustomersJSON struct {
	Data []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

type CustomersView struct {
	httpCli *http.Client
}

func NewCustomersView() *CustomersView {
	v := &CustomersView{
		httpCli: http.DefaultClient,
	}
	v.httpCli.Timeout = time.Second
	return v
}

func (v *CustomersView) Load(ctx context.Context) g.Node {
	endpoint := "http://localhost:8080/api/customers"
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return t.Div(t.P(g.Text(err.Error())))
	}
	req = req.WithContext(ctx)
	resp, err := v.httpCli.Do(req)
	if err != nil {
		return t.Div(t.P(g.Text(err.Error())))
	}
	var data CustomersJSON
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return t.Div(t.P(g.Text(err.Error())))
	}

	return v.buildNode(data)
}

func (v *CustomersView) buildNode(data CustomersJSON) g.Node {
	tableContents := []g.Node{
		t.Class("table table-sm table-bordered"),
		t.THead(
			t.Tr(
				t.Th(g.Text("USER-ID")),
				t.Th(g.Text("NAME")),
			),
		),
	}

	for _, v := range data.Data {
		tableContents = append(tableContents,
			t.TBody(
				t.Tr(
					t.Td(g.Text(fmt.Sprint(v.ID))),
					t.Td(g.Text(v.Name)),
				),
			),
		)
	}

	return t.Div(t.Table(tableContents...))
}
