package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readInput() (towels []string, patterns []string) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readingTowels := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if len(val) == 0 {
			readingTowels = false
			continue
		}

		if readingTowels {
			splits := strings.Split(val, ", ")
			towels = splits
		} else {
			patterns = append(patterns, val)
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

var seen = map[string]map[string]int{}

func solve(towels []string, pattern, towel string) int {
	if val, ok := seen[pattern]; ok {
		if val, ok := val[towel]; ok {
			return val
		}
	}
	if strings.HasPrefix(pattern, towel) {
		newPattern := strings.TrimPrefix(pattern, towel)
		if len(newPattern) == 0 {
			return 1
		}
		found := 0
		for _, t := range towels {
			found += solve(towels, newPattern, t)
		}
		if found > 0 {
			if seen[pattern] == nil {
				seen[pattern] = map[string]int{}
			}
			seen[pattern][towel] = found
		}
		return found
	}
	return 0
}

func sol1() {
	towels, patterns := readInput()

	running := 0
	for _, pattern := range patterns {
		if solve(towels, pattern, "") > 0 {
			running++
		}
	}
	log.Println(running)
}

func sol2() {
	towels, patterns := readInput()

	running := 0
	for _, pattern := range patterns {
		running += solve(towels, pattern, "")
	}
	log.Println(running)
}
