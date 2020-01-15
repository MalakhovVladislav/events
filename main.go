package main

import (
	"net/http"

	"github.com/antage/eventsource"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()

	api := NewApiHandler(es)
	http.Handle("/polling", es)
	http.HandleFunc("/list", api.GetList)
	http.HandleFunc("/add", api.AddObject)

	http.ListenAndServe(":8080", nil)
}
