package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/deny" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "welcome home %s!, with user %s and %s", getEnv("HOSTNAME", "default_hostname"), getEnv("DB_USER", "default_user"), getEnv("DB_URL", "default_url"))
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func wrapHandlerWithLogging(wrapperHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("--> %s %s", req.Method, req.URL.Path)
		wrapperHandler.ServeHTTP(w, req)
	})
}
func main() {
	//	router := mux.NewRouter().StrictSlash(true)
	rootHandler := wrapHandlerWithLogging(http.HandlerFunc(homeLink))
	http.Handle("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
