package util

// Returns a new slice removing the given index
// Making a whole new slice is useful for just guessing :>
func Remove(slice []int, index int) []int {
	var res []int
	for i := range slice {
		if i != index {
			res = append(res, slice[i])
		}
	}
	return res
}
