//go:build !solution

package spacecollapse

func CollapseSpaces(input string) string {
	// return ""
	if len(input) <= 1 {
		return input
	}
	r := []rune(input)
	collapseRune := []rune{}
	tmp := r[0]
	if r[0] != 10 && r[0] != 9 && r[0] != 11 && r[0] != 12 && r[0] != 13 && r[0] != 32 {
		collapseRune = append(collapseRune, tmp)
	} else if r[0] == 10 || r[0] == 9 || r[0] == 11 || r[0] == 12 || r[0] == 13 || r[0] == 32 {
		collapseRune = append(collapseRune, 32)
	}
	for i := 1; i < len(r); i++ {
		if r[i] != 10 && r[i] != 9 && r[i] != 11 && r[i] != 12 && r[i] != 13 && r[i] != 32 {
			collapseRune = append(collapseRune, r[i])
		} else if (r[i] == 10 || r[i] == 9 || r[i] == 11 || r[i] == 12 || r[i] == 13 || r[i] == 32) && (tmp != 32 && tmp != 10 && tmp != 9 && tmp != 11 && tmp != 12 && tmp != 13) {
			collapseRune = append(collapseRune, 32)
		}
		tmp = r[i]
	}
	outp := string(collapseRune)
	return outp
}

// go fmt ./utf8/spacecollapse/collapse.go
// golangci-lint run ./utf8/spacecollapse/...
// go test -v ./utf8/spacecollapse/...
