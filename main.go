package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/healthz" {
		http.Error(w, "wrong handler path", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "bad method", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "ok\n")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "wrong handler path", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "parse form error", http.StatusForbidden)
	}

	name := r.PostFormValue("name")
	if name == "" {
		http.Error(w, "name required", http.StatusForbidden)
	}

	addr := r.PostFormValue("addr")
	if addr == "" {
		http.Error(w, "addr required", http.StatusForbidden)
	}

	fmt.Printf("form params: name=%s, addr=%s\n", name, addr)

	fmt.Fprintf(w, "Hello, %s from %s!\n", name, addr)
}

func main() {
	httpFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", httpFileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/healthz", healthzHandler)
	const port = 8080

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("couldn't start server", err)
	}

	fmt.Printf("Server started at port %d\n", port)
}
