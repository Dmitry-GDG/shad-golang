//go:build !solution

package hogwarts

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

// go fmt ./hogwarts/courselist.go
// golangci-lint run ./hogwarts/...
// go test -v ./hogwarts/...
