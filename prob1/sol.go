package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() (a []int, b []int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vals := scanner.Text()
		splits := strings.Split(vals, "   ")
		val1, err := strconv.Atoi(splits[0])
		if err != nil {
			log.Fatal(err)
		}
		val2, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		a = append(a, val1)
		b = append(b, val2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	arr1, arr2 := readInput()
	slices.SortFunc(arr1, func(i int, j int) int {
		return i - j
	})
	slices.SortFunc(arr2, func(i int, j int) int {
		return i - j
	})

	running := 0
	for i, val := range arr1 {
		diff := arr2[i] - val
		if diff < 0 {
			diff *= -1
		}
		running += diff
	}
	log.Println(running)

	simularity := 0
	mult := 1
	for i, val := range arr1 {
		if i < len(arr1)-1 {
			if arr1[i+1] == val {
				mult++
				continue
			}
		}
		for j, val2 := range arr2 {
			if val2 == val {
				simularity += val * mult
			} else if val2 > val {
				arr2 = arr2[j:]
				break
			}
		}
		if i < len(arr1)-1 {
			mult = 1
		}
	}

	log.Println(simularity)
}
