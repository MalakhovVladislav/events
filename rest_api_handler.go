package main

import (
	"encoding/json"
	"net/http"

	"github.com/antage/eventsource"
)

type ApiHandler struct {
	eventEmmiter eventsource.EventSource
	objects      []*Object
}

func NewApiHandler(eventEmmiter eventsource.EventSource) *ApiHandler {
	return &ApiHandler{
		eventEmmiter: eventEmmiter,
		objects: []*Object{
			{
				ID:   "1",
				Data: "1Data",
			},
			{
				ID:   "2",
				Data: "2Data",
			},
		},
	}
}

func (h *ApiHandler) GetList(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(h.objects)
}

func (h *ApiHandler) AddObject(w http.ResponseWriter, r *http.Request) {

	h.objects = append(h.objects, &Object{
		ID:   "3",
		Data: "3Data",
	})
	j, _ := json.Marshal(h.objects)

	h.eventEmmiter.SendEventMessage(string(j), "list", "id")
}
