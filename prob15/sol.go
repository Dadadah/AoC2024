package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() (grid [][]rune, instructions []rune) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readingInstructions := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if val == "" {
			readingInstructions = true
			continue
		}

		if !readingInstructions {
			grid = append(grid, []rune(val))
		} else {
			instructions = append(instructions, []rune(val)...)
		}
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

func sol1() {
	grid, instructions := readInput()

	robotx := 0
	roboty := 0

	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				robotx = x
				roboty = y
				break
			}
		}
	}

	for _, instruction := range instructions {
		switch instruction {
		case '^':
			if grid[roboty-1][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty-1][robotx] = '@'
				roboty--
				break
			}
			if grid[roboty-1][robotx] == '#' {
				break
			}
			i := 1
			for ; grid[roboty-i][robotx] == 'O'; i++ {
			}
			if grid[roboty-i][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty-1][robotx] = '@'
				grid[roboty-i][robotx] = 'O'
				roboty--
				break
			}
			if grid[roboty-i][robotx] == '#' {
				break
			}
		case '>':
			if grid[roboty][robotx+1] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx+1] = '@'
				robotx++
				break
			}
			if grid[roboty][robotx+1] == '#' {
				break
			}
			i := 1
			for ; grid[roboty][robotx+i] == 'O'; i++ {
			}
			if grid[roboty][robotx+i] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx+1] = '@'
				grid[roboty][robotx+i] = 'O'
				robotx++
				break
			}
			if grid[roboty][robotx+i] == '#' {
				break
			}
		case 'v':
			if grid[roboty+1][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty+1][robotx] = '@'
				roboty++
				break
			}
			if grid[roboty+1][robotx] == '#' {
				break
			}
			i := 1
			for ; grid[roboty+i][robotx] == 'O'; i++ {
			}
			if grid[roboty+i][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty+1][robotx] = '@'
				grid[roboty+i][robotx] = 'O'
				roboty++
				break
			}
			if grid[roboty+i][robotx] == '#' {
				break
			}
		case '<':
			if grid[roboty][robotx-1] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx-1] = '@'
				robotx--
				break
			}
			if grid[roboty][robotx-1] == '#' {
				break
			}
			i := 1
			for ; grid[roboty][robotx-i] == 'O'; i++ {
			}
			if grid[roboty][robotx-i] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx-1] = '@'
				grid[roboty][robotx-i] = 'O'
				robotx--
				break
			}
			if grid[roboty][robotx-i] == '#' {
				break
			}
		}
	}

	score := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' {
				score += y*100 + x
			}
		}
	}
	log.Println(score)
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for y, row := range grid {
		newGrid[y] = make([]rune, len(row))
		copy(newGrid[y], row)
	}
	return newGrid
}

func move(grid [][]rune, beingMoved rune, dir, x, y int) bool {
	if grid[y][x] == '#' {
		return false
	}

	if grid[y][x] == '.' {
		grid[y][x] = beingMoved
		return true
	}

	succeeded := true
	succeeded = succeeded && move(grid, grid[y][x], dir, x, y+dir)
	if grid[y][x] == '[' {
		succeeded = succeeded && move(grid, ']', dir, x+1, y+dir)
		grid[y][x+1] = '.'
	}
	if grid[y][x] == ']' {
		succeeded = succeeded && move(grid, '[', dir, x-1, y+dir)
		grid[y][x-1] = '.'
	}
	grid[y][x] = beingMoved
	return succeeded
}

func sol2() {
	g, instructions := readInput()

	grid := make([][]rune, len(g))

	for y, row := range g {
		grid[y] = make([]rune, len(row)*2)
		for x, cell := range row {
			if cell == '#' {
				grid[y][2*x] = '#'
				grid[y][2*x+1] = '#'
			} else if cell == '.' {
				grid[y][2*x] = '.'
				grid[y][2*x+1] = '.'
			} else if cell == '@' {
				grid[y][2*x] = '@'
				grid[y][2*x+1] = '.'
			} else if cell == 'O' {
				grid[y][2*x] = '['
				grid[y][2*x+1] = ']'
			}
		}
	}

	robotx := 0
	roboty := 0

	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				robotx = x
				roboty = y
				break
			}
		}
	}

	for _, instruction := range instructions {
		switch instruction {
		case '^':
			if grid[roboty-1][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty-1][robotx] = '@'
				roboty--
				break
			}
			if grid[roboty-1][robotx] == '#' {
				break
			}
			gridCopy := copyGrid(grid)
			moved := move(gridCopy, grid[roboty-1][robotx], -1, robotx, roboty-1)
			if moved {
				grid = gridCopy
				grid[roboty][robotx] = '.'
				grid[roboty-1][robotx] = '@'
				roboty--
				break
			}
		case '>':
			if grid[roboty][robotx+1] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx+1] = '@'
				robotx++
				break
			}
			if grid[roboty][robotx+1] == '#' {
				break
			}
			i := 1
			for ; grid[roboty][robotx+i] == '['; i += 2 {
			}
			if grid[roboty][robotx+i] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx+1] = '@'
				cur := '['
				for j := 1; j < i; j++ {
					grid[roboty][robotx+1+j] = cur
					if cur == '[' {
						cur = ']'
					} else {
						cur = '['
					}
				}
				robotx++
				break
			}
			if grid[roboty][robotx+i] == '#' {
				break
			}
		case 'v':
			if grid[roboty+1][robotx] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty+1][robotx] = '@'
				roboty++
				break
			}
			if grid[roboty+1][robotx] == '#' {
				break
			}
			gridCopy := copyGrid(grid)
			moved := move(gridCopy, grid[roboty+1][robotx], 1, robotx, roboty+1)
			if moved {
				grid = gridCopy
				grid[roboty][robotx] = '.'
				grid[roboty+1][robotx] = '@'
				roboty++
				break
			}
		case '<':
			if grid[roboty][robotx-1] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx-1] = '@'
				robotx--
				break
			}
			if grid[roboty][robotx-1] == '#' {
				break
			}
			i := 1
			for ; grid[roboty][robotx-i] == ']'; i += 2 {
			}
			if grid[roboty][robotx-i] == '.' {
				grid[roboty][robotx] = '.'
				grid[roboty][robotx-1] = '@'
				cur := ']'
				for j := 1; j < i; j++ {
					grid[roboty][robotx-(j+1)] = cur
					if cur == ']' {
						cur = '['
					} else {
						cur = ']'
					}
				}
				robotx--
				break
			}
			if grid[roboty][robotx-i] == '#' {
				break
			}
		}
	}

	score := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '[' {
				score += y*100 + x
			}
		}
	}
	log.Println(score)

}
