package main

// // sqrt
// import (
// 	"fmt"
// )
// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	n := 1.0
// 	for i := 0; i < 16; i++ {
// 		for z * z < x {
// 			z += n
// 		}
// 		z -= n
// 		n /= 10
// 	}
// 	return z
// }
// func main() {
// 	var a float64
// 	fmt.Scan(&a)
// 	fmt.Println(Sqrt(a))
// }

// import "fmt"
// func main() {
// 	s := []int{2, 3, 5, 7, 11, 13}
// 	printSlice(s)
// 	// Slice the slice to give it zero length.
// 	s = s[:0]
// 	printSlice(s)
// 	// Extend its length.
// 	s = s[:4]
// 	printSlice(s)
// 	// Drop its first two values.
// 	s = s[2:]
// 	printSlice(s)
// }
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

// import "fmt"
// func main() {
// 	a := make([]int, 5)
// 	printSlice("a", a)
// 	b := make([]int, 0, 5)
// 	printSlice("b", b)
// 	c := b[:2]
// 	printSlice("c", c)
// 	d := c[2:5]
// 	printSlice("d", d)
// 	b = b[:5]
// 	d[1] = 100
// 	printSlice("b", b)
// }
// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// // Implement Pic
// // !!! ТОЛЬКО В ПЛЭЙГРАУНДЕ ЗАПУСТИТСЯ!!!!!
// // It should return a slice of length dy, each element of which
// // is a slice of dx 8-bit unsigned integers. When you run the program,
// // it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
// // The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
// // (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// // (Use uint8(intValue) to convert between types.)
// import "golang.org/x/tour/pic"
// func Pic(dx, dy int) [][]uint8 {
//     p := make([][]uint8, dy)
//     for y := range p {
//         p[y] = make([]uint8, dx)
//         for x := range p[y] {
// 			p[y][x] = uint8(x^y)
// 			// p[y][x] = uint8(x*y)
// 			// p[y][x] = uint8((x+y)/2)
// 			// p[y][x] = uint8(x*x + y*y)
// 			// p[y][x] = uint8(y * 10000 / (x + 1))
//         }
//     }
//     return p
// }
// func main() {
// 	// pic.
// 	io.WriteString(Show(Pic), "\n")
// }


// // Exercise: Maps
// // Implement WordCount. It should return a map of the counts of each “word” in the string s. 
// // The wc.Test function runs a test suite against the provided function 
// // and prints success or failure.
// // You might find strings.Fields helpful.
// import (
// 	// "golang.org/x/tour/wc"
// 	"strings"
// 	"fmt"
// )
// func WordCount(s string) map[string]int {
// 	m := make(map[string]int)
// 	ss := strings.Split(s, " ")
// 	for i := 0; i < len(ss); i++ {
// 		// fmt.Println(ss[i])
// 		_, ok := m[ss[i]]
// 		if ok {
// 			m[ss[i]]++
// 		} else {
// 			m[ss[i]] = 1
// 		}
// 	}
// 	return m
// }
// func main() {
// 	// wc.Test(WordCount)
// 	var s, sep, tmp string
// 	for i := 0; i < 6; i++ {
// 		fmt.Scan(&tmp)
// 		s += sep + tmp
// 		sep = " "
// 	}
// 	fmt.Println(WordCount(s))
// }

// // Function closures
// // Go functions may be closures. A closure is a function value 
// // that references variables from outside its body. 
// // The function may access and assign to the referenced variables; 
// // in this sense the function is "bound" to the variables.
// // For example, the adder function returns a closure. Each closure is bound to its own sum variable.
// import "fmt"
// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }
// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(pos(i), neg(-2 * i))
// 	}
// }

// Exercise: Fibonacci closure
// Let's have some fun with functions.
// Implement a fibonacci function that returns a function (a closure) 
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
import "fmt"
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	iPrev := 0
	iNow := 0
	var tmp int
	return func() int {
		if iNow == 0 {
			iPrev = 0
			iNow = 1
			return iPrev
		} else if iNow == 1 && iPrev ==0 {
			iPrev = 1
			iNow = 1
			return iPrev
		}
		tmp = iPrev
		iPrev = iNow
		iNow += tmp
		return iPrev
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
