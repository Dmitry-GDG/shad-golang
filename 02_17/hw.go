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

// func BubbleSort(outpSlice []string, prereqs map[string][]string) {
// 	Swapped := true
// 	for Swapped {
// 		Swapped = false
// 		for i := 0; i < len(outpSlice)-1; i++ {
// 			// if numbers[i] > outpSlice[i+1] {
// 			for prereq, slicePrereqs := range(prereqs) {
// 				if prereq == outpSlice[i] {
// 					// for k := j + 1; k < len()
// 					if isPrereqInSlice(slicePrereqs, outpSlice[i + 1]) {
// 						outpSlice[i], outpSlice[i + 1] = outpSlice[i + 1], outpSlice[i]
// 						Swapped = true
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func GetCourseList(prereqs map[string][]string) []string {
	// return []string{}
	var subjSlice, outpSlice []string
	subjMap := make(map[int]string)
	// making a slice and a map of subjects
	for prereq, slicePrereqs := range(prereqs) {
		if !isPrereqInSlice(subjSlice, prereq) {
			subjSlice = append(subjSlice, prereq)
		}
		for i := 0; i < len(slicePrereqs); i++ {
			// fmt.Printf("i: %d,\tprereq: %s\n", i, slicePrereqs[i])
			if !isPrereqInSlice(subjSlice, slicePrereqs[i]) {
				subjSlice = append(subjSlice, slicePrereqs[i])
			}
		}
	}
	for i := 0; i < len(subjSlice); i++ {
		subjMap[i] = subjSlice[i]
	}
	// fmt.Println(subjMap)
	// // BubbleSorting
	// BubbleSort(outpSlice, prereqs)
	// // for i := 0; i < len(outpSlice) - 1; i++ {
	// // 	for j := 0; j < len(outpSlice) - i - 1; j++ {
	// // 		for prereq, slicePrereqs := range(prereqs) {
	// // 			if prereq == outpSlice[j] {
	// // 				// for k := j + 1; k < len()
	// // 				if isPrereqInSlice(slicePrereqs, outpSlice[j + 1]) {
	// // 					outpSlice[j], outpSlice[j + 1] = outpSlice[j + 1], outpSlice[j]
	// // 					// break
	// // 				}
	// // 			}
	// // 		}
	// // 	}
	// // }

	// Sorting
	for i := 0; i < len(subjMap); {
		found := false
		// check := 0
		for prereq, slicePrereqs := range(prereqs) {
			if prereq == subjMap[i] {
				for k := 0; k < len(slicePrereqs); k++ {
					fmt.Println("Here")
					for j := i + 1; j < len(subjMap); j++ {
						fmt.Printf("j = %d, i = %d\n", j, i)
						fmt.Printf("before: %v\n", subjMap)
						if subjMap[j] == slicePrereqs[k] {
							tmp := subjMap[j]
							for l := j; l > i; l-- {
								subjMap[l] = subjMap[l - 1]
							}
							subjMap[i] = tmp
							fmt.Printf("after: %v\n", subjMap)
							found = true
							j = i + 1
							// check++
							// fmt.Println(check)
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
	// for i := 0; i < len(outpSlice) - 1; i++ {
	// 	for j := i + 1; j < len(outpSlice);  {
	// 		for prereq, slicePrereqs := range(prereqs) {
	// 			if prereq == outpSlice[i] {
	// 				for k := 0, l := j; k < len(slicePrereqs), l < len(outpSlice); {
	// 					// for k := j + 1; k < len()
	// 					if isPrereqInSlice(slicePrereqs, outpSlice[k]) {

	// 						outpSlice[j], outpSlice[j + 1] = outpSlice[j + 1], outpSlice[j]
	// 						// break
	// 						// i = 0
	// 					} else {
	// 						k++
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// Checking
	// for i := 0; i < len(outpSlice) - 1; i++ {
	// 	for j := i + 1; j < len(outpSlice); j++ {
	// 		for prereq, slicePrereqs := range(prereqs) {
	// 			if prereq == outpSlice[i] {
	// 				if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
	// 					// panic("cyclic dependency")
	// 					fmt.Println("here")
	// 				}
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	// filling outpSlice
	for i := 0; i < len(subjMap); i++ {
		outpSlice = append(outpSlice, subjMap[i])
	}

	// // Checking
	// for i := 0; i < len(outpSlice) - 1; i++ {
	// 	for j := i + 1; j < len(outpSlice); j++ {
	// 		for prereq, slicePrereqs := range(prereqs) {
	// 			if prereq == outpSlice[i] {
	// 				if isPrereqInSlice(slicePrereqs, outpSlice[j]) {
	// 					panic("cyclic dependency")
	// 				}
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	return outpSlice
}

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
	var computerScience = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	// for prereq, slicePrereqs := range(computerScience) {
	// 	fmt.Printf("prereq:\t%v\nslicePrereq:\t%v\n", prereq, slicePrereqs)
	// }

	fmt.Println(GetCourseList(computerScience))

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

	var strangeScience = map[string][]string{
		"1": {"0"},
		"2": {"1", "3"},
		"3": {"2"},
	}
	// require.Panics(t, func() {
	// 	impl(t, strangeScience, GetCourseList(strangeScience))
	// }
	fmt.Println(GetCourseList(strangeScience))
}