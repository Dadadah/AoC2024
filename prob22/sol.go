package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readInput() (nums []int) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		num, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func next(a int) int {
	newVal := a << 6
	a = prune(mix(a, newVal))
	newVal = a >> 5
	a = prune(mix(a, newVal))
	newVal = a << 11
	return prune(mix(a, newVal))
}

func ones(a int) int {
	return a % 10
}

func main() {
	sol1()
	sol2()
}

func sol1() {
	nums := readInput()

	running := 0
	for _, num := range nums {
		cur := num
		for i := 0; i < 2000; i++ {
			cur = next(cur)
		}
		running += cur
	}

	log.Println(running)
}

type sequence struct {
	a int
	b int
	c int
	d int
}

func (s sequence) getHash() int {
	return s.a + s.b + s.c + s.d
}

func sol2() {
	nums := readInput()

	ons := make([][]int, len(nums))
	seqs := make([][]int, len(nums))
	adds := make([][]int, len(nums))

	for j, num := range nums {
		cur := num
		last := cur
		for i := 0; i < 2000; i++ {
			cur = next(cur)
			seqs[j] = append(seqs[j], ones(cur)-ones(last))
			ons[j] = append(ons[j], ones(cur))
			last = cur
		}
	}

	for j, seq := range seqs {
		for i := 0; i < len(seq)-3; i++ {
			toAdd := 0
			toAdd += seq[i]
			toAdd += seq[i+1]
			toAdd += seq[i+2]
			toAdd += seq[i+3]
			adds[j] = append(adds[j], toAdd)
		}
	}

	seen := make(map[sequence]bool)

	maxBananers := 0
	for seqID, seq := range seqs {
		log.Println(seqID)
		for i := 0; i < len(seq)-3; i++ {
			bananers := 0

			s := sequence{seq[i], seq[i+1], seq[i+2], seq[i+3]}

			if seen[s] {
				continue
			} else {
				seen[s] = true
			}

			target := s.getHash()
			for j, addSeq := range adds {
				for i2, adder := range addSeq {
					if adder == target {
						if seqs[j][i2] == seq[i] &&
							seqs[j][i2+1] == seq[i+1] &&
							seqs[j][i2+2] == seq[i+2] &&
							seqs[j][i2+3] == seq[i+3] {
							bananers += ons[j][i2+3]
							break
						}
					}

				}
			}
			if bananers > maxBananers {
				log.Println(seq[i], seq[i+1], seq[i+2], seq[i+3])
				maxBananers = bananers
			}
		}
	}

	log.Println(maxBananers)
}
