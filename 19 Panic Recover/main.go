package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered in: %v", r)
				debug.PrintStack()
				http.Error(w, "Something went wrong!", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Panic!")
	})

	http.ListenAndServe(":8080", recoverMiddleware(mux))
}
