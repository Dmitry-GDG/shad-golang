//go:build !solution

package varfmt

// see https://coderlessons.com/tutorials/kompiuternoe-programmirovanie/uchebnik-python/47-format-stroki-python

import (
	"fmt"
	"strconv"
	"strings"
)

func removeBraces(strInp []string) []string {
	outpSlice := []string{}
	for _, str := range strInp {
		if str[len(str)-1] == '}' {
			str = str[:len(str)-1]
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
	for i < len(str)-j {
		if str[j+i] > 47 && str[j+i] < 58 {
			i++
		} else {
			break
		}
	}
	if (i < len(str)-j) && i != 0 && str[j+i] == 125 {
		return (j + i + 1)
	}
	return minIndex
}

func Sprintf(format string, args ...interface{}) string {
	inputSlice := strings.Split(format, " ")
	inputSlice = removeBraces(inputSlice)
	sep := ""
	var outpStr string
	for i := 0; i < len(inputSlice); i++ {
		nmb, err := strconv.Atoi(inputSlice[i])
		if err == nil {
			switch v := args[nmb].(type) {
			case int:
				outpStr += sep + strconv.Itoa(v)
			case string:
				outpStr += sep + v
			}
		} else {
			if inputSlice[i] == "" {
				switch v := args[i].(type) {
				case int:
					outpStr += sep + strconv.Itoa(v)
				case string:
					outpStr += sep + v
				}
			} else {
				minIndex := 0
				for j := 0; j < len(inputSlice[i]); {
					if inputSlice[i][j] != 123 {
						j++
					} else {
						outpStr += inputSlice[i][minIndex:j]
						minIndexOld := minIndex
						j++
						minIndex = checkIsFormatInside(j, minIndex, inputSlice[i])
						if minIndex != minIndexOld {
							nextNbrSlice := inputSlice[i][j : minIndex-1]
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
								}
							}
							j = minIndex
						}
					}
				}
				outpStr += inputSlice[i][minIndex:]
			}
		}
		sep = " "
	}

	return outpStr
}

// func Sprintf(format string, args ...interface{}) string {
// 	inputSlice := strings.Split(format, " ")
// 	inputSlice = removeBraces(inputSlice)
// 	sep := ""
// 	next := 0
// 	var outpStr string
// 	for i := 0; i < len(inputSlice); i++ {
// 		if inputSlice[i] == "" {
// 			switch v := args[next].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				outpStr += sep + v
// 			}
// 		} else {
// 			nmb, _ := strconv.Atoi(inputSlice[i])
// 			switch v := args[nmb].(type) {
// 			case int:
// 				outpStr += sep + strconv.Itoa(v)
// 			case string:
// 				outpStr += sep + v
// 			}
// 		}
// 		sep = " "
// 		next++
// 	}
// 	return outpStr
// }

// go fmt ./varfmt/fmt.go
// golangci-lint run ./varfmt/...
// go test -v ./varfmt/...
// go test -v -bench=. ./varfmt/...
