//go:build !solution

package reverse

// import (
// 	"fmt"
// 	"bytes"
// )

func Reverse(input string) string {
	// return ""
	r := []rune(input)
	rOutp := []rune{}
	// fmt.Println("Input runes: ", r)
	for i := len(r) - 1; i >= 0; i-- {
		rOutp = append(rOutp, r[i])
	}
	// fmt.Println("Reverse runes: ", rOutp)
	outp := string(rOutp)
	return (outp)
}

// func main() {
// 	s := "Привет"
// 	fmt.Println(Reverse(s))
// }

// go fmt ./utf8/reverse/reverse.go
// golangci-lint run ./utf8/reverse/...
// go test -v ./utf8/reverse/...
