package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/headers", printRequestHeaders)
	http.HandleFunc("/request", sendAuthenticatedRequest)
	http.ListenAndServe(":8080", nil)
}

func printRequestHeaders(w http.ResponseWriter, r *http.Request) {
	printHeaders(r.Header, w)
}

func sendAuthenticatedRequest(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet,
		targetURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Cannot create HTTP request: %v", err)))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Request failed: %v", err)))
		return
	}

	w.Write([]byte(fmt.Sprintf("Status: %d\nHeaders:\n", resp.StatusCode)))
	printHeaders(resp.Header, w)
}

func printHeaders(h http.Header, w http.ResponseWriter) {
	for name, values := range h {
		for _, value := range values {
			w.Write([]byte(fmt.Sprintf("%s = %s\n", name, value)))
		}
	}
}
