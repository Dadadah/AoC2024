package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func main() {
	sol1()
	sol2()
}

func sol1() {
	ss := readInput()

	running := 0
	for _, s := range ss {
		cursor := 0
		curMul := ""
		for cursor < len(s)-3 {
			curCursor := cursor
			if s[cursor:cursor+4] == "mul(" {
				curMul = "mul("
				cursor += 4
				valid, comma, end := true, false, false
				for valid {
					valid, comma, end = isValidRune(rune(s[cursor]), comma)
					if !valid {
						curMul = ""
					} else {
						curMul += string(s[cursor])
						cursor++
						if end {
							if comma {
								running += runMult(curMul)
							}
							curMul = ""
							break
						}
					}
				}
			}
			cursor = curCursor + 1
		}
	}

	log.Println(running)
}

func sol2() {
	ss := readInput()

	running := 0
	enabled := true
	for _, s := range ss {
		curMul := ""
		for cursor := 0; cursor < len(s)-3; cursor++ {
			if cursor < len(s)-6 && s[cursor:cursor+7] == "don't()" {
				enabled = false
			}
			if !enabled {
				if s[cursor:cursor+4] == "do()" {
					enabled = true
				}
				continue
			}
			if s[cursor:cursor+4] == "mul(" {
				curMul = "mul("
				cursor += 4
				valid, comma, end := true, false, false
				for valid {
					valid, comma, end = isValidRune(rune(s[cursor]), comma)
					if !valid {
						curMul = ""
					} else {
						curMul += string(s[cursor])
						if end {
							if comma {
								running += runMult(curMul)
							}
							curMul = ""
							break
						}
						cursor++
					}
				}
			}
		}
	}

	log.Println(running)
}

var validRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-'}

func isValidRune(r rune, commaFound bool) (bool, bool, bool) {
	if !commaFound {
		if r == ',' {
			return true, true, false
		}
	}
	if r == ')' {
		return true, commaFound, true
	}
	for _, t := range validRunes {
		if r == t {
			return true, commaFound, false
		}
	}
	return false, false, false
}

func runMult(s string) int {
	s = strings.TrimPrefix(s, "mul(")
	s = strings.TrimSuffix(s, ")")
	splits := strings.Split(s, ",")
	val1, err := strconv.Atoi(splits[0])
	if err != nil {
		log.Fatal(err)
	}
	val2, err := strconv.Atoi(splits[1])
	if err != nil {
		log.Fatal(err)
	}
	return val1 * val2
}
