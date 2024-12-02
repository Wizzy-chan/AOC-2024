package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

// TODO: Make using each day easier than having to modify fp + func call
func main() {
	filepath := "day2.txt"
	contentBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ERROR: Could not read file %s: %s\n", filepath, err) 
	}
	content := string(contentBytes)
	fmt.Println(day2(content))
}

// There is no math.Abs for ints its kinda fucked up
func Abs(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}

// Day 1 solution
func day1(input string) (int, int) {
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
		difference += Abs(slice1[i] - slice2[i])
		n1 := slice1[i]
	  for _, n2 := range slice2 {
			if n1 == n2 {
				similarity += n1
			}
		}
	}
	return difference, similarity
}

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
				return Safe(Remove(report, 0), false) || Safe(Remove(report, i), false) || Safe(Remove(report, i+1), false)
			}
			return false
		}
	}
	return true
}

// Day 2 solution
func day2(input string) (int, int) {
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
