package main

import "fmt"

// // https://github.com/adonovan/gopl.io/blob/master/ch2/popcount/main.go
// // https://www.reddit.com/r/golang/comments/4i4ja9/purpose_of_the_init_instructions_in_popcount_the/
// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var inpt uint64
	for _, err := fmt.Scan(&inpt); err == nil; _, err = fmt.Scan(&inpt){
		fmt.Println(PopCount(inpt))
	}
	// fmt.Println(PopCount(10))
}
// go run pop.go