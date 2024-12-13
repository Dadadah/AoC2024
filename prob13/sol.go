package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	button1 button
	button2 button

	prizex int
	prizey int
}

type button struct {
	x    int
	y    int
	cost int
}

func readInput() (problems []*problem) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := &problem{}
	index := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if len(val) == 0 {
			continue
		}

		splots := strings.Split(val, ": ")
		sploots := strings.Split(splots[1], ", ")
		x, err := strconv.Atoi(sploots[0][2:])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(sploots[1][2:])
		if err != nil {
			log.Fatal(err)
		}
		switch index % 3 {
		case 0:
			p.button1.x = x
			p.button1.y = y
			p.button1.cost = 3
		case 1:
			p.button2.x = x
			p.button2.y = y
			p.button2.cost = 1
		case 2:
			p.prizex = x + 10000000000000
			p.prizey = y + 10000000000000
			problems = append(problems, p)
			p = &problem{}
		}
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	// sol1()
	sol2()
}

func sol1() {
	problems := readInput()

	running := 0
	for _, p := range problems {
		cheapest := -1
		for i := 100; i > 0; i-- {
			for j := 100; j > 0; j-- {
				if p.button1.x*i+p.button2.x*j == p.prizex && p.button1.y*i+p.button2.y*j == p.prizey {
					val := p.button1.cost*i + p.button2.cost*j
					if cheapest == -1 || val < cheapest {
						cheapest = val
					}
				}
			}
		}
		if cheapest != -1 {
			running += cheapest
		}
	}
	log.Println(running)
}

func sol2() {
	problems := readInput()

	running := 0
	for _, p := range problems {
		match := true
		xMax := p.prizex / p.button2.x
		yMax := p.prizey / p.button2.y
		j := xMax
		if yMax < xMax {
			j = yMax
		}
		i := 0
		for xval, yval := p.button2.x*j+p.button1.x*i, p.button2.y*j+p.button1.y*i; xval != p.prizex || yval != p.prizey; xval, yval = p.button2.x*j+p.button1.x*i, p.button2.y*j+p.button1.y*i {
			xMaxDiff := (p.prizex - xval) / p.button1.x
			yMaxDiff := (p.prizey - yval) / p.button1.y

			if xMaxDiff > 1000 || yMaxDiff > 1000 {
				if xMaxDiff > yMaxDiff {
					i += xMaxDiff
					j -= (xMaxDiff*p.button1.y + p.button2.y) / p.button2.y
				}

				if yMaxDiff > xMaxDiff {
					i += yMaxDiff
					j -= (yMaxDiff*p.button1.x + p.button2.x) / p.button2.x
				}
			} else {
				j--
				for xval, yval = p.button2.x*j+p.button1.x*i, p.button2.y*j+p.button1.y*i; xval < p.prizex && yval < p.prizey; xval, yval = p.button2.x*j+p.button1.x*i, p.button2.y*j+p.button1.y*i {
					i++
				}
			}

			if i <= -1 || j <= -1 {
				match = false
				break
			}
			if i%100 == 0 {
				// log.Println(j, i, p.button2.x*j+p.button1.x*i, p.button2.y*j+p.button1.y*i, p.prizex, p.prizey)
			}
		}
		if match {
			running += i*3 + j
		}
	}
	log.Println(running)
}
