package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", verbose(http.FileServer(http.Dir("./www"))))
	http.ListenAndServe(":8000", nil)
}

func verbose(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
