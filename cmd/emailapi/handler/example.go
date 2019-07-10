package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ExampletGetEndpoint will return the welcome message for test in JSON as an example.
func (h *Handler) ExampletGetEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test GET endpoint is being hit now!")

	// Call to some function to do some work.
	// If you get an error back from a function you can respond with a header
	// of some error to inform the user, e.g.
	// w.WriteHeader(http.StatusInternalServerError)

	// Build response if it was not received fom function.
	m := make(map[string]string)
	m["endpoint"] = "/v1/test"
	m["result"] = "You have hit the GET test endpoint and it has answered!"

	// Marshal response and write it back.
	w.WriteHeader(http.StatusOK)
	respondJSON, _ := json.Marshal(m)
	w.Write(respondJSON)
}

type testSubmit struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

// ExamplePostEndpoint will return the welcome message for test in JSON as an example.
// A mock of the payload for this function is below, it could be adjusted to any other
// POST payload format that you would like.
/*
	{
		"sender": "mmcken3",
		"message": "this is my test message"
	}
*/
func (h *Handler) ExamplePostEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test POST endpoint is being hit now!")

	// Build response if it was not received fom function.
	m := make(map[string]string)
	m["endpoint"] = "/v1/test/submit"
	m["result"] = "You have hit the GET test endpoint and it has answered!"

	var ex testSubmit
	err := json.NewDecoder(r.Body).Decode(&ex)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m["result"] = "error decoding json payload"
		log.Printf("error decoding json payload")

		// Marshal response and write it back.
		respondJSON, _ := json.Marshal(m)
		w.Write(respondJSON)
		return
	}

	// Call to some function to do some work.
	// If you get an error back from a function you can respond with a header
	// of some error to inform the user, e.g.
	// w.WriteHeader(http.StatusInternalServerError)

	m["sender"] = ex.Sender
	m["messageEcho"] = ex.Message

	// Marshal response and write it back.
	respondJSON, _ := json.Marshal(m)
	w.Write(respondJSON)
}
