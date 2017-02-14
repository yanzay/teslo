package main

import (
	"net/http"

	"github.com/yanzay/teslo"
	"github.com/yanzay/teslo/cmd/templates"
)

func main() {
	server := teslo.NewServer()
	items := []*templates.Item{
		{ID: "1", Content: "write a todo app"},
		{ID: "2", Content: "extract framework", Done: true},
		{ID: "3", Content: "write a shop app"},
	}
	server.Render = func(w http.ResponseWriter) {
		templates.WriteHello(w, items)
	}
	server.Subscribe("todo", func(event *teslo.Event) {
		if event.Type == "click" {
			for _, item := range items {
				if event.ID == item.ID {
					item.Done = !item.Done
				}
			}
			server.Respond("todo", templates.Todo(items))
		}
	})
	server.Start()
}
