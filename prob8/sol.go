package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() (ss [][]rune) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		ss = append(ss, []rune(val))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func nc(x, y int) coordinate {
	return coordinate{x, y}
}

type coordinate struct {
	x int
	y int
}

func (c coordinate) antinodes(other coordinate, repeat bool, x, y int) []coordinate {
	xDiff := c.x - other.x
	yDiff := c.y - other.y
	var ncs []coordinate
	for curX, curY := c.x, c.y; curX >= 0 && curY >= 0 && curX <= x && curY <= y; curX, curY = curX+xDiff, curY+yDiff {
		if !repeat && curX == c.x && curY == c.y {
			continue
		}
		ncs = append(ncs, nc(curX, curY))
		if !repeat {
			break
		}
	}
	return ncs
}

func (c coordinate) isDupe(other coordinate) bool {
	return c.x == other.x && c.y == other.y
}

func main() {
	sol1()
	sol2()
}

func sol1() {
	grid := readInput()

	ants := make(map[rune][]coordinate)
	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				ants[cell] = append(ants[cell], nc(x, y))
			}
		}
	}

	var antinodes []coordinate
	for _, coords := range ants {
		for i, coord1 := range coords {
			for j, coord2 := range coords {
				if i == j {
					continue
				}
				antinodes = append(antinodes, coord1.antinodes(coord2, false, len(grid[0])-1, len(grid)-1)...)
			}
		}
	}

	var valid []coordinate

	for i, anti1 := range antinodes {
		isDupe := false
		for j, anti2 := range antinodes {
			if j <= i {
				continue
			}
			if anti1.isDupe(anti2) {
				isDupe = true
				break
			}
		}
		if !isDupe {
			valid = append(valid, anti1)
		}
	}

	for _, anti := range valid {
		if grid[anti.y][anti.x] == '.' {
			grid[anti.y][anti.x] = '#'
		}
	}

	log.Println(len(valid))
}

func sol2() {
	grid := readInput()

	ants := make(map[rune][]coordinate)
	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				ants[cell] = append(ants[cell], nc(x, y))
			}
		}
	}

	var antinodes []coordinate
	for _, coords := range ants {
		for i, coord1 := range coords {
			for j, coord2 := range coords {
				if i == j {
					continue
				}
				antinodes = append(antinodes, coord1.antinodes(coord2, true, len(grid[0])-1, len(grid)-1)...)
			}
		}
	}

	var valid []coordinate

	for i, anti1 := range antinodes {
		isDupe := false
		for j, anti2 := range antinodes {
			if j <= i {
				continue
			}
			if anti1.isDupe(anti2) {
				isDupe = true
				break
			}
		}
		if !isDupe {
			valid = append(valid, anti1)
		}
	}

	for _, anti := range valid {
		if grid[anti.y][anti.x] == '.' {
			grid[anti.y][anti.x] = '#'
		}
	}

	log.Println(len(valid))
}
