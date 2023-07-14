//go:build !solution

package main

import (
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"strings"
)

// type Evaluator struct {
// }

// // NewEvaluator creates evaluator.
// func NewEvaluator() *Evaluator {
// 	return &Evaluator{}
// }

// // Process evaluates sequence of words or definition.
// //
// // Returns resulting stack state and an error.
// func (e *Evaluator) Process(row string) ([]int, error) {
// 	return nil, nil
// }

type Stack []int
// func (s *Stack) len() int {
// 	return len(*s)
// }
func (s *Stack) push(x int) {
	*s = append(*s, x)
}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	last := (*s)[len(*s) - 1]
	*s = (*s)[:(len(*s) - 1)]
	return last, nil
}

type KnownWord struct {
	word string
	function func() error
	qtyArgsNeed int
	qtyStackNeed int
	qtyArgsGave int
}

type Evaluator struct {
	stack Stack
	dictionary []KnownWord
	err error
	counter int
}
func (e *Evaluator) mathOperation(mathOp func(int, int) int) func() error {
	return func() error {
			a, err := e.stack.pop()
		if err != nil {
			return err
		}
		b, err := e.stack.pop()
		if err != nil {
			return err
		}
		e.stack.push(mathOp(a, b))
		return nil
	}
}
func (e *Evaluator) plus() error {
	return e.mathOperation(func(a int, b int) int {return a + b})()
}
func (e *Evaluator) minus() error {
	return e.mathOperation(func(a int, b int) int {return a - b})()
}
func (e *Evaluator) multiply() error {
	return e.mathOperation(func(a int, b int) int {return a * b})()
}
func (e *Evaluator) divide() error {
	return e.mathOperation(func(a int, b int) int {return b / a})()
}
func (e *Evaluator) dup() error {
	a, err := e.stack.pop()
	if err != nil {
		return err
	}
	e.stack.push(a)
	e.stack.push(a)
	return nil
}
func (e *Evaluator) over() error {
	a, err := e.stack.pop()
	if err != nil {
		return err
	}
	b, nil := e.stack.pop()
	if err != nil {
		return err
	}
	e.stack.push(b)
	e.stack.push(a)
	e.stack.push(b)
	return nil
}
func (e *Evaluator) drop() error {
	_, err := e.stack.pop()
	return err
}
func (e *Evaluator) swap() error {
	a, err := e.stack.pop()
	if err != nil {
		return err
	}
	b, nil := e.stack.pop()
	if err != nil {
		return err
	}
	e.stack.push(a)
	e.stack.push(b)
	return nil
}
func (e *Evaluator) splitInputData(s string) [][]string {
	sSlice := [][]string{}
	begin := 0
	end := 32
	for i := 0; i < len(s); i++{
		// fmt.Println(s[i])
		if s[i] == 32 {
			i++
		}
		if s[i] == ':' {
			end = 59
			// fmt.Println("Here")
			i++
			if s[i] == 32 {
				i++
			}
		}
		begin = i
		for ; (i < len(s)) && (s[i] != byte(end)); i++{}
		tokenSlice := strings.Split(strings.TrimSpace(string(s[begin:i])), " ")
		// fmt.Printf("11 %v\n", tokenSlice)
		sSlice = append(sSlice, tokenSlice)
		end = 32
	}
	return sSlice
}


// tmpStr := ""
// for ; s[i] != 32; i++ {
// 	tmpStr += string(s[i])
// }
// tmpVal := strings.ToLower(tmpStr)
// if tmpVal == "dup" || tmpVal == "over" || tmpVal == "drop" || tmpVal == "swap" {

// }
// _, err := strconv.Atoi(tmpStr)
// if err != nil {
// 	outpSlice = outpSlice[:0]
// 	return outpSlice, err
// }
// tmpStr = ""
// for ; s[i] != 32; i++ {
// 	tmpStr += string(s[i])
// }

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(s string) ([]int, error) {
	// counter := 0
	var outpErr error
	outpErr = nil
	// var a, b int
	// return nil, nil
	// outpSlice := []int{}
	s = strings.TrimSpace(s)
	// fmt.Println(s)
	sSlice := e.splitInputData(s)
	// fmt.Println(sSlice)
	// for i, val := range sSlice {
	// 	fmt.Printf("len(sSlice[%d]) = %d\n", i, len(val))
	// }
	// dict := e.getDictionary
	// for i := 0; i < len(*dict); i++ {
	// 	fmt.Println((*dict)[i])
	// 	counter++
	// }
	// for _, en := range e.dictionary {
	// 	fmt.Println(en.word)
	// }
	// fmt.Println((e).dictionary)

	for _, val := range sSlice {
		// fmt.Println(val)
		if len(val) == 1 { // known words or data
			sign := false
			for _, entry := range e.dictionary {
				if val[0] == entry.word { // known words
					sign = true
					// fmt.Printf("len(e.stack) = %d, e.counter = %d\n", len(e.stack), e.counter)
					if e.counter >= entry.qtyArgsNeed && len(e.stack) >= entry.qtyStackNeed {
						e.err = entry.function()
						// if e.err != nil {
						// 	outpErr := e.err
						// 	e.err = nil
						// 	return e.stack, outpErr
						// }
						if e.err == nil {
							if entry.qtyArgsGave == 0 {
								e.counter = 0
							} else {
								e.counter += entry.qtyArgsGave
							}
						}
					} else {
						e.err = fmt.Errorf("parsing %v: invalid number of arguments or stack size", val[0])
					}
				}
			}
			if !sign {	// trying to add a data to the stack
				v, err := strconv.Atoi(val[0])
				// if err != nil {
				// 	e.err = fmt.Errorf("%v", err)
				// }
				if err == nil {
					e.stack = append(e.stack, v)
					e.counter++
				} else {
					e.err = err
				}
				
			}
			// if e.err != nil {
			// 	outpErr = e.err
			// 	e.err = nil
			// }
			// return e.stack, outpErr





			// switch counter {
			// case 0:
			// 	switch val[0] {
			// 		case "dup":
			// 			e.err = e.dup()
			// 			counter = 1
			// 		case "over":
			// 			e.err = e.over()
			// 			counter = 1
			// 		case "drop":
			// 			e.err = e.drop()
			// 			counter = 1
			// 		case "swap":
			// 			e.err = e.swap()
			// 			counter = 1
			// 		case "+", "-", "*", "/": 
			// 			if counter != 2 {
			// 				e.err = fmt.Errorf("parsing %v: invalid syntax", val[0])
			// 			} else {
			// 				e.err =
			// 			}

			// 		default:
			// 			v, err := strconv.Atoi(val[0])
			// 			if err != nil {
			// 				e.err = fmt.Errorf("%v", err)
			// 			}
			// 			fmt.Println(v)
			// 	}
			// 	if e.err != nil {
			// 		outpErr := e.err
			// 		e.err = nil
			// 		return e.stack, outpErr
			// 	}
			// }
		} else { //definition of new words
			fmt.Println(val)
		}
	}
	if e.err != nil {
		outpErr = e.err
		e.err = nil
	}
	return e.stack, outpErr

	// return e.stack, e.err
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := &Evaluator{}
	e.dictionary = []KnownWord{
		{word: "+", function: e.plus, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1},
		{word: "-", function: e.minus, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1},
		{word: "*", function: e.multiply, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1},
		{word: "/", function: e.divide, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1},
		{word: "dup", function: e.dup, qtyArgsNeed: 0, qtyStackNeed: 1, qtyArgsGave: 0},
		{word: "over", function: e.over, qtyArgsNeed: 0, qtyStackNeed: 2, qtyArgsGave: 0},
		{word: "drop", function: e.drop, qtyArgsNeed: 0, qtyStackNeed: 1, qtyArgsGave: 0},
		{word: "swap", function: e.swap, qtyArgsNeed: 0, qtyStackNeed: 2, qtyArgsGave: 0},
	}
	e.err = nil
	e.counter = 0
	// e.stack = e.stack[:0]
	// return &Evaluator{}
	return e
}

