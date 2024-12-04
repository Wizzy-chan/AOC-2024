package main

import (
	"fmt"
	"io/ioutil"
	"github.com/Wizzy-chan/AOC-2024/solutions"
)

// TODO: Make using each day easier than having to modify fp + func call
func main() {
	filepath := "inputs/day4.txt"
	contentBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ERROR: Could not read file %s: %s\n", filepath, err) 
	}
	content := string(contentBytes)
	fmt.Println(solutions.Day4(content))
}
