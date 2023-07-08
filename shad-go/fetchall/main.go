//go:build !solution

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	mapTime := make(map[int]time.Time)
	var maxTime time.Time
	start := time.Now()
	for i, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string, start time.Time, mapTime map[int]time.Time, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			} else {
				res, err1 := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				timeNow := time.Now()
				mapTime[i] = timeNow
				if err1 != nil {
					fmt.Fprintf(os.Stderr, "%.2fs\tGet %s\t%v\n", timeNow.Sub(start).Seconds(), url, err1)
				} else {
					fmt.Printf("%.2fs\t%d\t%v\n", timeNow.Sub(start).Seconds(), len(res), url)
				}
			}
		}(url, start, mapTime, i, wg)
	}
	wg.Wait()
	// time.Sleep(time.Second * 4)
	for _, ttime := range mapTime {
		if ttime.After(maxTime) {
			maxTime = ttime
		}
	}
	fmt.Printf("%.2fs elapsed\n", maxTime.Sub(start).Seconds())
}

// go run main.go https://gopl.io golang.org http://golang.org
// go run main.go https://google.com ya.r https://ya.ru https://google.com
// go fmt main.go
// goimports -w main.go
// cd ..
// golangci-lint run fetchall/...
