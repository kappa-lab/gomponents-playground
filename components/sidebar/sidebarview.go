package sidebar

import (
	"net/http"
	"strings"

	g "github.com/maragudk/gomponents"
	t "github.com/maragudk/gomponents/html"
)

type Item struct {
	Title  string
	Path   string
	Active bool
}

func CreateItems(r *http.Request) []Item {
	items := []Item{
		{
			Title: "Home",
			Path:  "/",
		},
		{
			Title: "Dashboard",
			Path:  "/dashboard",
		},
		{
			Title: "Orders",
			Path:  "/orders",
		},
		{
			Title: "Products",
			Path:  "/products",
		},
		{
			Title: "Customers",
			Path:  "/customers",
		},
	}

	if r.URL.Path == "/" {
		items[0].Active = true
	} else {
		for i, v := range items[1:] {
			if strings.HasPrefix(r.URL.Path, v.Path) {
				items[i+1].Active = true
				break
			}
		}

	}

	return items
}

func Sidebar(items []Item) g.Node {
	return t.Div(
		t.Class("d-flex flex-column p-3 text-white bg-dark"),
		t.StyleAttr("width: 280px; height: 100vh;"),
		t.A(
			t.Href("/"),
			t.Class("d-flex align-items-center mb-3 mb-md-0 me-md-auto text-white text-decoration-none"),
			t.Span(t.Class("fs-4"), g.Text("My Shop")),
		),
		t.Hr(),
		buildItemNodes(items),
		t.Hr(),
		t.Div(
			t.Class("text-center small"),
			t.P(g.Text("ver 0.0.1")),
		),
	)
}

func buildItemNodes(items []Item) g.Node {
	contents := []g.Node{
		t.Class("nav nav-pills flex-column mb-auto"),
	}

	for _, v := range items {
		state := "text-white"
		if v.Active {
			state = "active"
		}
		contents = append(contents, t.Li(
			t.Class("nav-item"),
			t.A(
				t.Class("nav-link "+state),
				t.Href(v.Path),
				g.Text(v.Title),
			),
		))
	}

	return t.Ul(contents...)
}
