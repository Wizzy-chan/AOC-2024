package solutions

import (
	"strings"
	"strconv"
	"github.com/Wizzy-chan/AOC-2024/util"
)

func CorrectOrder(update []int, rules [100][100]bool) bool {
	for i := range update {
		for j := 0; j < i; j++ {
			if rules[update[i]][update[j]] {
				return false
			}
		}
	}
	return true
}

func Sort(update []int, rules [100][100]bool) []int {
	// While we aren't sorted
	for !CorrectOrder(update, rules) {
		// If we find two incorrectly order pages, swap them
		for i := range update {
			for j := 0; j < i; j++ {
				if rules[update[i]][update[j]] {
					update[i], update[j] = update[j], update[i]
				}
			}
		}
	}
	return update
}

func Day5(input string) (int, int) {
	var rules [100][100]bool
	rulesString, updates, _ := strings.Cut(input, "\n\n")
	for _, line := range strings.Split(rulesString, "\n") {
		if line == "" { continue }
		before, after, _ := strings.Cut(line, "|")
		n, _ := strconv.Atoi(before)
		m, _ := strconv.Atoi(after)
		rules[n][m] = true
	}
	total1 := 0
	total2 := 0
	for _, line := range strings.Split(updates, "\n") {
		if line == "" { continue }
		update := util.ToIntSlice(line, ",")
		if CorrectOrder(update, rules) {
			total1 += update[(len(update) - 1) / 2]
		} else {
			update = Sort(update, rules)
			total2 += update[(len(update) - 1) / 2]
		}
	}
	return total1, total2
}
