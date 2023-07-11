package main

// import "fmt"
// func Reverse(input string) string {
// 	// return ""
// 	r := []rune(input)
// 	rOutp := []rune{}
// 	// fmt.Println("Input runes: ", r)
// 	for i := len(r) - 1; i >= 0; i-- {
// 		rOutp = append(rOutp, r[i])
// 	}
// 	// fmt.Println("Reverse runes: ", rOutp)
// 	outp := string(rOutp)
// 	return (outp)
// }
// func main() {
// 	s := "Привет"
// 	fmt.Println(Reverse(s))
// }

import "fmt"
func CollapseSpaces(input string) string {
	// return ""
	if len(input) <= 1 {
		return input
	}
	r := []rune(input)
	collapseRune := []rune{}
	fmt.Println("Input rune:\t", r)
	tmp := r[0]
	if r[0] != 10 && r[0] != 9 && r[0] != 11 && r[0] != 12 && r[0] != 13 && r[0] != 32 {
		collapseRune = append(collapseRune, tmp)
	} else if (r[0] == 10 || r[0] == 9 || r[0] == 11 || r[0] == 12 || r[0] == 13 || r[0] == 32) {
		collapseRune = append(collapseRune, 32)
		fmt.Println("Here3")
	}
	for i := 1; i < len(r); i++ {
		if r[i] != 10 && r[i] != 9 && r[i] != 11 && r[i] != 12 && r[i] != 13 && r[i] != 32 {
			collapseRune = append(collapseRune, r[i])
			fmt.Println("Here1")
		} else if (r[i] == 10 || r[i] == 9 || r[i] == 11 || r[i] == 12 || r[i] == 13 || r[i] == 32) && (tmp != 32 && tmp != 10 && tmp != 9 && tmp != 11 && tmp != 12 && tmp != 13) {
			collapseRune = append(collapseRune, 32)
			fmt.Println("Here2")
		}
		tmp = r[i]
	}
	fmt.Println("collapse rune:\t", collapseRune)
	outp := string(collapseRune)
	return outp
}
func main() {
	// s := "Привет  Андрей    привет"
	// s := "\n\t\v\r\f*"
	s := "\n*"
	fmt.Println(CollapseSpaces(s))
}