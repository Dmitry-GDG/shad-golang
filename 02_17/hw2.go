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

//--------------------------------------------

// import "fmt"
// func CollapseSpaces(input string) string {
// 	// return ""
// 	if len(input) <= 1 {
// 		return input
// 	}
// 	r := []rune(input)
// 	collapseRune := []rune{}
// 	fmt.Println("Input rune:\t", r)
// 	tmp := r[0]
// 	if r[0] != 10 && r[0] != 9 && r[0] != 11 && r[0] != 12 && r[0] != 13 && r[0] != 32 {
// 		collapseRune = append(collapseRune, tmp)
// 	} else if (r[0] == 10 || r[0] == 9 || r[0] == 11 || r[0] == 12 || r[0] == 13 || r[0] == 32) {
// 		collapseRune = append(collapseRune, 32)
// 		fmt.Println("Here3")
// 	}
// 	for i := 1; i < len(r); i++ {
// 		if r[i] != 10 && r[i] != 9 && r[i] != 11 && r[i] != 12 && r[i] != 13 && r[i] != 32 {
// 			collapseRune = append(collapseRune, r[i])
// 			fmt.Println("Here1")
// 		} else if (r[i] == 10 || r[i] == 9 || r[i] == 11 || r[i] == 12 || r[i] == 13 || r[i] == 32) && (tmp != 32 && tmp != 10 && tmp != 9 && tmp != 11 && tmp != 12 && tmp != 13) {
// 			collapseRune = append(collapseRune, 32)
// 			fmt.Println("Here2")
// 		}
// 		tmp = r[i]
// 	}
// 	fmt.Println("collapse rune:\t", collapseRune)
// 	outp := string(collapseRune)
// 	return outp
// }
// func main() {
// 	// s := "Привет  Андрей    привет"
// 	// s := "\n\t\v\r\f*"
// 	s := "\n*"
// 	fmt.Println(CollapseSpaces(s))
// }

//--------------------------------------------

import (
	"fmt"
	"strings"
	"strconv"
)
// func removeBraces(strInp []string) []string {
// 	outpSlice := []string{}
// 	for _, str := range strInp {
// 		if str[len(str) - 1] == '}' {
// 			str = str[:len(str) - 1]
// 		}
// 		if str[0] == '{' {
// 			str = str[1:]
// 		}
// 		outpSlice = append(outpSlice, fmt.Sprint(str))
// 	}
// 	return outpSlice
// }

func removeBraces(strInp []string) []string {
	outpSlice := []string{}
	for _, str := range strInp {
		if str[len(str) - 1] == '}' {
			str = str[:len(str) - 1]
		}
		if str[0] == '{' {
			str = str[1:]
		}
		outpSlice = append(outpSlice, fmt.Sprint(str))
	}
	return outpSlice
}


func checkIsFormatInside(j, minIndex int, str string) int {
	i := 0
	for ; i < len(str) - j; {
		if (str[j + i] > 47 && str[j + i] < 58) {
			i++
		} else {
			break
		}
	}
	if (i < len(str) - j) && i != 0 && str[j + i] == 125 {
		return(j + i + 1)
	}
	return minIndex
}

func Sprintf(format string, args ...interface{}) string {
	inputSlice := strings.Split(format, " ")
	inputSlice = removeBraces(inputSlice)
	fmt.Println(inputSlice)
	sep := ""
	// next := 0
	var outpStr string
	for i := 0; i < len(inputSlice); i++ {
		// nmb, _ := strconv.Atoi(inputSlice[i])
		nmb, err := strconv.Atoi(inputSlice[i])
		if err == nil {
			// outpStr += sep + args[nmb]
			switch v := args[nmb].(type) {
			case int:
				outpStr += sep + strconv.Itoa(v)
			case string:
				outpStr += sep + v
			} 
		} else {
			if inputSlice[i] == "" {
				// fmt.Print("Here - 1")
				switch v := args[i].(type) {
				case int:
					outpStr += sep + strconv.Itoa(v)
				case string:
					outpStr += sep + v
			}
			// fmt.Println(" next = ", next)
			} else {
				// strTmp := []string
				minIndex := 0
				for j := 0; j < len(inputSlice[i]); {
					fmt.Println("minindex = ", minIndex, " j = ", j)
					if inputSlice[i][j] != 123 {
						j++
					} else {
						outpStr += inputSlice[i][minIndex:j]
						// newFormat := []string{}
						minIndexOld := minIndex
						j++
						minIndex = checkIsFormatInside(j, minIndex, inputSlice[i])
						if minIndex != minIndexOld {
							nextNbrSlice := inputSlice[i][j:minIndex - 1]
							// fmt.Println(nextNbrSlice)
							if nextNbrSlice == "" {
								switch v := args[i].(type) {
								case int:
									outpStr += strconv.Itoa(v)
								case string:
									outpStr += v
								}
							} else {
								nextNmbr, _ := strconv.Atoi(nextNbrSlice)
								switch v := args[nextNmbr].(type) {
								case int:
									outpStr += strconv.Itoa(v)
								case string:
									outpStr += v
									fmt.Println("Here")
								}
							}
							j = minIndex
						}
					}
				}
				outpStr += inputSlice[i][minIndex:]

			// 	// fmt.Print("Here - 2")
				// outpStr += sep + inputSlice[i]
			}
			// fmt.Print(" next = ", next, " nmb = ", nmb)
		}
		sep = " "
		// next++
	}

	return outpStr
}

// func Sprintf(format string, args ...interface{}) string {
// 	inputSlice := strings.Split(format, " ")
// 	inputSlice = removeBraces(inputSlice)
// 	fmt.Println(inputSlice)
// 	sep := ""
// 	next := 0
// 	var outpStr string
// 	for i := 0; i < len(inputSlice); i++ {
// 		if inputSlice[i] == "" {
// 			// fmt.Print("Here - 1")
// 			switch v := args[next].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				outpStr += sep + v
// 			}
// 			// fmt.Println(" next = ", next)
// 		} else {
// 			// fmt.Print("Here - 2")
// 			nmb, _ := strconv.Atoi(inputSlice[i])
// 			switch v := args[nmb].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				outpStr += sep + v
// 			}
// 			// fmt.Print(" next = ", next, " nmb = ", nmb)
// 		}
// 		sep = " "
// 		next++
// 	}

// 	return outpStr
// }

// func Sprintf(format string, args ...interface{}) string {
// 	// return ""
// 	// fmt.Println(format)
// 	inputSlice := strings.Split(format, " ")
// 	// fmt.Println(inputSlice)
// 	inputSlice = removeBraces(inputSlice)
// 	// fmt.Println(inputSlice)
// 	// for _, elem := range args {
// 	// 	fmt.Println("1", elem)
// 	// }
// 	// for _, elem := range args {
// 	// 	fmt.Println("2", elem)
// 	// }
// 	// fmt.Println(args[2], args[0], args[1])

// 	// fmt.Printf("%T, %T\n", inputSlice, inputSlice[0])
// 	// fmt.Printf("%T, %T\n", args[0], args[1])
	
// 	// for i:= 0; i < len(inputSlice); i++ {
// 	// 	fmt.Println("i = ", i, "inputSlice[i] =", inputSlice[i])
// 	// }

// 	sep := ""
// 	next := 0
// 	var outpStr string
// 	for i := 0; i < len(inputSlice); i++ {
// 		if inputSlice[i] == "" {
// 			fmt.Print("Here - 1")
// 			// tmp := args[i]
// 			switch v := args[next].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				outpStr += sep + v
// 			}
// 			// outpStr += sep + str(args[i])
// 			// fmt.Printf("type arg: %T", args[i])
// 			if next < i {
// 				next = i
// 				fmt.Println("next0:, ", next)
// 			}
// 			next++
// 			fmt.Println(" next = ", next)
// 		} else {
// 			fmt.Print("Here - 2")
// 			nmb, _ := strconv.Atoi(inputSlice[i])
// 			switch v := args[nmb].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				// nmb, _ := strconv.Atoi(v)
// 				outpStr += sep + v
// 			}
// 			// tmp := args[nmb]
// 			// switch v := args[nmb].(type) {
// 			// case int:
// 			// 	outpStr += sep + strconv.Itoa(v)
// 			// case string:
// 			// 	outpStr += sep + v
// 			// }
// 			// outpStr += sep + args[nmb]
// 			// fmt.Printf("type arg: %T", args[nmb])
// 			// fmt.Println("Here")
// 			fmt.Print(" next = ", next, " nmb = ", nmb)
// 			if nmb >= next {
// 				next = nmb + 1
// 			}
// 			fmt.Println(" next = ", next)
// 		}
// 		sep = " "
// 	}

// 	return outpStr
// 	// return ""
// }
func main() {
	// Sprintf("{1} {}", "Hello", "World")
	// Sprintf("{} {} {2}", 1000, 1001, 1002)
	// fmt.Println(fmt.Sprintf("{1} {}", 42, 43))
	// fmt.Println(Sprintf("{1} {} {3} {2} {} {} {} {}", 1000, 1001, 1002, 1003, 4, 5, 6, 7, 8))
	// fmt.Println(Sprintf("{} {0} {0} {0} {}", 0, 1, 2, 3, 4))
	// fmt.Println(Sprintf("Hello, {2}", 0, 1, "World"))
	fmt.Println(Sprintf("He{2}o{2}o", 0, 1, "ll"))
}

//--------------------------------------------