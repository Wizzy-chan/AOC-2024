package solutions

import (
	"strings"
	"strconv"
	"github.com/Wizzy-chan/AOC-2024/util"
)

func concat(x, y int) int {
	// ~> min 10^n > y (i.e. 100 > 95)
	// x * 10^n + y (i.e. 8 * 100 + 95 = 895)
	n := 1
	for n <= y {
		n *= 10
	}
	return x * n + y
}

func getSumsRecursive(part2 bool, n int, total int, ns []int) bool {
	if len(ns) == 0 && n == total {
		return true
	}
	if len(ns) == 0 && n != total {
		return false
	}
	if part2 {
		return getSumsRecursive(part2, n, total+ns[0], ns[1:]) || getSumsRecursive(part2, n, total*ns[0], ns[1:]) || getSumsRecursive(part2, n, concat(total, ns[0]), ns[1:])
	}
	return getSumsRecursive(part2, n, total+ns[0], ns[1:]) || getSumsRecursive(part2, n, total*ns[0], ns[1:])
}


func getSums(part2 bool, n int, ns []int) bool {
	return getSumsRecursive(part2, n, ns[0], ns[1:])
}

func Day7(input string) (int, int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	total1 := 0
	total2 := 0
	for _, line := range lines {
		v, vs, _ := strings.Cut(line, ": ")
		sum, _ := strconv.Atoi(v)
		nums := util.ToIntSlice(vs, " ")
		if getSums(false, sum, nums) {
			total1 += sum
		}
		if getSums(true, sum, nums) {
			total2 += sum
		}
	}
	return total1, total2
}
