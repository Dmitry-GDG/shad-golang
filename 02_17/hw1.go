//courselist.go
package main
import "fmt"

func isPrereqInSlice(outpSlice []string, prereq string) bool {
	for i := 0; i < len(outpSlice); i++ {
		if prereq == outpSlice[i] {
			return true
		}
	}
	return false
}

func GetCourseList(prereqs map[string][]string) []string {
	// return []string{}
	var subjSlice, outpSlice []string
	subjMap := make(map[int]string)

	// making a slice and a map of subjects
	for prereq, slicePrereqs := range prereqs {
		if !isPrereqInSlice(subjSlice, prereq) {
			subjSlice = append(subjSlice, prereq)
		}
		for i := 0; i < len(slicePrereqs); i++ {
			if !isPrereqInSlice(subjSlice, slicePrereqs[i]) {
				subjSlice = append(subjSlice, slicePrereqs[i])
			}
		}
	}
	for i := 0; i < len(subjSlice); i++ {
		subjMap[i] = subjSlice[i]
	}

	// Sorting
	check := 0
	for i := 0; i < len(subjMap); {
		found := false
		for prereq, slicePrereqs := range prereqs {
			if prereq == subjMap[i] {
				for k := 0; k < len(slicePrereqs); k++ {
					for j := i + 1; j < len(subjMap); j++ {
						if subjMap[j] == slicePrereqs[k] {
							tmp := subjMap[j]
							for l := j; l > i; l-- {
								subjMap[l] = subjMap[l-1]
							}
							subjMap[i] = tmp
							found = true
							j = i + 1
							check++
							if check >= len(subjMap) {
								panic("cyclic dependency")
							}
						}
					}
				}
			}
		}
		if !found {
			i++
		}
	}

	// filling outpSlice
	for i := 0; i < len(subjMap); i++ {
		outpSlice = append(outpSlice, subjMap[i])
	}

	return outpSlice
}

// func isPrereqInSlice(outpSlice []string, prereq string) bool {
// 	for i := 0; i < len(outpSlice); i++ {
// 		if prereq == outpSlice[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func compare(x, pivot string, prereqs map[string][]string) int {
// 	for prereq, slicePrereqs := range(prereqs) {
// 		if x == prereq {
// 			for i := 0; i < len(slicePrereqs); i++ {
// 				if slicePrereqs[i] == pivot {
// 					return -1
// 				}
// 			}
// 		}
// 	}
// 	return 1
// }

// func qSort(outpSlice []string, prereqs map[string][]string) []string {
// 	if len(outpSlice) <= 1 {
// 		return outpSlice
// 	}

// 	left := make([]string, 0)
// 	right := make([]string, 0)


// 	// // for len(outpSlice) > 0 {
// 	// 	pivot := outpSlice[0]
// 	// 	outpSlice = outpSlice[1:]

// 	// 	// for _, x := range outpSlice {

			
// 	// 	// 	if x.Compare(pivot) < 0 {
// 	// 	// 		left = append(left, x)
// 	// 	// 	} else {
// 	// 	// 		right = append(right, x)
// 	// 	// 	}
// 	// 	// }

// 	// 	for _, x := range outpSlice {
// 	// 		// fmt.Printf("x: %s, pivot: %s, compare(x, pivot, prereqs): %d\n", x, pivot, compare(x, pivot, prereqs))
// 	// 		if compare(x, pivot, prereqs) < 0 {
// 	// 			left = append(left, x)
// 	// 		} else {
// 	// 			right = append(right, x)
// 	// 		}
// 	// 	}
// 	// 	// fmt.Printf("for left: %s\n", left)
// 	// 	// fmt.Printf("for right: %s\n", right)
// 	// 	// fmt.Printf("for: %s\n", outpSlice)

// 	// 	result := append(qSort(left, prereqs), pivot)
// 	// 	result = append(result, qSort(right, prereqs)...)
// 	// 	return(result)
// 	// // }


// 	// for len(outpSlice) > 0 {
// 		pivot := outpSlice[0]
// 		outpSlice = outpSlice[1:]

// 		// for _, x := range outpSlice {

			
// 		// 	if x.Compare(pivot) < 0 {
// 		// 		left = append(left, x)
// 		// 	} else {
// 		// 		right = append(right, x)
// 		// 	}
// 		// }

// 		for _, x := range outpSlice {
// 			// fmt.Printf("x: %s, pivot: %s, compare(x, pivot, prereqs): %d\n", x, pivot, compare(x, pivot, prereqs))
// 			if compare(x, pivot, prereqs) < 0 {
// 				left = append(left, x)
// 			} else {
// 				right = append(right, x)
// 			}
// 		}
// 		fmt.Printf("for left: %s\n", left)
// 		fmt.Printf("for right: %s\n", right)
// 		fmt.Printf("for: %s\n", outpSlice)

// 		if len(left) == 0 {
// 			outpSlice = right
// 			right = nil
// 		} else if len(right) == 0 {
// 			outpSlice = left
// 			left = nil
// 		} else {
// 			outpSlice = append(qSort(left, prereqs), pivot)
// 			outpSlice = append(outpSlice, qSort(right, prereqs)...)
// 			left = nil
// 			right = nil
// 			// outpSlice = result
// 		}
// 	// }
// 	// return result
// 	return outpSlice
// }

// func GetCourseList(prereqs map[string][]string) []string {
// 	// return []string{}

// 	var outpSlice []string
// 	// making a slice of subjects
// 	for prereq, slicePrereqs := range(prereqs) {
// 		if !isPrereqInSlice(outpSlice, prereq) {
// 			outpSlice = append(outpSlice, prereq)
// 		}
// 		for i := 0; i < len(slicePrereqs); i++ {
// 			if !isPrereqInSlice(outpSlice, slicePrereqs[i]) {
// 				outpSlice = append(outpSlice, slicePrereqs[i])
// 			}
// 		}
// 	}
// 	// fmt.Println(outpSlice)

// 	// Sorting
// 	outpSlice = qSort(outpSlice, prereqs)
// 	// fmt.Println(outpSlice)

// 	// Checking
// 	// for i := 0; i < len(outpSlice) - 1; i++ {
// 	// 	for j := i + 1; j < len(outpSlice); j++ {
// 	// 		for prereq, slicePrereqs := range(prereqs) {
// 	// 			if prereq == outpSlice[i] {
// 	// 				if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 	// 					// panic("cyclic dependency")
// 	// 					fmt.Println("here")
// 	// 				}
// 	// 				break
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// 	// filling outpSlice
// 	// for i := 0; i < len(subjMap); i++ {
// 	// 	outpSlice = append(outpSlice, subjMap[i])
// 	// }

// 	// // Checking
// 	// for i := 0; i < len(outpSlice) - 1; i++ {
// 	// 	for j := i + 1; j < len(outpSlice); j++ {
// 	// 		for prereq, slicePrereqs := range(prereqs) {
// 	// 			if prereq == outpSlice[i] {
// 	// 				if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 	// 					panic("cyclic dependency")
// 	// 				}
// 	// 				break
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	return outpSlice
// }

// func GetCourseList(prereqs map[string][]string) []string {
// 	// return []string{}
// 	var outpSlice []string
// 	// making a slice of subjects
// 	for prereq, slicePrereqs := range(prereqs) {
// 		if !isPrereqInSlice(outpSlice, prereq) {
// 			outpSlice = append(outpSlice, prereq)
// 		}
// 		for i := 0; i < len(slicePrereqs); i++ {
// 			if !isPrereqInSlice(outpSlice, slicePrereqs[i]) {
// 				outpSlice = append(outpSlice, slicePrereqs[i])
// 			}
// 		}
// 	}
// 	// BubbleSorting
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := 0; j < len(outpSlice) - i - 1; j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[j] {
// 					// for k := j + 1; k < len()
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j + 1]) {
// 						outpSlice[j], outpSlice[j + 1] = outpSlice[j + 1], outpSlice[j]
// 						// break
// 					}
// 				}
// 			}
// 		}
// 	}
// 	// Checking
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := i + 1; j < len(outpSlice); j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[i] {
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 						// panic("cyclic dependency")
// 						fmt.Println("here")
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return outpSlice
// }

// func isPrereqInSlice(outpSlice []string, prereq string) bool {
// 	for i := 0; i < len(outpSlice); i++ {
// 		if prereq == outpSlice[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetCourseList(prereqs map[string][]string) []string {
// 	// return []string{}
// 	var outpSlice []string
// 	// making a slice of subjects
// 	for prereq, slicePrereqs := range(prereqs) {
// 		if !isPrereqInSlice(outpSlice, prereq) {
// 			outpSlice = append(outpSlice, prereq)
// 		}
// 		for i := 0; i < len(slicePrereqs); i++ {
// 			if !isPrereqInSlice(outpSlice, slicePrereqs[i]) {
// 				outpSlice = append(outpSlice, slicePrereqs[i])
// 			}
// 		}
// 	}
// 	// BubbleSorting
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := 0; j < len(outpSlice) - i - 1; j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[j] {
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j + 1]) {
// 						outpSlice[j], outpSlice[j + 1] = outpSlice[j + 1], outpSlice[j]
// 					}
// 					// break
// 				}
// 			}
// 		}
// 	}
// 	// Checking
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := i + 1; j < len(outpSlice); j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[i] {
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 						panic("cyclic dependency")
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return outpSlice
// }

// func GetCourseList(prereqs map[string][]string) []string {
// 	// return []string{}
// 	var outpSlice []string
// 	// making a slice of subjects
// 	for prereq, slicePrereqs := range(prereqs) {
// 		if !isPrereqInSlice(outpSlice, prereq) {
// 			outpSlice = append(outpSlice, prereq)
// 		}
// 		for i := 0; i < len(slicePrereqs); i++ {
// 			if !isPrereqInSlice(outpSlice, slicePrereqs[i]) {
// 				outpSlice = append(outpSlice, slicePrereqs[i])
// 			}
// 		}
// 	}
// 	// BubbleSorting
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := i + 1; j < len(outpSlice); j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[i] {
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 						outpSlice[i], outpSlice[j] = outpSlice[j], outpSlice[i]
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	// Checking
// 	for i := 0; i < len(outpSlice) - 1; i++ {
// 		for j := i + 1; j < len(outpSlice); j++ {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[i] {
// 					if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
// 						panic("cyclic dependency")
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return outpSlice
// }

func main() {
	// var computerScience = map[string][]string{
	// 	"algorithms": {"data structures"},
	// 	"calculus":   {"linear algebra"},
	// 	"compilers": {
	// 		"data structures",
	// 		"formal languages",
	// 		"computer organization",
	// 	},
	// 	"data structures":       {"discrete math"},
	// 	"databases":             {"data structures"},
	// 	"discrete math":         {"intro to programming"},
	// 	"formal languages":      {"discrete math"},
	// 	"networks":              {"operating systems"},
	// 	"operating systems":     {"data structures", "computer organization"},
	// 	"programming languages": {"data structures", "computer organization"},
	// }
	// // for prereq, slicePrereqs := range(computerScience) {
	// // 	fmt.Printf("prereq:\t%v\nslicePrereq:\t%v\n", prereq, slicePrereqs)
	// // }

	// fmt.Println(GetCourseList(computerScience))

	// var linearScience = map[string][]string{
	// 	"1": {"0"},
	// 	"2": {"1"},
	// 	"3": {"2"},
	// 	"4": {"3"},
	// 	"5": {"4"},
	// 	"6": {"5"},
	// 	"7": {"6"},
	// 	"8": {"7"},
	// 	"9": {"8"},
	// }
	// fmt.Println(GetCourseList(linearScience))

	// var weirdScience = map[string][]string{
	// 	"купи":   {"продай"},
	// 	"продай": {"купи"},
	// }
	// // require.Panics(t, func() {
	// // 	impl(t, weirdScience, GetCourseList(weirdScience))
	// // })
	// fmt.Println(GetCourseList(weirdScience))

	// var strangeScience = map[string][]string{
	// 	"1": {"0"},
	// 	"2": {"1", "3"},
	// 	"3": {"2"},
	// }
	// // require.Panics(t, func() {
	// // 	impl(t, strangeScience, GetCourseList(strangeScience))
	// // }
	// fmt.Println(GetCourseList(strangeScience))

	var acristin = map[string][]string{
		"1": {"0"},
		"4": {"3"},
		"3": {"2"},
		"2": {"1"},
	}
	fmt.Printf("before: %v\n", acristin)
	fmt.Printf("after: %v\n", GetCourseList(acristin))

	var acristin1 = map[string][]string{
		"1": {"3"},
		"4": {"2"},
		"3": {"0"},
		"2": {"1"},
		"0": {},
	}
	fmt.Printf("before: %v\n", acristin1)
	fmt.Printf("after: %v\n", GetCourseList(acristin1))
}