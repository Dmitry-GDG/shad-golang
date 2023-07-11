//go:build !solution

package reverse

func Reverse(input string) string {
	r := []rune(input)
	rOutp := []rune{}
	for i := len(r) - 1; i >= 0; i-- {
		rOutp = append(rOutp, r[i])
	}
	outp := string(rOutp)
	return (outp)
}

// go fmt ./utf8/reverse/reverse.go
// golangci-lint run ./utf8/reverse/...
// go test -v ./utf8/reverse/...
