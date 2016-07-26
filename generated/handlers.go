package main

import (
	"net/http"
	"golang.org/x/net/context"
	"encoding/json"
)

var controller Controller

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	input, err := NewGetBookInput(r)
	if err != nil {
		// TODO: Think about this whether this is usually an internal error or it could
		// be from a bad request format...
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = input.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Actually handle the context correctly...
	var ctx context.Context
	resp, err := controller.GetBook(ctx, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
