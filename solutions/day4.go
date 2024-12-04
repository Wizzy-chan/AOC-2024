package solutions

import (
	"strings"
)

func Search(grid []string, x, y int) int {
	total := 0
	height := len(grid)
	width := len(grid[0]) // TODO: empty input?
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if (dx == 0 && dy == 0) { continue }
			buf := make([]byte, 4)
			for i := 0; i < 4; i++ {
				if (y + i*dy < 0 || y + i*dy >= height || x + i*dx < 0 || x + i*dx >= width) { continue }
				buf[i] += grid[y + i*dy][x + i*dx]
			}
			if string(buf) == "XMAS" {
				total++
			}
		}
	}
	return total
}

func Search2(grid []string, x, y int) bool {
	total := 0
	height := len(grid)
	width := len(grid[0]) // TODO: empty input?
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if (dx == 0 || dy == 0) { continue }
			buf := make([]byte, 3)
			for i := -1; i < 2; i++ {
				if (y + i*dy < 0 || y + i*dy >= height || x + i*dx < 0 || x + i*dx >= width) { continue }
				buf[i+1] += grid[y + i*dy][x + i*dx]
			}
			if string(buf) == "MAS" {
				total++
			}
		}
	}
	return total == 2
}

func Day4(input string) (int, int) {
	grid := strings.Split(input, "\n")
	grid = grid[:len(grid)-1]
	height := len(grid)
	width := len(grid[0]) // TODO: empty input?
	total1 := 0
	total2 := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			total1 += Search(grid, x, y)
			if (Search2(grid, x, y)) {
				total2++
			}
		}
	}
	return total1, total2
}
