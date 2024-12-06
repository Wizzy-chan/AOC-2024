package solutions

import (
	"strings"
	"github.com/Wizzy-chan/AOC-2024/util"
)

type state struct { pos, dir util.Vec2 }

func run(grid [][]bool, pos, dir util.Vec2) (map[util.Vec2]bool, bool) {
	visited := make(map[util.Vec2]bool)
	states := make(map[state]bool)
	height, width := len(grid), len(grid[0])
	loop := false
	for {
		visited[pos] = true
		if states[state{pos, dir}] {
			// Loop!!
			loop = true
			break
		}
		states[state{pos, dir}] = true
		newPos := pos.Add(dir)
		if 0 > newPos.X || 0 > newPos.Y || newPos.X >= width || newPos.Y >= height {
			break
		} else if grid[newPos.Y][newPos.X] {
			switch dir {
			case util.VEC2_UP:
				dir = util.VEC2_RIGHT
			case util.VEC2_RIGHT:
				dir = util.VEC2_DOWN
			case util.VEC2_DOWN:
				dir = util.VEC2_LEFT
			case util.VEC2_LEFT:
				dir = util.VEC2_UP
			}
		} else {
			pos = newPos
		}
	}
	return visited, loop
}

func Day6(input string) (int, int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // trim final empty string
	width := len(lines[0])
	height := len(lines)
	grid := make([][]bool, height)
	var pos util.Vec2
	dir := util.VEC2_UP
	for y := 0; y < height; y++ {
		grid[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			if lines[y][x] == '#' {
				grid[y][x] = true
			}
			if lines[y][x] == '^' {
				pos = util.Vec2{x, y}
			}
		}
	}
	resA, _ := run(grid, pos, dir)
	resB := 0
	
	for key := range resA {
		grid[key.Y][key.X] = true
		_, loop := run(grid, pos, dir)
		if loop {
			resB++
		}
		grid[key.Y][key.X] = false
	}

	
	return len(resA), resB
}
