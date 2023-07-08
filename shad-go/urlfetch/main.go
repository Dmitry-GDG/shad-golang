//go:build !solution

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}

}

// cd ..
// go run ./urlfetch/main.go "https://google.com" "https://ya.ru"
// go run ./urlfetch/main.go https://google.com https://ya.ru
// go run ./urlfetch/main.go golang.org
// go fmt ./urlfetch/main.go
// go goimports ./urlfetch/main.go // go goimports -w ./urlfetch/main.go
// golangci-lint run ./urlfetch/main.go
// go test ./urlfetch/...
