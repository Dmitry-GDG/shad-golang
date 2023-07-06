//go:build !solution

package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
)
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
func readFile(fileName string) []string {
	var fileInside []string
	f, err := os.Open(fileName)
    checkErr(err)
	defer f.Close()
	rd := bufio.NewReader(f)
	for i := 0; ; i++ {
		tmp, err := rd.ReadString('\n')
		if err != io.EOF {
			checkErr(err)
		}
		tmp = strings.TrimSuffix(tmp, "\n")
		tmp = strings.TrimSuffix(tmp, " ")
		// if tmp != "" {							// Горе от ума: удалил пробелы, из-за этого не проходил тест
			// for i := 0; tmp[0] == ' '; i++ {
			// 	tmp = tmp[1:]
			// }
			// for i := 0; tmp[len(tmp) - 1] == ' '; i++ {
			// 	tmp = tmp[:len(tmp) - 1]
			// }
			// if tmp != "" {
				fileInside = append(fileInside, tmp)
			// }
		// }
		if err == io.EOF {
			break
		}
	}
	return fileInside
}
func main() {
	var filesInside []([]string)
	result := make(map[string]int)
	files := os.Args[1:]
	// fmt.Println(files)
	for i := 0; i < len(files); i++ {
		filesInside = append(filesInside, readFile(files[i]))
		// fmt.Println(filesInside[i])
	}
	// fmt.Println(filesInside)
	for _, file := range filesInside {
		for _, str := range file {
			_, ok := result[str]
			if ok {
				result[str]++
			} else {
				result[str] = 1
			}
		}
	}
	// fmt.Println(result)
	for str, sum := range result {
		if sum > 1 {
			fmt.Printf("%d\t%s\n", sum, str)
		}
	}
}
