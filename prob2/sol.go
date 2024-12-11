package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() (arrs [][]int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vals := scanner.Text()
		splits := strings.Split(vals, " ")
		var xarr []int
		for _, split := range splits {
			val, err := strconv.Atoi(split)
			if err != nil {
				log.Fatal(err)
			}
			xarr = append(xarr, val)
		}
		arrs = append(arrs, xarr)
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
	arrs := readInput()

	safeCount := 0
	for _, arr := range arrs {
		if isSafe(arr) {
			safeCount++
		}
	}
	log.Println(safeCount)
}

func sol2() {
	arrs := readInput()

	safeCount := 0
	for _, arr := range arrs {
		if !isSafe(arr) {
			for i := range arr {
				var n []int
				n = append(n, arr[:i]...)
				n = append(n, arr[i+1:]...)
				if isSafe(n) {
					safeCount++
					break
				}
			}
		} else {
			safeCount++
		}
	}
	log.Println(safeCount)
}

func isSafe(arr []int) bool {
	lastval := 0
	var derArr []int
	for i, val := range arr {
		if i == 0 {
			lastval = val
			continue
		}
		derArr = append(derArr, val-lastval)
		lastval = val
	}
	sign := 0
	for _, val := range derArr {
		if val > 3 || val < -3 || val == 0 {
			return false
		}
		if sign == 0 {
			if val > 0 {
				sign = 1
			} else {
				sign = -1
			}
		} else {
			if (val > 0 && sign == -1) || (val < 0 && sign == 1) {
				return false
			}
		}
	}
	return true
}
