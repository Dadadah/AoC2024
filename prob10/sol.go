package main

import (
	"bufio"
	"log"
	"os"
)

var runeToInt = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func readInput() (ss [][]int) {
	f, err := os.Open("input.test")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		var row []int

		for _, v := range val {
			row = append(row, runeToInt[v])
		}
		ss = append(ss, row)
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

func bfs(grid [][]int, base, x, y int) (nines map[int]map[int]struct{}) {
	if base == 9 {
		return map[int]map[int]struct{}{x: {y: struct{}{}}}
	}
	nines = map[int]map[int]struct{}{}
	for offY := -1; offY <= 1; offY++ {
		if y+offY < 0 || y+offY >= len(grid) {
			continue
		}
		for offX := -1; offX <= 1; offX++ {
			if x+offX < 0 || x+offX >= len(grid[y+offY]) {
				continue
			}
			if offY == 0 && offX == 0 {
				continue
			}

			if offY != 0 && offX != 0 {
				continue
			}

			if grid[y+offY][x+offX] == base+1 {
				newNines := bfs(grid, base+1, x+offX, y+offY)
				for nx, ys := range newNines {
					if _, ok := nines[nx]; !ok {
						nines[nx] = make(map[int]struct{})
					}
					for ny := range ys {
						nines[nx][ny] = struct{}{}
					}
				}
			}
		}
	}
	return nines
}

func bfs2(grid [][]int, base, x, y int) int {
	if base == 9 {
		return 1
	}
	nines := 0
	for offY := -1; offY <= 1; offY++ {
		if y+offY < 0 || y+offY >= len(grid) {
			continue
		}
		for offX := -1; offX <= 1; offX++ {
			if x+offX < 0 || x+offX >= len(grid[y+offY]) {
				continue
			}
			if offY == 0 && offX == 0 {
				continue
			}

			if offY != 0 && offX != 0 {
				continue
			}

			if grid[y+offY][x+offX] == base+1 {
				count := bfs2(grid, base+1, x+offX, y+offY)
				nines += count
			}
		}
	}
	return nines
}

func sol1() {
	grid := readInput()

	running := 0
	for y, row := range grid {
		for x, val := range row {
			if val == 0 {
				nines := bfs(grid, 0, x, y)
				for _, ys := range nines {
					for _ = range ys {
						running++
					}
				}
			}
		}
	}
	log.Println(running)
}

func sol2() {
	grid := readInput()

	running := 0
	for y, row := range grid {
		for x, val := range row {
			if val == 0 {
				nines := bfs2(grid, 0, x, y)
				running += nines
			}
		}
	}
	log.Println(running)
}
