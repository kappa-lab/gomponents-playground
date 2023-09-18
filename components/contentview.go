package components

import (
	"context"
	"net/http"
	"strings"

	g "github.com/maragudk/gomponents"
	t "github.com/maragudk/gomponents/html"
)

type ContentView struct {
}

func (v ContentView) Node(req *http.Request) g.Node {
	title := strings.TrimPrefix(req.URL.Path, "/")
	var content g.Node
	switch req.URL.Path {
	case "/":
		title = "home"
		content = v.home()
	case "/dashboard":
	case "/products":
		content = v.products()
	case "/orders":
		content = v.orders()
	case "/customers":
		content = v.customers(req.Context())
	}

	return t.Div(
		t.Class("d-flex flex-column align-items-stretch bg-white"),
		t.StyleAttr("margin: 80px"),
		t.H1(
			t.Class("text-uppercase"),
			g.Text(title),
		),
		content,
	)
}

func (ContentView) home() g.Node {
	return t.Div(
		t.P(g.Text("SHOP-ID: 1")),
	)
}

func (ContentView) products() g.Node {
	return t.Div(
		t.Table(
			t.Class("table table-sm table-bordered"),
			t.THead(
				t.Tr(
					t.Th(g.Text("PRODUCT-ID")),
					t.Th(g.Text("NAME")),
					t.Th(g.Text("PRICE")),
				),
			),
			t.TBody(
				t.Tr(
					t.Td(g.Text("1")),
					t.Td(g.Text("bag")),
					t.Td(g.Text("$100")),
				),
			),
			t.TBody(
				t.Tr(
					t.Td(g.Text("2")),
					t.Td(g.Text("pen")),
					t.Td(g.Text("$110")),
				),
			),
		),
	)

}

func (ContentView) orders() g.Node {
	return t.Div(
		t.Table(
			t.Class("table table-sm table-bordered"),
			t.THead(
				t.Tr(
					t.Th(g.Text("ORDER-ID")),
					t.Th(g.Text("PRODUCT-ID")),
					t.Th(g.Text("DATE")),
				),
			),
			t.TBody(
				t.Tr(
					t.Td(g.Text("1001")),
					t.Td(g.Text("1")),
					t.Td(g.Text("2023-01-01 10:00:00")),
				),
			),
			t.TBody(
				t.Tr(
					t.Td(g.Text("1002")),
					t.Td(g.Text("2")),
					t.Td(g.Text("2023-01-01 11:01:01")),
				),
			),
		),
	)

}
func (ContentView) customers(ctx context.Context) g.Node {
	return t.Div()
}
