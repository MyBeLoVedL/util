package slice

// return true if a slice contain an element
func ContainInt(source []int, element int) bool {
	for _, val := range source {
		if val == element {
			return true
		}
	}
	return false
}
