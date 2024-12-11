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

func main() {
	sol1()
	sol2()
}

func sol1() {
	grid := readInput()

	guardX, guardY := 0, 0

	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				guardX, guardY = x, y
				break
			}
		}
		if guardX != 0 && guardY != 0 {
			break
		}
	}

	grid[guardY][guardX] = 'X'

	dir := 0
	for guardX >= 0 && guardX <= len(grid[0])-1 && guardY >= 0 && guardY <= len(grid)-1 {
		turned := false
		switch dir {
		case 0:
			if guardY-1 >= 0 {
				if grid[guardY-1][guardX] == '#' {
					dir++
					turned = true
				}
			}
		case 1:
			if guardX+1 <= len(grid[0])-1 {
				if grid[guardY][guardX+1] == '#' {
					dir++
					turned = true
				}
			}
		case 2:
			if guardY+1 <= len(grid)-1 {
				if grid[guardY+1][guardX] == '#' {
					dir++
					turned = true
				}
			}
		case 3:
			if guardX-1 >= 0 {
				if grid[guardY][guardX-1] == '#' {
					dir = 0
					turned = true
				}
			}
		}
		grid[guardY][guardX] = 'X'
		if turned {
			continue
		}
		switch dir {
		case 0:
			guardY--
		case 1:
			guardX++
		case 2:
			guardY++
		case 3:
			guardX--
		}
	}
	running := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 'X' {
				running++
			}
		}
	}
	log.Println(running)
}

func sol2() {
	grid := readInput()

	initX, initY := 0, 0
	guardX, guardY := 0, 0

	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				guardX, guardY = x, y
				initX, initY = x, y
				break
			}
		}
		if guardX != 0 && guardY != 0 {
			break
		}
	}

	grid[guardY][guardX] = 'X'

	running := 0
	dir := 0
	for guardX >= 0 && guardX <= len(grid[0])-1 && guardY >= 0 && guardY <= len(grid)-1 {
		turned := false
		switch dir {
		case 0:
			if guardY-1 >= 0 {
				if grid[guardY-1][guardX] == '#' {
					dir++
					turned = true
				} else if !(initX == guardX && initY == guardY-1) {
					if grid[guardY-1][guardX] != 'O' {
						grid[guardY-1][guardX] = '#'
						if loopCheck(grid, initX, initY) {
							running++
						}
						grid[guardY-1][guardX] = 'O'
					}
				}
			}
		case 1:
			if guardX+1 <= len(grid[0])-1 {
				if grid[guardY][guardX+1] == '#' {
					dir++
					turned = true
				} else if !(initX == guardX+1 && initY == guardY) {
					if grid[guardY][guardX+1] != 'O' {
						grid[guardY][guardX+1] = '#'
						if loopCheck(grid, initX, initY) {
							running++
						}
						grid[guardY][guardX+1] = 'O'
					}
				}
			}

		case 2:
			if guardY+1 <= len(grid)-1 {
				if grid[guardY+1][guardX] == '#' {
					dir++
					turned = true
				} else if !(initX == guardX && initY == guardY+1) {
					if grid[guardY+1][guardX] != 'O' {
						grid[guardY+1][guardX] = '#'
						if loopCheck(grid, initX, initY) {
							running++
						}
						grid[guardY+1][guardX] = 'O'
					}
				}
			}

		case 3:
			if guardX-1 >= 0 {
				if grid[guardY][guardX-1] == '#' {
					dir = 0
					turned = true
				} else if !(initX == guardX-1 && initY == guardY) {
					if grid[guardY][guardX-1] != 'O' {
						grid[guardY][guardX-1] = '#'
						if loopCheck(grid, initX, initY) {
							running++
						}
						grid[guardY][guardX-1] = 'O'
					}
				}
			}
		}
		if turned {
			continue
		}
		switch dir {
		case 0:
			guardY--
		case 1:
			guardX++
		case 2:
			guardY++
		case 3:
			guardX--
		}
	}
	log.Println(running)
}

func loopCheck(grid [][]rune, guardX, guardY int) bool {
	check, dirs := loopCheckNoPrint(grid, guardX, guardY)
	if check {
		printIt(grid, dirs)
	}
	return check
}
func loopCheckNoPrint(grid [][]rune, guardX, guardY int) (bool, [][]int) {
	dir := 0
	directions := make([][]int, len(grid))
	for guardX >= 0 && guardX <= len(grid[0])-1 && guardY >= 0 && guardY <= len(grid)-1 {
		turned := false
		switch dir {
		case 0:
			if guardY-1 >= 0 {
				if directions[guardY-1] != nil {
					if directions[guardY-1][guardX]&(1<<dir) != 0 {
						return true, directions
					}
				}
				if grid[guardY-1][guardX] == '#' {
					dir++
					turned = true
				}
			}
		case 1:
			if guardX+1 <= len(grid[0])-1 {
				if directions[guardY] != nil {
					if directions[guardY][guardX+1]&(1<<dir) != 0 {
						return true, directions
					}
				}
				if grid[guardY][guardX+1] == '#' {
					dir++
					turned = true
				}
			}
		case 2:
			if guardY+1 <= len(grid)-1 {
				if directions[guardY+1] != nil {
					if directions[guardY+1][guardX]&(1<<dir) != 0 {
						return true, directions
					}
				}
				if grid[guardY+1][guardX] == '#' {
					dir++
					turned = true
				}
			}
		case 3:
			if guardX-1 >= 0 {
				if directions[guardY] != nil {
					if directions[guardY][guardX-1]&(1<<dir) != 0 {
						return true, directions
					}
				}
				if grid[guardY][guardX-1] == '#' {
					dir = 0
					turned = true
				}
			}
		}
		if directions[guardY] == nil {
			directions[guardY] = make([]int, len(grid[0]))
		}
		directions[guardY][guardX] |= 1 << dir
		if turned {
			directions[guardY][guardX] |= 1 << ((dir + 3) % 4)
			continue
		}
		switch dir {
		case 0:
			guardY--
		case 1:
			guardX++
		case 2:
			guardY++
		case 3:
			guardX--
		}
	}
	return false, directions
}

func printIt(grid [][]rune, directions [][]int) {
	if true {
		return
	}
	for y, row := range directions {
		toPrint := ""
		if row == nil {
			row = make([]int, len(grid[0]))
		}
		for x, dirs := range row {
			p := ""
			if dirs&1 != 0 || dirs&4 != 0 {
				if dirs&2 != 0 || dirs&8 != 0 {
					p = "+"
				} else {
					p = "|"
				}
			} else {
				if dirs&2 != 0 || dirs&8 != 0 {
					p = "-"
				} else {
					p = string(grid[y][x])
				}
			}
			toPrint += p
		}
		log.Println(toPrint)
	}
	log.Println()
}
