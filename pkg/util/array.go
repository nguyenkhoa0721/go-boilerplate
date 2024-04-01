package util

func RemoveDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueArr []string

	for _, element := range arr {
		if !uniqueMap[element] {
			uniqueMap[element] = true
			uniqueArr = append(uniqueArr, element)
		}
	}

	return uniqueArr
}

func Contains(elems []string, elem string) bool {
	for _, e := range elems {
		if elem == e {
			return true
		}
	}
	return false
}
