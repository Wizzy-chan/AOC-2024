package solutions

import (
	"regexp"
	"strconv"
)

// Day 3 solution
func Day3(input string) (int, int) {
	total1 := 0
	total2 := 0
	enabled := true
	exp := regexp.MustCompile(`(mul|do|don't)\((?:([0-9]{1,3}),([0-9]{1,3}))?\)`)
	matches := exp.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		switch match[1] {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		case "mul":
			n1, _ := strconv.Atoi(match[2])
			n2, _ := strconv.Atoi(match[3])
			total1 += n1 * n2
			if enabled {
				total2 += n1 * n2
			}
		}
	}
	return total1, total2
}
