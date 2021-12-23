package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, err := fmt.Fprintf(w, "Count %d\n", count)
	if err != nil {
		return
	}
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err3 := fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	if err3 != nil {
		return
	}
	for k, v := range r.Header {
		_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
	_, err := fmt.Fprintf(w, "Host = %q\n", r.Host)
	if err != nil {
		return
	}
	_, err1 := fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	if err1 != nil {
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	_, err2 := fmt.Fprintf(w, "url.Path = %q\n", r.URL.Path)
	if err2 != nil {
		return
	}
}
