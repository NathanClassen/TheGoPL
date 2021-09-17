// Server2 is an "echo" server with request counter
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
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// hanlder echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // Lock locks m. If the lock is already in use, the calling goroutine blocks until the mutex is available.
	count++
	/* Unlock unlocks m. It is a run-time error if m is not locked on entry to Unlock.
	A locked Mutex is not associated with a particular goroutine. It is allowed for one goroutine to lock a Mutex and then
	arrange for another goroutine to unlock it. */
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// request handler for requests made to baseurl/count
func counter(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter is the interface used by an HTTP handler to construct an HTTP response.
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count) // write to the ResponseWriter to create the response for this request.
	mu.Unlock()
}