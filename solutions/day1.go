package solutions

import (
	"strings"
	"strconv"
	"sort"
	"github.com/Wizzy-chan/AOC-2024/util"
)

// Day 1 solution
func Day1(input string) (int, int) {
	var slice1, slice2 []int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		ns := strings.Split(line, "   ")
		n1, _ := strconv.Atoi(ns[0])
		n2, _ := strconv.Atoi(ns[1])
		slice1 = append(slice1, n1)
		slice2 = append(slice2, n2)
	}
	sort.Ints(slice1)
	sort.Ints(slice2)
	difference := 0
	similarity := 0
	for i := range slice1 {
		difference += util.Abs(slice1[i] - slice2[i])
		n1 := slice1[i]
	  for _, n2 := range slice2 {
			if n1 == n2 {
				similarity += n1
			}
		}
	}
	return difference, similarity
}
