package main

import "net/http"

func main() {
	// Start the server
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	http.ListenAndServe(":8080", mux)
}
