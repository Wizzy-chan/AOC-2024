package solutions

import (
	"strings"
	"github.com/Wizzy-chan/AOC-2024/util"
)

type Beacon struct {
	pos util.Vec2;
	freq byte;
}

func countAntinodes(part2 bool, beacons []Beacon, tl, br util.Vec2) int {
	nodes := make(map[util.Vec2]bool, 0)
	// Find all pairs of the same beacon
	for i := range beacons {
		for j := 0; j < i; j++ {
			if beacons[i].freq != beacons[j].freq {
				continue
			}
			// Find out where their antinodes are (d = b1 - b2, a1 = b1 + d, a2 = b2 - d)
			d := beacons[i].pos.Sub(beacons[j].pos) // d = ->b2b
			if !part2 {
				n1 := beacons[i].pos.Add(d)
				n2 := beacons[j].pos.Sub(d)
				// Check if they are in bounds
				// Add them to a set
				if n1.InBounds(tl, br) {
					nodes[n1] = true
				}
				if n2.InBounds(tl, br) {
					nodes[n2] = true
				}
			} else {
				n1 := beacons[i].pos
				n2 := beacons[j].pos
				for n1.InBounds(tl, br) {
					nodes[n1] = true
					n1 = n1.Add(d)
				}
				for n2.InBounds(tl, br) {
					nodes[n2] = true
					n2 = n2.Sub(d)
				}
			}
		}
	}
	return len(nodes)
}

func Day8(input string) (int, int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	beacons := make([]Beacon, 0)
	// Find all beacons
	for y, line := range lines {
		for x := range line {
			if line[x] != '.' {
				beacons = append(beacons, Beacon{util.Vec2{x, y}, line[x]})
			}
		}
	}
	tl := util.Vec2{0, 0}
	br := util.Vec2{len(lines[0]), len(lines)}
	return countAntinodes(false, beacons, tl, br), countAntinodes(true, beacons, tl, br)
}
