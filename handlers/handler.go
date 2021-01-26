// Package handlers defines the HTTP handlers for this GraphQL API.
package handlers

import (
	"bytes"
	"fmt"
	"net/http"
)

// respond
func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

// isSupported determines if a request method is supported
func isSupported(method string) bool {
	return method == "POST" || method == "GET"
}

// errorJSON
func errorJSON(msg string) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `{"error": "%s"}`, msg)
	return buf.Bytes()
}
