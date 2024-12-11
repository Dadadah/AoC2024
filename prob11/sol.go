package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput() (stones []int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		splits := strings.Split(val, " ")
		for _, s := range splits {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			stones = append(stones, num)
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

func applyBlinkRules(s int) []int {
	if s == 0 {
		return []int{1}
	}
	power := 0
	runner := s
	for {
		runner = runner / 10
		power++
		if runner == 0 {
			break
		}
	}
	if power%2 == 0 {
		num2 := s % int(math.Pow10(power/2))
		num1 := s / int(math.Pow10(power/2))
		return []int{num1, num2}
	}
	return []int{s * 2024}
}

func sol1() {
	stones := readInput()
	stoneMap := make(map[int]int)

	for _, s := range stones {
		if _, ok := stoneMap[s]; ok {
			stoneMap[s] += 1
		} else {
			stoneMap[s] = 1
		}
	}

	for i := 0; i < 25; i++ {
		newStoneMap := make(map[int]int)
		for s, c := range stoneMap {
			newStones := applyBlinkRules(s)
			for _, newStone := range newStones {
				if _, ok := newStoneMap[newStone]; ok {
					newStoneMap[newStone] += c
				} else {
					newStoneMap[newStone] = c
				}
			}
		}
		stoneMap = newStoneMap
	}

	running := 0
	for _, s := range stoneMap {
		running += s
	}

	log.Println(running)
}

func sol2() {
	stones := readInput()
	stoneMap := make(map[int]int)

	for _, s := range stones {
		if _, ok := stoneMap[s]; ok {
			stoneMap[s] += 1
		} else {
			stoneMap[s] = 1
		}
	}

	for i := 0; i < 75; i++ {
		newStoneMap := make(map[int]int)
		for s, c := range stoneMap {
			newStones := applyBlinkRules(s)
			for _, newStone := range newStones {
				if _, ok := newStoneMap[newStone]; ok {
					newStoneMap[newStone] += c
				} else {
					newStoneMap[newStone] = c
				}
			}
		}
		stoneMap = newStoneMap
	}

	running := 0
	for _, s := range stoneMap {
		running += s
	}

	log.Println(running)
}
