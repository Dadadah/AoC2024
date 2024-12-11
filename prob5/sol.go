package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() (mustBefore map[int][]int, updates [][]int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mustBefore = make(map[int][]int)
	sw := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if val == "" {
			sw = true
			continue
		}

		if !sw {
			splits := strings.Split(val, "|")
			num1, err := strconv.Atoi(splits[0])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(splits[1])
			if err != nil {
				log.Fatal(err)
			}
			mustBefore[num1] = append(mustBefore[num1], num2)
		} else {
			splits := strings.Split(val, ",")
			var nums []int
			for _, split := range splits {
				num, err := strconv.Atoi(split)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, num)
			}
			updates = append(updates, nums)
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

func sol1() {
	mustBefore, updates := readInput()

	running := 0
	for _, update := range updates {
		valid := true
		mid := 0
		for i, num := range update {
			if i == len(update)/2 {
				mid = num
			}

			mustBeBefore := mustBefore[num]
			for _, t := range mustBeBefore {
				for _, b := range update[:i] {
					if t == b {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			running += mid
		}
	}
	log.Println(running)
}

func sol2() {
	mustBefore, updates := readInput()

	running := 0
	for _, update := range updates {
		valid := true
		for i, num := range update {
			mustBeBefore := mustBefore[num]
			for _, t := range mustBeBefore {
				for _, b := range update[:i] {
					if t == b {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}
		if !valid {
			slices.SortStableFunc(update, func(a, b int) int {
				mustBeBefore := mustBefore[a]
				for _, t := range mustBeBefore {
					if t == b {
						return -1
					}
				}
				return 1
			})
			running += update[len(update)/2]
		}
	}
	log.Println(running)
}
