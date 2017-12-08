package main

import (
    "net/http"
    "os"
    "github.com/gorilla/handlers"
)

// env returns the value for an environment variable or, if not set, a fallback
// value.
func env(key, fallback string) string {
	if value := os.Getenv(key); 0 < len(value) {
		return value
	}
	return fallback
}

func main() {
   // Collect environment variables.
   folder := env("FOLDER", "/opt") + "/"
   port := env("PORT", "8080")
   http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(folder))))
   http.ListenAndServe(":"+port, nil)
}
