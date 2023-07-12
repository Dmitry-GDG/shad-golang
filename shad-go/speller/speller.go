//go:build !solution

package speller

func convert999(n int64) string {
	c100 := map[int64]string{
		2:   "twenty",
		3:   "thirty",
		4:   "forty",
		5:   "fifty",
		6:   "sixty",
		7:   "seventy",
		8:   "eighty",
		9:   "ninety",
		100: "hundred",
	}

	c19 := [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	var str string
	sep := " "
	nbr := n % 100
	n /= 100
	if n > 0 {
		str += c19[n] + sep + c100[100]
	}
	if nbr < 20 && nbr != 0 {
		if n > 0 {
			str += sep
		}
		str += c19[nbr]
	} else {
		if n > 0 && nbr > 0 {
			str += sep
		}
		str += c100[nbr/10]
		if nbr%10 > 0 {
			str += "-" + c19[nbr%10]
		}
	}
	return str
}

func capacity(i int) string {
	switch i {
	case 3:
		return "billion"
	case 2:
		return "million"
	case 1:
		return "thousand"
	default:
		return ""
	}
}

func Spell(n int64) string {
	// return ""
	outpStr := ""
	minus := false
	sep := " "
	outpSlice := []string{}
	if n < 0 {
		n *= -1
		minus = true
	}
	if n == 0 {
		return "zero"
	}
	for n > 0 {
		outpSlice = append(outpSlice, convert999(n%1000))
		n /= 1000
	}

	if minus {
		outpStr += "minus "
	}
	for i := len(outpSlice) - 1; i >= 0; i-- {
		if i != len(outpSlice)-1 && outpSlice[i] != "" {
			outpStr += sep
		}
		outpStr += outpSlice[i]
		if len(outpSlice) > 1 && outpSlice[i] != "" && capacity(i) != "" {
			outpStr += sep + capacity(i)
		}
	}
	return (outpStr)
}

// go fmt ./speller/speller.go
// golangci-lint run ./speller/...
// go test -v ./speller/...
