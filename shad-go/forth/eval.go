//go:build !solution

package main

import (
	// "bufio"
	"fmt"
	// "os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) push(a int) {
	*s = append(*s, a)
}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	last := (*s)[len(*s)-1]
	*s = (*s)[:(len(*s) - 1)]
	return last, nil
}

type FF struct {
	name         func() error
	qtyArgsNeed  int
	qtyStackNeed int
	qtyArgsGave  int
}

type KnownWord struct {
	word     string
	function []FF
	numbers  []int
}

type Evaluator struct {
	stack      Stack
	dictionary []KnownWord
	err        error
	counter    int
	funcNow    string
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
func (e *Evaluator) mathOperation1(mathOp func(int, int) int) func() error {
	return func() error {
		a, err := e.stack.pop()
		if err != nil {
			return err
		}
		if a == 0 {
			return fmt.Errorf("error: division by zero")
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
	return e.mathOperation(func(a int, b int) int { return a + b })()
}
func (e *Evaluator) minus() error {
	return e.mathOperation(func(a int, b int) int { return b - a })()
}
func (e *Evaluator) multiply() error {
	return e.mathOperation(func(a int, b int) int { return a * b })()
}
func (e *Evaluator) divide() error {
	return e.mathOperation1(func(a int, b int) int { return b / a })()
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
func (e *Evaluator) push(a int) {
	e.stack.push(a)
}
func (e *Evaluator) pushNumberFromFunc() error { // qtyArgsNeed = 0, qtyStackNeed = 0, qtyArgsGave 1
	for i := 0; i < len(e.dictionary); i++ {
		if e.funcNow == (e.dictionary)[i].word {
			first := (e.dictionary)[i].numbers[0]
			e.dictionary[i].numbers = e.dictionary[i].numbers[1:]
			e.stack.push(first)
			break
		}
	}
	return nil
}
func (e *Evaluator) splitInputData(s string) [][]string {
	sSlice := [][]string{}
	begin := 0
	end := 32
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			i++
		}
		if s[i] == ':' {
			end = 59
			i++
			if s[i] == 32 {
				i++
			}
		}
		begin = i
		for ; (i < len(s)) && (s[i] != byte(end)); i++ {
		}
		tokenSlice := strings.Split(strings.TrimSpace(s[begin:i]), " ")
		sSlice = append(sSlice, tokenSlice)
		end = 32
	}
	return sSlice
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(s string) ([]int, error) {
	var outpErr error
	outpErr = nil
	s = strings.TrimSpace(s)
	sSlice := e.splitInputData(s)

	for _, val := range sSlice {
		if len(val) == 1 { // tryinf to find known words or data
			sign := false
			for _, entry := range e.dictionary {
				if strings.ToLower(val[0]) == entry.word { // known words
					e.funcNow = strings.ToLower(val[0])
					sign = true
					for i := 0; i < len(entry.function); i++ {
						if e.counter >= (entry.function)[i].qtyArgsNeed && len(e.stack) >= (entry.function)[i].qtyStackNeed {
							e.err = entry.function[i].name()
							if e.err != nil {
								break
							}
							if e.err == nil {
								e.counter += entry.function[i].qtyArgsGave
							}
						} else {
							e.err = fmt.Errorf("parsing %v: invalid number of arguments or stack size", val[0])
						}
					}
					e.funcNow = ""
				}
			}
			if !sign { // trying to add a data to the stack
				v, err := strconv.Atoi(val[0])
				if err == nil {
					e.push(v)
					e.counter++
				} else {
					e.err = err
				}
			}
		} else { //definition of new words
			if len(val) < 2 {
				e.err = fmt.Errorf("parsing %v: invalid syntax", val)
			} else {
				_, err := strconv.Atoi(val[0])
				if err == nil { // can not override number
					return e.stack, fmt.Errorf("parsing %v: can not override number", val[0])
				}
				newWord := new(KnownWord)
				newWord.word = strings.ToLower(val[0])
				sign := false
				for i := 1; i < len(val); i++ {
					sign = false
					for _, entry := range e.dictionary {
						if strings.ToLower(val[i]) == entry.word { // known words
							sign = true
							newWord.function = append(newWord.function, entry.function...)
							newWord.numbers = append(newWord.numbers, entry.numbers...)
							break
						}
					}
					if !sign { // there is no such func in db, add it
						v, err := strconv.Atoi(val[i])
						if err == nil {
							newFunction := FF{name: e.pushNumberFromFunc, qtyArgsNeed: 0, qtyStackNeed: 0, qtyArgsGave: 1}
							newWord.function = append(newWord.function, newFunction)
							newWord.numbers = append(newWord.numbers, v)
							e.counter = 0
						} else {
							e.err = fmt.Errorf("parsing %v: %v", val[i], err)
						}
					}
				}
				ss := new([]KnownWord)
				for j := 0; j < len(e.dictionary); j++ {
					if newWord.word != e.dictionary[j].word { // known word // !!!!! redefining an existing word !!!!!!
						*ss = append(*ss, e.dictionary[j])
					}
				}
				*ss = append(*ss, *newWord)
				e.dictionary = *ss
			}
		}
	}
	if e.err != nil {
		outpErr = e.err
		e.err = nil
	}
	return e.stack, outpErr
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := &Evaluator{}
	e.dictionary = []KnownWord{
		{word: "+", function: []FF{{name: e.plus, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1}}},
		{word: "-", function: []FF{{name: e.minus, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1}}},
		{word: "*", function: []FF{{name: e.multiply, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1}}},
		{word: "/", function: []FF{{name: e.divide, qtyArgsNeed: 2, qtyStackNeed: 0, qtyArgsGave: 1}}},
		{word: "dup", function: []FF{{name: e.dup, qtyArgsNeed: 0, qtyStackNeed: 1, qtyArgsGave: 1}}},
		{word: "over", function: []FF{{name: e.over, qtyArgsNeed: 0, qtyStackNeed: 2, qtyArgsGave: 0}}},
		{word: "drop", function: []FF{{name: e.drop, qtyArgsNeed: 0, qtyStackNeed: 1, qtyArgsGave: -1}}},
		{word: "swap", function: []FF{{name: e.swap, qtyArgsNeed: 0, qtyStackNeed: 2, qtyArgsGave: 0}}},
	}
	e.err = nil
	e.counter = 0
	e.funcNow = ""
	return e
}

// clear
// go fmt ./forth/eval.go
// golangci-lint run ./forth/...
// go test -v ./forth/...
