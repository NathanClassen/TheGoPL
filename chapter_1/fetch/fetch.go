// Fetch prints the content found at each URL
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		urlStart := time.Now()
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			fmt.Printf("\n%dms elapsed\n", time.Since(start).Milliseconds())
			os.Exit(1)
		}
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fethc: reading %s: %v\n", url, err)
			fmt.Printf("\n%dms elapsed\n", time.Since(start).Milliseconds())
			os.Exit(1)
		}
		ms := time.Since(urlStart).Milliseconds()
		fmt.Printf("%dms %7d %s\n", ms, nbytes, url)
	}
	fmt.Printf("\n%dms elapsed\n", time.Since(start).Milliseconds())
}