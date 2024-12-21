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
	x, y  int
	score int
}

var gottaCacheTheMaze = make(map[int]map[int]int)

func solveMaze(grid [][]rune, p point) int {
	if gottaCacheTheMaze[p.y][p.x] > 0 {
		return gottaCacheTheMaze[p.y][p.x] + p.score
	}

	lowest, _ := solveMazeAndGetScoreGrid(grid, p)

	if gottaCacheTheMaze[p.y] == nil {
		gottaCacheTheMaze[p.y] = make(map[int]int)
	}
	gottaCacheTheMaze[p.y][p.x] = lowest - p.score

	return lowest
}

func solveMazeAndGetScoreGrid(grid [][]rune, p point) (int, [][]int) {
	var scoreGrid [][]int
	var visitQueue []point

	for y, row := range grid {
		scoreGrid = append(scoreGrid, make([]int, len(row)))
		for x := range row {
			scoreGrid[y][x] = -1
		}
	}

	visitQueue = append(visitQueue, p)

	lowest := -1
	for i := 0; i < len(visitQueue); i++ {
		// moves
		visitX := visitQueue[i].x
		visitY := visitQueue[i].y
		visitScore := visitQueue[i].score

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

		if visitY-1 > 0 {
			if grid[visitY-1][visitX] != '#' {
				// don't consider this path
				if lowest > 0 && visitScore+1 > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX, visitY - 1, visitScore + 1})
			}
		}
		if visitX+1 < len(grid[0])-1 {
			if grid[visitY][visitX+1] != '#' {
				// don't consider this path
				if lowest > 0 && visitScore+1 > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX + 1, visitY, visitScore + 1})
			}
		}
		if visitY+1 < len(grid)-1 {
			if grid[visitY+1][visitX] != '#' {
				// don't consider this path
				if lowest > 0 && visitScore+1 > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX, visitY + 1, visitScore + 1})

			}
		}
		if visitX-1 > 0 {
			if grid[visitY][visitX-1] != '#' {
				// don't consider this path
				if lowest > 0 && visitScore+1 > lowest {
					continue
				}
				visitQueue = append(visitQueue, point{visitX - 1, visitY, visitScore + 1})

			}
		}
	}

	return lowest, scoreGrid
}

func sol1() {
	grid := readInput()
	lowest := 0
	var scoreGrid [][]int
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				lowest, scoreGrid = solveMazeAndGetScoreGrid(grid, point{x, y, 0})
			}
		}
	}

	log.Println(lowest)

	var visitQueue []point
	visitQueue = nil
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				visitQueue = append(visitQueue, point{x, y, 0})
			}
		}
	}

	cheats := 0

	for i := 0; i < len(visitQueue); i++ {
		// moves
		visitX := visitQueue[i].x
		visitY := visitQueue[i].y
		visitScore := visitQueue[i].score

		// If we've visited this cell before and our score is too high, terminate this path
		if visitScore != scoreGrid[visitY][visitX] {
			continue
		}

		if visitY-1 > 0 {
			if grid[visitY-1][visitX] != '#' {
				visitQueue = append(visitQueue, point{visitX, visitY - 1, visitScore + 1})
			} else {
				// cheat
				newGrid := copyGrid(grid)
				newGrid[visitY-1][visitX] = '.'
				score := solveMaze(newGrid, point{visitX, visitY - 1, visitScore + 1})
				if score <= lowest-20 {
					cheats++
				}
			}
		}
		if visitX+1 < len(grid[0])-1 {
			if grid[visitY][visitX+1] != '#' {
				visitQueue = append(visitQueue, point{visitX + 1, visitY, visitScore + 1})
			} else {
				// cheat
				newGrid := copyGrid(grid)
				newGrid[visitY][visitX+1] = '.'
				score := solveMaze(newGrid, point{visitX + 1, visitY, visitScore + 1})
				if score <= lowest-20 {
					cheats++
				}
			}
		}
		if visitY+1 < len(grid)-1 {
			if grid[visitY+1][visitX] != '#' {
				visitQueue = append(visitQueue, point{visitX, visitY + 1, visitScore + 1})
			} else {
				// cheat
				newGrid := copyGrid(grid)
				newGrid[visitY+1][visitX] = '.'
				score := solveMaze(newGrid, point{visitX, visitY + 1, visitScore + 1})
				if score <= lowest-20 {
					cheats++
				}
			}
		}
		if visitX-1 > 0 {
			if grid[visitY][visitX-1] != '#' {
				visitQueue = append(visitQueue, point{visitX - 1, visitY, visitScore + 1})
			} else {
				// cheat
				newGrid := copyGrid(grid)
				newGrid[visitY][visitX-1] = '.'
				score := solveMaze(newGrid, point{visitX - 1, visitY, visitScore + 1})
				if score <= lowest-20 {
					cheats++
				}
			}
		}
	}

	log.Println(cheats)
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for y, row := range grid {
		newGrid[y] = make([]rune, len(row))
		copy(newGrid[y], row)
	}
	return newGrid
}

type gridPoint struct {
	grid  [][]rune
	point point
}

func cheatbfs(grid [][]rune, p point) []point {
	var gps []point
	var solvedPoints []point
	solvedPointsMap := make(map[int]map[int]bool)

	gps = append(gps, p)

	for i := 0; i < 20; i++ {
		var newPoints []point
		for _, gp := range gps {
			// Add new points
			visitX := gp.x
			visitY := gp.y
			visitScore := gp.score
			if visitY-1 > 0 {
				newY := visitY - 1
				newX := visitX
				if solvedPointsMap[newY] == nil || !solvedPointsMap[newY][newX] {
					solvable := grid[newY][newX] != '#'
					newPoint := point{newX, newY, visitScore + 1}
					newPoints = append(newPoints, newPoint)
					if solvable {
						score := solveMaze(grid, newPoint)
						newPoint.score = score
						solvedPoints = append(solvedPoints, newPoint)
					}
					if solvedPointsMap[newY] == nil {
						solvedPointsMap[newY] = map[int]bool{}
					}
					solvedPointsMap[newY][newX] = true
				}
			}
			if visitX+1 < len(grid[0])-1 {
				newY := visitY
				newX := visitX + 1
				if solvedPointsMap[newY] == nil || !solvedPointsMap[newY][newX] {
					solvable := grid[newY][newX] != '#'
					newPoint := point{newX, newY, visitScore + 1}
					newPoints = append(newPoints, newPoint)
					if solvable {
						score := solveMaze(grid, newPoint)
						newPoint.score = score
						solvedPoints = append(solvedPoints, newPoint)
					}
					if solvedPointsMap[newY] == nil {
						solvedPointsMap[newY] = map[int]bool{}
					}
					solvedPointsMap[newY][newX] = true
				}
			}
			if visitY+1 < len(grid)-1 {
				newY := visitY + 1
				newX := visitX
				if solvedPointsMap[newY] == nil || !solvedPointsMap[newY][newX] {
					solvable := grid[newY][newX] != '#'
					newPoint := point{newX, newY, visitScore + 1}
					newPoints = append(newPoints, newPoint)
					if solvable {
						score := solveMaze(grid, newPoint)
						newPoint.score = score
						solvedPoints = append(solvedPoints, newPoint)
					}
					if solvedPointsMap[newY] == nil {
						solvedPointsMap[newY] = map[int]bool{}
					}
					solvedPointsMap[newY][newX] = true
				}
			}
			if visitX-1 > 0 {
				newY := visitY
				newX := visitX - 1
				if solvedPointsMap[newY] == nil || !solvedPointsMap[newY][newX] {
					solvable := grid[newY][newX] != '#'
					newPoint := point{newX, newY, visitScore + 1}
					newPoints = append(newPoints, newPoint)
					if solvable {
						score := solveMaze(grid, newPoint)
						newPoint.score = score
						solvedPoints = append(solvedPoints, newPoint)
					}
					if solvedPointsMap[newY] == nil {
						solvedPointsMap[newY] = map[int]bool{}
					}
					solvedPointsMap[newY][newX] = true
				}
			}
		}
		gps = newPoints
	}
	return solvedPoints
}

func sol2() {
	grid := readInput()
	lowest := 0
	var scoreGrid [][]int
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				lowest, scoreGrid = solveMazeAndGetScoreGrid(grid, point{x, y, 0})
			}
		}
	}

	log.Println(lowest)

	var visitQueue []point
	visitQueue = nil
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
				visitQueue = append(visitQueue, point{x, y, 0})
			}
		}
	}

	cheats := 0

	for i := 0; i < len(visitQueue); i++ {
		// moves
		visitX := visitQueue[i].x
		visitY := visitQueue[i].y
		visitScore := visitQueue[i].score

		// If we've visited this cell before and our score is too high, terminate this path
		if visitScore != scoreGrid[visitY][visitX] {
			continue
		}

		cheatedPoints := cheatbfs(copyGrid(grid), visitQueue[i])
		for _, pe := range cheatedPoints {
			if pe.score <= lowest-100 {
				cheats++
			}
		}

		if visitY-1 > 0 {
			if grid[visitY-1][visitX] != '#' {
				visitQueue = append(visitQueue, point{visitX, visitY - 1, visitScore + 1})
			}
		}
		if visitX+1 < len(grid[0])-1 {
			if grid[visitY][visitX+1] != '#' {
				visitQueue = append(visitQueue, point{visitX + 1, visitY, visitScore + 1})
			}
		}
		if visitY+1 < len(grid)-1 {
			if grid[visitY+1][visitX] != '#' {
				visitQueue = append(visitQueue, point{visitX, visitY + 1, visitScore + 1})
			}
		}
		if visitX-1 > 0 {
			if grid[visitY][visitX-1] != '#' {
				visitQueue = append(visitQueue, point{visitX - 1, visitY, visitScore + 1})
			}
		}
	}

	log.Println(cheats)
}
