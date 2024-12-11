package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type equationsHolder struct {
	result   int
	operands []int
}

type equation struct {
	operands  []int
	operators []int8
}

func (eq equation) copy() equation {
	var operands []int
	var operators []int8
	operands = append(operands, eq.operands...)
	operators = append(operators, eq.operators...)
	return equation{operands, operators}
}

func (eq equation) copyWith(operand int, operator int8) equation {
	newEq := eq.copy()
	newEq.operands = append(newEq.operands, operand)
	newEq.operators = append(newEq.operators, operator)
	return newEq
}

func (eq equation) value() int {
	running := 0
	for i, operator := range eq.operators {
		if i == 0 {
			running = eq.operands[0]
		}
		if operator == 0 {
			running += eq.operands[i+1]
		} else if operator == 1 {
			running *= eq.operands[i+1]
		} else if operator == 2 {
			asString := strconv.Itoa(eq.operands[i+1])
			running *= int(math.Pow10(len(asString)))
			running += eq.operands[i+1]
		}
	}
	return running
}

func (eqh equationsHolder) value() int {
	var eqs []equation
	eqs = append(eqs, equation{operands: []int{eqh.operands[0]}})
	for i := 1; i < len(eqh.operands); i++ {
		var newEqs []equation
		for _, eq := range eqs {
			newEqs = append(newEqs, eq.copyWith(eqh.operands[i], 0))
			newEqs = append(newEqs, eq.copyWith(eqh.operands[i], 1))
			newEqs = append(newEqs, eq.copyWith(eqh.operands[i], 2))
		}
		eqs = newEqs
	}

	for _, eq := range eqs {
		if eq.value() == eqh.result {
			return eqh.result
		}
	}
	return 0
}

func readInput() (eqs []equationsHolder) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()
		eq := equationsHolder{}
		splits := strings.Split(val, ": ")
		eq.result, err = strconv.Atoi(splits[0])
		if err != nil {
			log.Fatal(err)
		}

		opSplits := strings.Split(splits[1], " ")
		for _, opSplit := range opSplits {
			op, err := strconv.Atoi(opSplit)
			if err != nil {
				log.Fatal(err)
			}
			eq.operands = append(eq.operands, op)
		}
		eqs = append(eqs, eq)
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
	eqhs := readInput()
	running := 0
	for _, eqh := range eqhs {
		running += eqh.value()
	}
	log.Println(running)
}

func sol2() {

}
