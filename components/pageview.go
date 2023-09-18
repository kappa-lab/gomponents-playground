package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	t "github.com/maragudk/gomponents/html"
)

type Props struct {
	Title string
	Path  string
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
			t.Script(
				t.Type("text/javascript"),
				g.Raw(`
				console.log('Ready...');
				console.log(location.href);
				function log(val) {
					console.log(val);
				}
				function addMessage(evt) {
					evt.preventDefault()
					var messages = document.getElementById("messageList");
					var li = document.createElement("li")
					li.textContent = evt.srcElement.messageInput.value
					messages.appendChild(li)
					evt.srcElement.messageInput.value = ""
				}
				function clear(evt) {
					evt.preventDefault()
					console.log("clear");
					var messages = document.getElementById("messageList");
					while (messages.firstChild) {
						messages.removeChild(messages.firstChild);
					}
				}

				window.onload = function () {
					console.log('onload');
					const clearBtn = document.getElementById("clearBtn");
					clearBtn.onclick = clear

					const messageInput = document.getElementById("messageForm");
					messageInput.onsubmit = addMessage
				}
				`),
			),
		},
		Body: []g.Node{
			t.StyleAttr("margin:10px"),

			t.Div(
				t.Class("display-2"),
				t.P(g.Text("HelloWorld--------------")),
			),
			t.Div(
				t.FormEl(
					t.ID("messageForm"),
					t.Input(
						t.Type("text"),
						t.ID("messageInput"),
						t.Placeholder("Enter Message"),
						t.AutoFocus(),
					),
				),
			),
			t.Div(
				t.Ul(t.ID("messageList")),
			),
			t.Div(
				t.FormEl(
					t.Button(
						t.ID("clearBtn"),
						g.Raw("Clear")),
				),
			),
		},
	})
}
