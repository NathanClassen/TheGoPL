package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	count++
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto) // GET /counter HTTP/1.1
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }

// request handler for requests made to baseurl/count
func counter(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter is the interface used by an HTTP handler to construct an HTTP response.
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count) // write to the ResponseWriter to create the response for this request.
	mu.Unlock()
}