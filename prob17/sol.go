package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	A int
	B int
	C int

	tape []int
}

func readInput() (mach machine) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	line := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if line == 0 {
			splits := strings.Split(val, " ")
			mach.A, _ = strconv.Atoi(splits[len(splits)-1])
		} else if line == 1 {
			splits := strings.Split(val, " ")
			mach.B, _ = strconv.Atoi(splits[len(splits)-1])
		} else if line == 2 {
			splits := strings.Split(val, " ")
			mach.C, _ = strconv.Atoi(splits[len(splits)-1])
		} else if line == 4 {
			splits := strings.Split(val, " ")
			splots := strings.Split(splits[len(splits)-1], ",")
			for _, s := range splots {
				val, _ := strconv.Atoi(s)
				mach.tape = append(mach.tape, val)
			}
		}

		line++
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

func (m machine) combo(val int) int {
	switch val {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return val
	case 4:
		return m.A
	case 5:
		return m.B
	case 6:
		return m.C
	case 7:
		fallthrough
	default:
		panic("Ah!")
	}
}

func sol1() {
	mach := readInput()

	var out []int

	for i := 0; i < len(mach.tape); i += 2 {
		operator := mach.tape[i]
		operand := mach.tape[i+1]

		switch operator {
		case 0:
			mach.A = mach.A / (1 << mach.combo(operand))
		case 1:
			mach.B = mach.B ^ operand
		case 2:
			mach.B = mach.combo(operand) & 7
		case 3:
			if mach.A != 0 {
				i = operand - 2
			}
		case 4:
			mach.B = mach.B ^ mach.C
		case 5:
			out = append(out, mach.combo(operand)&7)
		case 6:
			mach.B = mach.A / (1 << mach.combo(operand))
		case 7:
			mach.C = mach.A / (1 << mach.combo(operand))
		}
	}

	strVal := ""
	for i, v := range out {
		if i == 0 {
			strVal += strconv.Itoa(v)
		} else {
			strVal += "," + strconv.Itoa(v)
		}
	}

	log.Println(strVal)
}

func solve(mach machine, a, depth int) bool {
	mach.A = a
	var out []int
	for i := 0; i < len(mach.tape); i += 2 {
		operator := mach.tape[i]
		operand := mach.tape[i+1]

		switch operator {
		case 0:
			mach.A = mach.A / (1 << mach.combo(operand))
		case 1:
			mach.B = mach.B ^ operand
		case 2:
			mach.B = mach.combo(operand) & 7
		case 3:
			if mach.A != 0 {
				i = operand - 2
			}
		case 4:
			mach.B = mach.B ^ mach.C
		case 5:
			val := mach.combo(operand) & 7
			out = append(out, val)
			if len(out) == depth+1 {
				// check all outputs for correctness
				for i, o := range out {
					if o != mach.tape[len(mach.tape)-len(out)+i] {
						return false
					}
				}

				if len(out) == len(mach.tape) {
					log.Println(a)
					return true
				}

				newA := a * 8
				if !solve(mach, newA, depth+1) && !solve(mach, newA+1, depth+1) && !solve(mach, newA+2, depth+1) && !solve(mach, newA+3, depth+1) && !solve(mach, newA+4, depth+1) && !solve(mach, newA+5, depth+1) && !solve(mach, newA+6, depth+1) && !solve(mach, newA+7, depth+1) {
					return false
				}
			}

		case 6:
			mach.B = mach.A / (1 << mach.combo(operand))
		case 7:
			mach.C = mach.A / (1 << mach.combo(operand))
		}
	}
	return false
}

func sol2() {
	mach := readInput()

	// solve(mach, 0, 0)
	solve(mach, 1, 0)
	solve(mach, 2, 0)
	solve(mach, 3, 0)
	solve(mach, 4, 0)
	solve(mach, 5, 0)
	solve(mach, 6, 0)
	solve(mach, 7, 0)
}
