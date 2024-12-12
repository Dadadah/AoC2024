package main

import (
	"bufio"
	"log"
	"os"
)

type region struct {
	grid       [][]rune
	identifier rune
	perim      int
}

func readInput() (grid [][]rune) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		grid = append(grid, []rune(val))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	sol1()
	sol2()
}

func (r *region) populate(grid [][]rune, x, y int) bool {
	if r.grid[y][x] == r.identifier {
		return true
	}
	if grid[y][x] == r.identifier {
		grid[y][x] = '.'
		r.grid[y][x] = r.identifier
	} else {
		if r.grid[y][x] == '.' {
			r.grid[y][x] = '1'
		} else {
			r.grid[y][x]++
		}
		return false
	}
	sides := 4
	if x-1 >= 0 {
		if r.populate(grid, x-1, y) {
			sides--
		}
	}
	if x+1 < len(grid[0]) {
		if r.populate(grid, x+1, y) {
			sides--
		}
	}
	if y-1 >= 0 {
		if r.populate(grid, x, y-1) {
			sides--
		}
	}
	if y+1 < len(grid) {
		if r.populate(grid, x, y+1) {
			sides--
		}
	}
	r.perim += sides
	return true
}

var runes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^'}

var indexFromRune = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '!': 10, '@': 11, '#': 12, '$': 13, '%': 14, '^': 15}

func (r *region) populate2(grid [][]rune, x, y int, dir int) bool {
	baseGridX := x - 1
	baseGridY := y - 1
	if r.grid[y][x] == r.identifier {
		return true
	}
	if baseGridX < 0 || baseGridX >= len(grid[0]) || baseGridY < 0 || baseGridY >= len(grid) {
		if r.grid[y][x] == '.' {
			r.grid[y][x] = runes[1<<dir]
		}
		index := indexFromRune[r.grid[y][x]]
		index |= 1 << dir
		r.grid[y][x] = runes[index]
		return false
	}
	if grid[baseGridY][baseGridX] == r.identifier {
		grid[baseGridY][baseGridX] = '.'
		r.grid[y][x] = r.identifier
	} else {
		if r.grid[y][x] == '.' {
			r.grid[y][x] = runes[1<<dir]
		}
		index := indexFromRune[r.grid[y][x]]
		index |= 1 << dir
		r.grid[y][x] = runes[index]
		return false
	}
	r.populate2(grid, x-1, y, 3)
	r.populate2(grid, x+1, y, 1)
	r.populate2(grid, x, y-1, 0)
	r.populate2(grid, x, y+1, 2)
	return true
}

func (r *region) calcArea() int {
	running := 0
	for _, row := range r.grid {
		for _, cell := range row {
			if cell == r.identifier {
				running++
			}
		}
	}
	return running
}

func (r *region) calcPerim() int {
	return r.perim
}

func (r *region) calcSides() int {
	running := 0
	for y, row := range r.grid {
		for x, cell := range row {
			if cell == r.identifier || cell == '.' {
				continue
			}
			index := indexFromRune[r.grid[y][x]]

			if index&1 != 0 {
				index -= 1
				r.grid[y][x] = runes[index]
				for x2 := x + 1; x2 < len(r.grid[0]); x2++ {
					newIndex := indexFromRune[r.grid[y][x2]]
					if newIndex&1 != 0 {
						newIndex -= 1
						r.grid[y][x2] = runes[newIndex]
					} else {
						break
					}
				}
				running++
			}
			if index&2 != 0 {
				index -= 2
				r.grid[y][x] = runes[index]
				for y2 := y + 1; y2 < len(r.grid); y2++ {
					newIndex := indexFromRune[r.grid[y2][x]]
					if newIndex&2 != 0 {
						newIndex -= 2
						r.grid[y2][x] = runes[newIndex]
					} else {
						break
					}
				}
				running++
			}
			if index&4 != 0 {
				index -= 4
				r.grid[y][x] = runes[index]
				for x2 := x + 1; x2 < len(r.grid[0]); x2++ {
					newIndex := indexFromRune[r.grid[y][x2]]
					if newIndex&4 != 0 {
						newIndex -= 4
						r.grid[y][x2] = runes[newIndex]
					} else {
						break
					}
				}
				running++
			}
			if index&8 != 0 {
				index -= 8
				r.grid[y][x] = runes[index]
				for y2 := y + 1; y2 < len(r.grid); y2++ {
					newIndex := indexFromRune[r.grid[y2][x]]
					if newIndex&8 != 0 {
						newIndex -= 8
						r.grid[y2][x] = runes[newIndex]
					} else {
						break
					}
				}
				running++
			}
		}
	}
	return running
}

func (r region) print() {
	log.Println(string(r.identifier), r.perim)
	for _, row := range r.grid {
		log.Println(string(row))
	}
}

func sol1() {
	grid := readInput()

	var regions []*region
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '.' {
				emptyGrid := make([][]rune, len(grid))
				for i := range emptyGrid {
					emptyGrid[i] = make([]rune, len(grid[0]))
					for j := range emptyGrid[i] {
						emptyGrid[i][j] = '.'
					}
				}
				r := &region{identifier: grid[y][x], grid: emptyGrid}
				r.populate(grid, x, y)
				// r.print()
				regions = append(regions, r)
			}
		}
	}

	running := 0
	for _, r := range regions {
		running += r.calcArea() * r.calcPerim()
	}

	log.Println(running)
}

func sol2() {
	grid := readInput()

	var regions []*region
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '.' {
				emptyGrid := make([][]rune, len(grid)+2)
				for i := range emptyGrid {
					emptyGrid[i] = make([]rune, len(grid[0])+2)
					for j := range emptyGrid[i] {
						emptyGrid[i][j] = '.'
					}
				}
				r := &region{identifier: grid[y][x], grid: emptyGrid}
				r.populate2(grid, x+1, y+1, -1)
				regions = append(regions, r)
			}
		}
	}

	running := 0
	for _, r := range regions {
		area := r.calcArea()
		sides := r.calcSides()
		running += area * sides
	}

	log.Println(running)
}
