package server

import (
	"net/http"
)

func StartServer() error {
	// Serve static files (Monaco Editor assets)
	http.Handle("/", http.FileServer(http.Dir("assets/monaco/")))

	// Start server on port 8080
	return http.ListenAndServe(":8080", nil)
}
