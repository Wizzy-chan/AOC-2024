package solutions

import (
	"strings"
	"strconv"
	"github.com/Wizzy-chan/AOC-2024/util"
)

// func for day 2
// this func is really bad and very botched
// TODO: Maybe come up with a better fix for part 2
func Safe(report []int, part2 bool) bool {
	sign := 1 // assume increasing
	// Check first two readings to work out which way
	// Maybe part 2 should check more than one pair
	if report[0] > report[1] {
		sign = -1 // actually decreasing
	}
	for i := 0; i < len(report) - 1; i++ {
		delta := report[i+1] - report[i]
		// 1 <= delta * sign <= 3 is valid
		if delta * sign < 1 || delta * sign > 3 {
			if part2 {
				// Adding in the 0 here is a botched solution to a specific edge case
				// If the first transition is the "bad" transition it can slip through
				// and 1-2 will be read as the bad transition
				// so we check both readings in the bad transition and the very first reading just in case
				return Safe(util.Remove(report, 0), false) || Safe(util.Remove(report, i), false) || Safe(util.Remove(report, i+1), false)
			}
			return false
		}
	}
	return true
}

// Day 2 solution
func Day2(input string) (int, int) {
	var reports [][]int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		var report []int
		ns := strings.Split(line, " ")
		for _, n := range ns {
			v, _ := strconv.Atoi(n)
			report = append(report, v)
		}
		reports = append(reports, report)
	}
	total1, total2 := 0, 0
	for _, report := range reports {
		if Safe(report, false) {
			total1++
		}
		if Safe(report, true) {
			total2++
		}
	}
	return total1, total2
}
