package main

import (
	"bufio"
	"log"
	"os"
)

type lock []int

type key []int

func buildKey(grid [][]rune) key {
	k := key{}
	for x := range grid[0] {
		count := 0
		for y := range grid {
			if grid[y][x] == '#' {
				count++
			}
		}
		k = append(k, count-1)
	}
	return k
}

func buildLock(grid [][]rune) lock {
	k := lock{}
	for x := range grid[0] {
		count := 0
		for y := range grid {
			if grid[y][x] == '#' {
				count++
			}
		}
		k = append(k, count-1)
	}
	return k
}

func readInput() (locks []lock, keys []key) {
	fileName := "input"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lockKeyGrid [][]rune

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if len(val) == 0 {
			if lockKeyGrid[0][0] == '.' {
				keys = append(keys, buildKey(lockKeyGrid))
			} else {
				locks = append(locks, buildLock(lockKeyGrid))
			}

			lockKeyGrid = nil
			continue
		}

		lockKeyGrid = append(lockKeyGrid, []rune(val))
	}

	if lockKeyGrid[0][0] == '.' {
		keys = append(keys, buildKey(lockKeyGrid))
	} else {
		locks = append(locks, buildLock(lockKeyGrid))
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
	locks, keys := readInput()

	running := 0
	for _, k := range keys {
		for _, l := range locks {
			match := true
			for i, v := range k {
				if l[i]+v > 5 {
					match = false
					break
				}
			}
			if match {
				running++
			}
		}
	}
	log.Println(running)
}

func sol2() {

}
