package components

import (
	"net/http"

	"github.com/kappa-lab/gomponents-playground/components/sidebar"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	t "github.com/maragudk/gomponents/html"
)

type Props struct {
	Title string
	Req   *http.Request
}

// Page is a whole document to output.
func Page(p Props) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: p.Title,
		Head: []g.Node{
			t.Link(
				t.Rel("stylesheet"),
				t.Href("https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css"),
			),
			t.Link(
				t.Rel("stylesheet"),
				t.Href("https://getbootstrap.jp/docs/5.0/examples/sidebars/sidebars.css"),
			),
			t.Script(
				t.Type("text/javascript"),
				g.Raw(`
				console.log('Ready...');
				`),
			),
		},
		Body: []g.Node{
			t.Main(
				sidebar.Sidebar(sidebar.CreateItems(p.Req)),
				ContentView{}.Node(p.Req),
			),
		},
	})
}
