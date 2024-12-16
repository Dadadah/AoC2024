package main

import (
	"bufio"
	"log"
	"os"
)

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

type point struct {
	x, y      int
	direction int
	score     int
}

func sol1() {
	grid := readInput()
	var scoreGrid [][]int

	var visitQueue []point
	for y, row := range grid {
		scoreGrid = append(scoreGrid, make([]int, len(row)))
		for x := range row {
			scoreGrid[y][x] = -1
			if grid[y][x] == 'S' {
				visitQueue = append(visitQueue, point{x, y, 1, 0})
			}
		}
	}

	lowest := -1
	for i := 0; i < len(visitQueue); i++ {
		// moves
		visitX := visitQueue[i].x
		visitY := visitQueue[i].y
		visitScore := visitQueue[i].score
		visitDirection := visitQueue[i].direction

		if grid[visitY][visitX] == 'E' {
			if lowest < 0 || visitScore < lowest {
				lowest = visitScore
				continue
			}
		}

		// If we've visited this cell before and our score is too high, terminate this path
		if scoreGrid[visitY][visitX] > 0 {
			if visitScore < scoreGrid[visitY][visitX] {
				scoreGrid[visitY][visitX] = visitScore
			} else {
				continue
			}
		} else if scoreGrid[visitY][visitX] == -1 {
			scoreGrid[visitY][visitX] = visitScore
		}

		if grid[visitY-1][visitX] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 1:
				fallthrough
			case 3:
				scoreDiff += 1000
				fallthrough
			case 0:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX, visitY - 1, 0, visitScore + scoreDiff})
			}
		}
		if grid[visitY][visitX+1] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 0:
				fallthrough
			case 2:
				scoreDiff += 1000
				fallthrough
			case 1:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX + 1, visitY, 1, visitScore + scoreDiff})
			}
		}
		if grid[visitY+1][visitX] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 1:
				fallthrough
			case 3:
				scoreDiff += 1000
				fallthrough
			case 2:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX, visitY + 1, 2, visitScore + scoreDiff})
			}
		}
		if grid[visitY][visitX-1] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 0:
				fallthrough
			case 2:
				scoreDiff += 1000
				fallthrough
			case 3:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX - 1, visitY, 3, visitScore + scoreDiff})
			}
		}

	}

	log.Println(lowest)
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for y, row := range grid {
		newGrid[y] = make([]rune, len(row))
		copy(newGrid[y], row)
	}
	return newGrid
}

type path struct {
	x         int
	y         int
	direction int
	score     int
	points    []point
}

func (p path) print(grid [][]rune) {
	for _, poi := range p.points {
		grid[poi.y][poi.x] = 'O'
	}
	for _, row := range grid {
		log.Println(string(row))
	}
}

func sol2() {
	grid := readInput()
	var scoreGrid [][]int

	var visitQueue []path
	for y, row := range grid {
		scoreGrid = append(scoreGrid, make([]int, len(row)))
		for x := range row {
			scoreGrid[y][x] = -1
			if grid[y][x] == 'S' {
				visitQueue = append(visitQueue, path{x, y, 1, 0, []point{{x: x, y: y}}})
			}
		}
	}

	lowest := -1
	var lowestPaths []path
	for i := 0; i < len(visitQueue); i++ {
		// moves
		visitX := visitQueue[i].x
		visitY := visitQueue[i].y
		visitScore := visitQueue[i].score
		visitDirection := visitQueue[i].direction
		visitPoints := visitQueue[i].points

		if grid[visitY][visitX] == 'E' {
			if lowest < 0 || visitScore < lowest {
				lowest = visitScore
				lowestPaths = []path{visitQueue[i]}
				continue
			} else if visitScore == lowest {
				lowestPaths = append(lowestPaths, visitQueue[i])
				continue
			}
		}

		// If we've visited this cell before and our score is too high, terminate this path
		if scoreGrid[visitY][visitX] > 0 {
			if visitScore < scoreGrid[visitY][visitX] {
				scoreGrid[visitY][visitX] = visitScore
			} else if visitScore > scoreGrid[visitY][visitX]+1000 {
				continue
			}
		} else if scoreGrid[visitY][visitX] == -1 {
			scoreGrid[visitY][visitX] = visitScore
		}

		if grid[visitY-1][visitX] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 1:
				fallthrough
			case 3:
				scoreDiff += 1000
				fallthrough
			case 0:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				pCopy := make([]point, len(visitPoints))
				copy(pCopy, visitPoints)
				pCopy = append(pCopy, point{x: visitX, y: visitY - 1})
				visitQueue = append(visitQueue, path{visitX, visitY - 1, 0, visitScore + scoreDiff, pCopy})
			}
		}
		if grid[visitY][visitX+1] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 0:
				fallthrough
			case 2:
				scoreDiff += 1000
				fallthrough
			case 1:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				pCopy := make([]point, len(visitPoints))
				copy(pCopy, visitPoints)
				pCopy = append(pCopy, point{x: visitX + 1, y: visitY})
				visitQueue = append(visitQueue, path{visitX + 1, visitY, 1, visitScore + scoreDiff, pCopy})
			}
		}
		if grid[visitY+1][visitX] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 1:
				fallthrough
			case 3:
				scoreDiff += 1000
				fallthrough
			case 2:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				pCopy := make([]point, len(visitPoints))
				copy(pCopy, visitPoints)
				pCopy = append(pCopy, point{x: visitX, y: visitY + 1})
				visitQueue = append(visitQueue, path{visitX, visitY + 1, 2, visitScore + scoreDiff, pCopy})
			}
		}
		if grid[visitY][visitX-1] != '#' {
			scoreDiff := 1
			switch visitDirection {
			case 0:
				fallthrough
			case 2:
				scoreDiff += 1000
				fallthrough
			case 3:
				// don't consider this path
				if lowest > 0 && visitScore+scoreDiff > lowest {
					continue
				}
				pCopy := make([]point, len(visitPoints))
				copy(pCopy, visitPoints)
				pCopy = append(pCopy, point{x: visitX - 1, y: visitY})
				visitQueue = append(visitQueue, path{visitX - 1, visitY, 3, visitScore + scoreDiff, pCopy})
			}
		}

	}

	log.Println(lowest)

	uniques := make(map[int]map[int]struct{})
	for _, p := range lowestPaths {
		for _, poi := range p.points {
			if _, ok := uniques[poi.x]; !ok {
				uniques[poi.x] = make(map[int]struct{})
			}
			uniques[poi.x][poi.y] = struct{}{}
		}
	}

	uniqueCount := 0
	for _, row := range uniques {
		uniqueCount += len(row)
	}

	log.Println(uniqueCount)
}
