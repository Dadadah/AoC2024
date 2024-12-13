package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var s tcell.Screen

const drawDelay = 200
const drawLoops = 1

type region struct {
	grid       [][]rune
	identifier rune
	perim      int
	sides      int
}

func readInput() (grid [][]rune) {
	f, err := os.Open("input.test")
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

func drawText(x, y int, style tcell.Style, text string) {
	drawTextNoSleep(x, y, style, text)
	time.Sleep(time.Second)
}

func drawTextNoSleep(x, y int, style tcell.Style, text string) {
	for _, r := range []rune(text) {
		s.SetContent(x, y, r, nil, style)
		x++
	}
	s.Show()
}

func drawGrid(grid [][]rune, blinkX, blinkY, offset int, style tcell.Style, wait bool) {
	old := grid[blinkY][blinkX]

	loops := drawLoops
	if !wait {
		loops = 1
	}

	for i := 0; i < loops; i++ {
		if grid[blinkY][blinkX] == '_' {
			grid[blinkY][blinkX] = old
		} else {
			grid[blinkY][blinkX] = '_'
		}
		for y, row := range grid {
			for x, cell := range row {
				s.SetContent(x+((5+len(row))*offset), y, cell, nil, style)
			}

		}
		s.Show()
		if wait {
			time.Sleep(drawDelay * time.Millisecond)
		}
	}

	grid[blinkY][blinkX] = old
	s.SetContent(blinkX+((5+len(grid[0]))*offset), blinkY, old, nil, style)
	s.Show()
}

func main() {
	var err error
	s, err = tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					quit()
					return
				}
			}
		}
	}()

	for {
		ev := s.PollEvent()
		closeout := false
		switch ev.(type) {
		case *tcell.EventKey:
			closeout = true
		}
		if closeout {
			break
		}
	}

	// sol1()
	sol2()

	quit()
}

func quit() {
	s.Fini()
	os.Exit(0)
}

func (r *region) populate(grid [][]rune, x, y int) bool {
	if r.grid[y][x] == r.identifier {
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
		return true
	}
	drawGrid(grid, x, y, 0, tcell.StyleDefault, true)
	if grid[y][x] == r.identifier {
		grid[y][x] = '.'
		r.grid[y][x] = r.identifier
		drawGrid(grid, x, y, 0, tcell.StyleDefault, false)
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
	} else {
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
		if r.grid[y][x] == '.' {
			r.grid[y][x] = '1'
		} else {
			r.grid[y][x]++
		}
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, false)
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

var runes = []rune{'0', '_', '[', 0x231E, 0x203E, '=', 0x231C, 0x228F, ']', 0x231F, 0x2016, 'u', 0x231D, 0x2290, 'n', 0x03F4}

var indexFromRune = map[rune]int{'0': 0, '_': 1, '[': 2, 0x231E: 3, 0x203E: 4, '=': 5, 0x231C: 6, 0x228F: 7, ']': 8, 0x231F: 9, 0x2016: 10, 'u': 11, 0x231D: 12, 0x2290: 13, 'n': 14, 0x03F4: 15}

func (r *region) populate2(grid [][]rune, x, y int, dir int) bool {
	baseGridX := x - 1
	baseGridY := y - 1
	if r.grid[y][x] == r.identifier {
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
		return true
	}
	if baseGridX < 0 || baseGridX >= len(grid[0]) || baseGridY < 0 || baseGridY >= len(grid) {
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
		if r.grid[y][x] == '.' {
			r.grid[y][x] = runes[1<<dir]
		}
		index := indexFromRune[r.grid[y][x]]
		index |= 1 << dir
		r.grid[y][x] = runes[index]
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, false)
		return false
	}
	drawGrid(grid, baseGridX, baseGridY, 0, tcell.StyleDefault, true)
	if grid[baseGridY][baseGridX] == r.identifier {
		grid[baseGridY][baseGridX] = '.'
		r.grid[y][x] = r.identifier
		drawGrid(grid, baseGridX, baseGridY, 0, tcell.StyleDefault, false)
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
	} else {
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
		if r.grid[y][x] == '.' {
			r.grid[y][x] = runes[1<<dir]
		}
		index := indexFromRune[r.grid[y][x]]
		index |= 1 << dir
		r.grid[y][x] = runes[index]
		drawGrid(r.grid, x, y, 1, tcell.StyleDefault, false)
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

func (r *region) calcScore1() int {
	return r.calcArea() * r.calcPerim()
}

func (r *region) calcScore2() int {
	return r.calcArea() * r.calcSides()
}

func (r *region) calcSides() int {
	if r.sides > 0 {
		return r.sides
	}
	running := 0
	for y, row := range r.grid {
		for x, cell := range row {
			if cell == r.identifier || cell == '.' || cell == '0' {
				continue
			}
			drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
			index := indexFromRune[r.grid[y][x]]

			if index&1 != 0 {
				index -= 1
				r.grid[y][x] = runes[index]
				for x2 := x + 1; x2 < len(r.grid[0]); x2++ {
					drawGrid(r.grid, x2, y, 1, tcell.StyleDefault, true)
					newIndex := indexFromRune[r.grid[y][x2]]
					if newIndex&1 != 0 {
						newIndex -= 1
						r.grid[y][x2] = runes[newIndex]
						drawGrid(r.grid, x2, y, 1, tcell.StyleDefault, true)
					} else {
						break
					}
				}
				running++
			}
			if index&2 != 0 {
				drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
				index -= 2
				r.grid[y][x] = runes[index]
				for y2 := y + 1; y2 < len(r.grid); y2++ {
					drawGrid(r.grid, x, y2, 1, tcell.StyleDefault, true)
					newIndex := indexFromRune[r.grid[y2][x]]
					if newIndex&2 != 0 {
						newIndex -= 2
						r.grid[y2][x] = runes[newIndex]
						drawGrid(r.grid, x, y2, 1, tcell.StyleDefault, true)
					} else {
						break
					}
				}
				running++
			}
			if index&4 != 0 {
				drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
				index -= 4
				r.grid[y][x] = runes[index]
				for x2 := x + 1; x2 < len(r.grid[0]); x2++ {
					drawGrid(r.grid, x2, y, 1, tcell.StyleDefault, true)
					newIndex := indexFromRune[r.grid[y][x2]]
					if newIndex&4 != 0 {
						newIndex -= 4
						r.grid[y][x2] = runes[newIndex]
						drawGrid(r.grid, x2, y, 1, tcell.StyleDefault, true)
					} else {
						break
					}
				}
				running++
			}
			if index&8 != 0 {
				drawGrid(r.grid, x, y, 1, tcell.StyleDefault, true)
				index -= 8
				r.grid[y][x] = runes[index]
				for y2 := y + 1; y2 < len(r.grid); y2++ {
					drawGrid(r.grid, x, y2, 1, tcell.StyleDefault, true)
					newIndex := indexFromRune[r.grid[y2][x]]
					if newIndex&8 != 0 {
						newIndex -= 8
						r.grid[y2][x] = runes[newIndex]
						drawGrid(r.grid, x, y2, 1, tcell.StyleDefault, true)
					} else {
						break
					}
				}
				running++
			}
		}
	}
	r.sides = running
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

	drawTextNoSleep(0, len(grid)+1, tcell.StyleDefault, "Region | Area | Perimeter | Score")
	drawTextNoSleep(0, len(grid)+2, tcell.StyleDefault, "------ | ---- | --------- | -----")

	var regions []*region
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '.' {
				drawGrid(grid, x, y, 0, tcell.StyleDefault, true)
				emptyGrid := make([][]rune, len(grid))
				for i := range emptyGrid {
					emptyGrid[i] = make([]rune, len(grid[0]))
					for j := range emptyGrid[i] {
						emptyGrid[i][j] = '.'
					}
				}
				r := &region{identifier: grid[y][x], grid: emptyGrid}
				r.populate(grid, x, y)
				drawText(0, len(grid)+3+len(regions), tcell.StyleDefault, fmt.Sprintf("%6s | %4d | %9d | %d", string(r.identifier), r.calcArea(), r.calcPerim(), r.calcScore1()))
				regions = append(regions, r)
			}
		}
	}

	running := 0
	for _, r := range regions {
		running += r.calcScore1()
	}

}

func sol2() {
	s.Clear()
	grid := readInput()

	drawTextNoSleep(0, len(grid)+3, tcell.StyleDefault, "Region | Area | Sides | Score")
	drawTextNoSleep(0, len(grid)+4, tcell.StyleDefault, "------ | ---- | ----- | -----")

	var regions []*region
	for y, row := range grid {
		for x := range row {
			if grid[y][x] != '.' {
				drawGrid(grid, x, y, 0, tcell.StyleDefault, true)
				emptyGrid := make([][]rune, len(grid)+2)
				for i := range emptyGrid {
					emptyGrid[i] = make([]rune, len(grid[0])+2)
					for j := range emptyGrid[i] {
						emptyGrid[i][j] = '.'
					}
				}
				r := &region{identifier: grid[y][x], grid: emptyGrid}
				r.populate2(grid, x+1, y+1, -1)
				drawText(0, len(grid)+5+len(regions), tcell.StyleDefault, fmt.Sprintf("%6s | %4d | %5d | %d", string(r.identifier), r.calcArea(), r.calcSides(), r.calcScore2()))
				regions = append(regions, r)
			}
		}
	}

	running := 0
	for _, r := range regions {
		running += r.calcScore2()
	}
}
