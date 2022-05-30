package controllers

import "net/http"

// HelloHandler is example handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, World!"))
}
