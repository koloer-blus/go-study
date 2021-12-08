package main

import (
	"fmt"
	"go-study/gee-web/day1-http-base/base3/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.path = %q \n", req.URL.Path)
	})
	r.GET("/hello", func(rw http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
