package main

import (
	"fmt"
	_ "embed"
)

//go:embed hello.txt
var s string
//var s []byte

func main() {
	fmt.Println(s)
}