package util

import (
	"strings"
	"strconv"
)

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

// Convert a string n1,n2,n3 into a int slice
func ToIntSlice(str, delim string) []int {
	strs := strings.Split(str, delim)
	res := make([]int, 0)
	for _, s := range strs {
		n, _:= strconv.Atoi(s)
		res = append(res, n)
	}
	return res
}
