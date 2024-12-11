package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() (ss []string) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		ss = append(ss, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

var validRunes1 = []rune{'X', 'M', 'A', 'S'}

func main() {
	sol1()
	sol2()
}

func sol1() {
	grid := readInput()

	occurences := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'X' {
				for offX := -1; offX <= 1; offX++ {
					for offY := -1; offY <= 1; offY++ {
						if offX == 0 && offY == 0 {
							continue
						}
						for i := 1; i < 4; i++ {
							totalY := offY * i
							totalX := offX * i
							if y+totalY < 0 || y+totalY >= len(grid) {
								break
							}
							if x+totalX < 0 || x+totalX >= len(row) {
								break
							}
							if rune(grid[y+totalY][x+totalX]) != validRunes1[i] {
								break
							}
							if i == 3 {
								occurences++
							}
						}
					}
				}

			}
		}
	}
	log.Println(occurences)
}

func sol2() {
	grid := readInput()

	occurences := 0
	for y, row := range grid {
		if y == 0 || y == len(grid)-1 {
			continue
		}
		for x, cell := range row {
			if x == 0 || x == len(row)-1 {
				continue
			}
			if cell == 'A' {
				ms, ss := 0, 0
				mx, my := 0, 0
				valid := true
				for offX := -1; offX <= 1; offX++ {
					for offY := -1; offY <= 1; offY++ {
						if offX == 0 || offY == 0 {
							continue
						}
						if grid[y+offY][x+offX] == 'M' {
							ms++
							if mx != 0 {
								if mx != x+offX && my != y+offY {
									valid = false
								}
							} else {
								mx = x + offX
								my = y + offY
							}
						}
						if grid[y+offY][x+offX] == 'S' {
							ss++
						}
					}
				}
				if ms == 2 && ss == 2 && valid {
					occurences++
				}

			}
		}
	}
	log.Println(occurences)
}
