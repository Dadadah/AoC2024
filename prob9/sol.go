package main

import (
	"bufio"
	"log"
	"os"
)

func readInput() (diskMap []rune) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		diskMap = []rune(val)
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

var runeToInt = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func sol1() {
	diskMap := readInput()

	var disk []int

	fileNum := 0
	file := true
	for _, d := range diskMap {
		val := runeToInt[d]
		if file {
			for i := 0; i < val; i++ {
				disk = append(disk, fileNum)
			}
			fileNum++
		} else {
			for i := 0; i < val; i++ {
				disk = append(disk, -1)
			}
		}
		file = !file
	}

	for end := len(disk) - 1; end >= 0; end-- {
		if disk[end] == -1 {
			continue
		}
		empty := -1
		for begin := 0; begin < end; begin++ {
			if disk[begin] == -1 {
				empty = begin
				break
			}
		}
		if empty == -1 {
			break
		}
		disk[empty] = disk[end]
		disk[end] = -1
	}

	running := 0
	for i, d := range disk {
		if d == -1 {
			continue
		}
		running += i * d
	}
	log.Println(running)
}

func sol2() {
	diskMap := readInput()

	var disk []int

	fileNum := 0
	file := true
	for _, d := range diskMap {
		val := runeToInt[d]
		if file {
			for i := 0; i < val; i++ {
				disk = append(disk, fileNum)
			}
			fileNum++
		} else {
			for i := 0; i < val; i++ {
				disk = append(disk, -1)
			}
		}
		file = !file
	}

	for end := len(disk) - 1; end >= 0; {
		// If the disk space is empty, continue
		if disk[end] == -1 {
			end--
			continue
		}
		fileSize := 1
		file := disk[end]
		for off := 1; end-off >= 0; off++ {
			if disk[end-off] == file {
				fileSize++
			} else {
				break
			}
		}
		fileWritten := false
		for begin := 0; begin <= end; {
			// If the disk space is taken, continue
			if disk[begin] != -1 {
				begin++
				continue
			}
			size := 1
			// count free space
			for off := 1; begin+off < end; off++ {
				if disk[begin+off] == -1 {
					size++
				} else {
					break
				}
			}

			if fileSize <= size {
				// We can move the file
				for i := 0; i < fileSize; i++ {
					disk[begin+i] = file
					disk[end-i] = -1
				}
				end -= fileSize
				fileWritten = true
				break
			} else {
				begin += size
			}
		}
		if !fileWritten {
			end -= fileSize
		}
	}

	running := 0
	for i, d := range disk {
		if d == -1 {
			continue
		}
		running += i * d
	}
	log.Println(running)
}
