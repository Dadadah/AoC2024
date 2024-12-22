package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput() (codes [][]rune) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		codes = append(codes, []rune(val))
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

var doorPad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'#', '0', 'A'},
}

var robotPad = [][]rune{
	{'#', '^', 'A'},
	{'<', 'v', '>'},
}

var shortestCodePathMap = make(map[rune]map[rune][]rune)

func shortestDoorCodePath(a, b rune) []rune {
	if shortestCodePathMap[a] != nil {
		if shortestCodePathMap[a][b] != nil {
			return shortestCodePathMap[a][b]
		}
	} else {
		shortestCodePathMap[a] = make(map[rune][]rune)
	}

	aX := 0
	aY := 0
	bX := 0
	bY := 0

	for y, row := range doorPad {
		for x, button := range row {
			if button == a {
				aX = x
				aY = y
			}
			if button == b {
				bX = x
				bY = y
			}
		}
	}

	var path []rune

	if (aY == 3 && bX == 0) || (aX == 0 && bY == 3) {
		// do y first
		for diff := aY - bY; diff != 0; diff = aY - bY {
			if diff > 0 {
				path = append(path, '^')
				aY--
			} else {
				path = append(path, 'v')
				aY++
			}
		}
	}

	for diff := aX - bX; diff > 0; diff = aX - bX {
		path = append(path, '<')
		aX--
	}

	for diff := aY - bY; diff != 0; diff = aY - bY {
		if diff > 0 {
			path = append(path, '^')
			aY--
		} else {
			path = append(path, 'v')
			aY++
		}
	}

	for diff := aX - bX; diff < 0; diff = aX - bX {
		path = append(path, '>')
		aX++
	}

	path = append(path, 'A')
	shortestCodePathMap[a][b] = path
	return path
}

func shortestRobotCodePath(a, b rune) []rune {
	if shortestCodePathMap[a] != nil {
		if shortestCodePathMap[a][b] != nil {
			return shortestCodePathMap[a][b]
		}
	} else {
		shortestCodePathMap[a] = make(map[rune][]rune)
	}

	aX := 0
	aY := 0
	bX := 0
	bY := 0

	for y, row := range robotPad {
		for x, button := range row {
			if button == a {
				aX = x
				aY = y
			}
			if button == b {
				bX = x
				bY = y
			}
		}
	}

	var path []rune

	if aY == 0 && bX == 0 {
		// do y first
		for diff := aY - bY; diff != 0; diff = aY - bY {
			if diff > 0 {
				path = append(path, '^')
				aY--
			} else {
				path = append(path, 'v')
				aY++
			}
		}
	} else if aY == 1 && bY == 0 {
		// do x first
		for diff := aX - bX; diff > 0; diff = aX - bX {
			path = append(path, '<')
			aX--
		}

		for diff := aX - bX; diff < 0; diff = aX - bX {
			path = append(path, '>')
			aX++
		}
	}

	for diff := aX - bX; diff > 0; diff = aX - bX {
		path = append(path, '<')
		aX--
	}

	for diff := aY - bY; diff != 0; diff = aY - bY {
		if diff > 0 {
			path = append(path, '^')
			aY--
		} else {
			path = append(path, 'v')
			aY++
		}
	}

	for diff := aX - bX; diff < 0; diff = aX - bX {
		path = append(path, '>')
		aX++
	}

	path = append(path, 'A')
	shortestCodePathMap[a][b] = path
	return path
}

func sol1() {
	ogCodes := readInput()

	var codes [][]rune
	var results [][]rune
	for _, c := range ogCodes {
		codes = append(codes, c)
	}

	for _, code := range codes {
		lastCode := 'A'
		var codePath []rune
		for _, r := range code {
			cp := shortestDoorCodePath(lastCode, r)
			codePath = append(codePath, cp...)
			lastCode = r
		}
		code = codePath
		for i := 0; i < 2; i++ {
			codePath = nil
			for _, r := range code {
				cp := shortestRobotCodePath(lastCode, r)
				codePath = append(codePath, cp...)
				lastCode = r
			}
			code = codePath
		}
		results = append(results, code)
	}

	running := 0
	for i, code := range results {
		ogCode := ogCodes[i]
		ogCode = ogCode[:len(ogCode)-1]
		val, err := strconv.Atoi(string(ogCode))
		if err != nil {
			log.Fatal(err)
		}
		running += val * len(code)
	}
	log.Println(running)
}

type codeSection struct {
	val          []rune
	nextSections []*codeSection
}

func (cs *codeSection) count(depth int) int {
	if depth == 1 {
		return len(cs.val)
	}
	total := 0
	for _, ns := range cs.nextSections {
		total += ns.count(depth - 1)
	}
	return total
}

func (cs *codeSection) getAsRunes(depth int) []rune {
	if depth == 1 {
		return cs.val
	}
	var str []rune
	for _, ns := range cs.nextSections {
		str = append(str, ns.getAsRunes(depth-1)...)
	}
	return str
}

var codeSectionMap = make(map[string]*codeSection)

func getOrCreateCodeSection(val []rune) *codeSection {
	if codeSectionMap[string(val)] == nil {
		codeSectionMap[string(val)] = &codeSection{val: val}
	}
	return codeSectionMap[string(val)]
}

func sol2() {
	codes := readInput()

	running := 0

	for _, code := range codes {
		lastCode := 'A'
		var newCode []rune
		for _, r := range code {
			cp := shortestDoorCodePath(lastCode, r)
			newCode = append(newCode, cp...)
			lastCode = r
		}

		var cp []rune
		var cs *codeSection
		var allCS []*codeSection
		cp = shortestRobotCodePath('A', '<')
		log.Println('A', '<', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('A', 'v')
		log.Println('A', 'v', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('A', '^')
		log.Println('A', '^', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('A', '>')
		log.Println('A', '>', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('A', 'A')
		log.Println('A', 'A', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('^', '<')
		log.Println('^', '<', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('^', 'v')
		log.Println('^', 'v', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('^', 'A')
		log.Println('^', 'A', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('^', '>')
		log.Println('^', '>', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('^', '^')
		log.Println('^', '^', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('>', '<')
		log.Println('>', '<', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('>', 'v')
		log.Println('>', 'v', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('>', '^')
		log.Println('>', '^', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('>', 'A')
		log.Println('>', 'A', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('>', '>')
		log.Println('>', '>', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('<', 'A')
		log.Println('<', 'A', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('<', 'v')
		log.Println('<', 'v', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('<', '^')
		log.Println('<', '^', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('<', '>')
		log.Println('<', '>', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)
		cp = shortestRobotCodePath('<', '<')
		log.Println('<', '<', string(cp))
		cs = getOrCreateCodeSection(cp)
		allCS = append(allCS, cs)

		for _, cs := range allCS {
			lastCode := 'A'
			for j, r := range cs.val {
				var cs2 *codeSection
				cp := shortestRobotCodePath(lastCode, r)
				cs2 = getOrCreateCodeSection(cp)
				if cs.nextSections == nil {
					cs.nextSections = make([]*codeSection, len(cs.val))
				}
				cs.nextSections[j] = cs2
				lastCode = r
			}
		}

		log.Println(string(newCode))

		lastCode = 'A'
		var newCodeSections []*codeSection
		for _, r := range newCode {
			cp := shortestRobotCodePath(lastCode, r)
			cs2 := getOrCreateCodeSection(cp)
			newCodeSections = append(newCodeSections, cs2)
			lastCode = r
		}

		r := 0
		for _, cs := range newCodeSections {
			r += cs.count(25)
		}
		log.Println(r)

		numCode := code[:len(code)-1]
		val, err := strconv.Atoi(string(numCode))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(val, r)
		running += val * r
	}
	log.Println(running)
}
