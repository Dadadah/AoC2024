package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var gridSize = 71
var dropCount = 1024

type coord struct{ x, y int }

func readInput() (coords []coord) {
	fileName := "input"
	if fileName == "input.test" {
		gridSize = 7
		dropCount = 12
	}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		spits := strings.Split(val, ",")
		x, _ := strconv.Atoi(spits[0])
		y, _ := strconv.Atoi(spits[1])
		coords = append(coords, coord{x, y})
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

type path struct {
	steps []coord
}

func sol1() {
	drops := readInput()

	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i := 0; i < dropCount; i++ {
		drop := drops[i]
		grid[drop.y][drop.x] = '#'
	}
	// bfs
	q := []*path{{steps: []coord{{0, 0}}}}
	for i := 0; i < len(q); i++ {
		// for _, row := range grid {
		// 	log.Println(string(row))
		// }
		curPath := q[i]
		latestStep := curPath.steps[len(curPath.steps)-1]
		if grid[latestStep.y][latestStep.x] == '#' || grid[latestStep.y][latestStep.x] == 'O' {
			continue
		}
		grid[latestStep.y][latestStep.x] = 'O'

		if latestStep.x == gridSize-1 && latestStep.y == gridSize-1 {
			log.Println(len(curPath.steps) - 1)
		}
		if latestStep.y > 0 {
			newSteps := make([]coord, len(curPath.steps))
			copy(newSteps, curPath.steps)
			newSteps = append(newSteps, coord{latestStep.x, latestStep.y - 1})
			p := &path{steps: newSteps}
			q = append(q, p)
		}
		if latestStep.x > 0 {
			newSteps := make([]coord, len(curPath.steps))
			copy(newSteps, curPath.steps)
			newSteps = append(newSteps, coord{latestStep.x - 1, latestStep.y})
			p := &path{steps: newSteps}
			q = append(q, p)
		}
		if latestStep.y < gridSize-1 {
			newSteps := make([]coord, len(curPath.steps))
			copy(newSteps, curPath.steps)
			newSteps = append(newSteps, coord{latestStep.x, latestStep.y + 1})
			p := &path{steps: newSteps}
			q = append(q, p)
		}
		if latestStep.x < gridSize-1 {
			newSteps := make([]coord, len(curPath.steps))
			copy(newSteps, curPath.steps)
			newSteps = append(newSteps, coord{latestStep.x + 1, latestStep.y})
			p := &path{steps: newSteps}
			q = append(q, p)
		}
	}

}

func sol2() {
	drops := readInput()

	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// binary search the bad drop
	upperbound := len(drops) - 1
	lowerbound := 0
	cur := len(drops) / 2
	for upperbound-lowerbound > 1 {
		for i := 0; i < cur; i++ {
			drop := drops[i]
			grid[drop.y][drop.x] = '#'
		}

		// bfs
		exited := false
		q := []*path{{steps: []coord{{0, 0}}}}
		for i := 0; i < len(q); i++ {
			curPath := q[i]
			latestStep := curPath.steps[len(curPath.steps)-1]
			if grid[latestStep.y][latestStep.x] == '#' || grid[latestStep.y][latestStep.x] == 'O' {
				continue
			}
			grid[latestStep.y][latestStep.x] = 'O'

			if latestStep.x == gridSize-1 && latestStep.y == gridSize-1 {
				exited = true
				break
			}
			if latestStep.y > 0 {
				newSteps := make([]coord, len(curPath.steps))
				copy(newSteps, curPath.steps)
				newSteps = append(newSteps, coord{latestStep.x, latestStep.y - 1})
				p := &path{steps: newSteps}
				q = append(q, p)
			}
			if latestStep.x > 0 {
				newSteps := make([]coord, len(curPath.steps))
				copy(newSteps, curPath.steps)
				newSteps = append(newSteps, coord{latestStep.x - 1, latestStep.y})
				p := &path{steps: newSteps}
				q = append(q, p)
			}
			if latestStep.y < gridSize-1 {
				newSteps := make([]coord, len(curPath.steps))
				copy(newSteps, curPath.steps)
				newSteps = append(newSteps, coord{latestStep.x, latestStep.y + 1})
				p := &path{steps: newSteps}
				q = append(q, p)
			}
			if latestStep.x < gridSize-1 {
				newSteps := make([]coord, len(curPath.steps))
				copy(newSteps, curPath.steps)
				newSteps = append(newSteps, coord{latestStep.x + 1, latestStep.y})
				p := &path{steps: newSteps}
				q = append(q, p)
			}
		}
		if !exited {
			log.Println(cur, " failed, setting it to upper bound")
			upperbound = cur
		} else {
			log.Println(cur, " succeeded, setting it to lower bound")
			lowerbound = cur
		}
		for y, row := range grid {
			for x, _ := range row {
				grid[y][x] = '.'
			}
		}
		cur = lowerbound + ((upperbound - lowerbound) / 2)
		log.Println(upperbound, cur, lowerbound)
	}
	log.Println(lowerbound)
	log.Println(drops[lowerbound])
}
