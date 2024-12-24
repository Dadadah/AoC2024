package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type gate struct {
	input1     string
	input1Gate *gate
	input2     string
	input2Gate *gate
	output     string
	operation  int
	solution   int
}

type wire struct {
	name string
	val  int
}

func readInput() (wires map[string]wire, gates []*gate) {
	fileName := "input"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wires = make(map[string]wire)
	scanningGates := false
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		if len(val) == 0 {
			scanningGates = true
			continue
		}

		if !scanningGates {
			splits := strings.Split(val, ": ")
			num, err := strconv.Atoi(splits[1])
			if err != nil {
				log.Fatal(err)
			}
			wires[splits[0]] = wire{splits[0], num}
		} else {
			splits := strings.Split(val, " ")
			g := &gate{}
			g.input1 = splits[0]
			operation := 0
			switch splits[1] {
			case "OR":
				operation = 1
			case "XOR":
				operation = 2
			}
			g.operation = operation
			g.input2 = splits[2]
			g.output = splits[4]
			g.solution = -1
			gates = append(gates, g)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func buildGateMap(gates []*gate) (ends []*gate) {
	gatesByInputs = make(map[string][]*gate)
	for _, g := range gates {
		if g.buildInputs(gates) {
			ends = append(ends, g)
		}
	}
	return
}

func resetGatesSolutions(gates []*gate) {
	for _, g := range gates {
		g.solution = -1
	}
}

var gatesByInputs map[string][]*gate

func (g *gate) buildInputs(gates []*gate) bool {
	g.input1Gate = getOutputGate(gates, g.input1)
	g.input2Gate = getOutputGate(gates, g.input2)
	gatesByInputs[g.input1] = append(gatesByInputs[g.input1], g)
	gatesByInputs[g.input2] = append(gatesByInputs[g.input2], g)
	return strings.HasPrefix(g.output, "z")
}

func (g *gate) printStack() {
	if g.input1Gate != nil {
		g.input1Gate.printStack()
	}
	if g.input2Gate != nil {
		g.input2Gate.printStack()
	}
	g.print()
}

func (g *gate) print() {

	op := "&"
	if g.operation == 1 {
		op = "|"
	} else if g.operation == 2 {
		op = "^"
	}
	log.Println(g.input1, op, g.input2, g.output)
}

func (g *gate) count() int {
	count := 0
	if g.input1Gate != nil {
		count += g.input1Gate.count()
	}
	if g.input2Gate != nil {
		count += g.input2Gate.count()
	}
	return count + 1
}

func (g *gate) solve(wires map[string]wire) int {
	if g.solution >= 0 {
		return g.solution
	}
	val1 := -1
	val2 := -1
	if g.input1Gate != nil {
		val1 = g.input1Gate.solve(wires)
	} else {
		val1 = wires[g.input1].val
	}
	if g.input2Gate != nil {
		val2 = g.input2Gate.solve(wires)
	} else {
		val2 = wires[g.input2].val
	}
	output := 0
	switch g.operation {
	case 0:
		output = val1 & val2
	case 1:
		output = val1 | val2
	case 2:
		output = val1 ^ val2
	}
	g.solution = output
	return output
}

func getOutputGate(gates []*gate, w string) *gate {
	for _, g := range gates {
		if g.output == w {
			return g
		}
	}
	return nil
}

func getOutput(ends []*gate, wires map[string]wire) int {
	running := 0
	for _, end := range ends {
		sol := end.solve(wires)
		running += sol
		running = running << 1
	}
	running = running >> 1
	return running
}

func main() {
	sol1()
	sol2()
}

func sol1() {
	wires, gates := readInput()

	ends := buildGateMap(gates)

	slices.SortFunc(ends, func(a, b *gate) int {
		return -strings.Compare(a.output, b.output)
	})

	log.Println(getOutput(ends, wires))
}

func checkAdder(bit int, carryGate *gate) {
	input1Gates := gatesByInputs[fmt.Sprintf("x%02d", bit)]
	// input2Gates := gatesByInputs["y" + bitAsString]

	if len(input1Gates) == 0 {
		log.Println("Found all bad gates")
		return
	}

	// find xor and and gates
	var xorGate, andGate *gate
	for _, g := range input1Gates {
		if g.operation == 2 {
			xorGate = g
		}
		if g.operation == 0 {
			andGate = g
		}
	}

	// no carry == first adder, no other checks
	if carryGate == nil {
		checkAdder(bit+1, andGate)
		return
	}

	// check the xorGate and carry gates are both outputting to one xor and one and gate
	xorOutGates := gatesByInputs[xorGate.output]
	carryOutGates := gatesByInputs[carryGate.output]
	var xorOutGate, andOutGate *gate
	if len(xorOutGates) != 2 {
		xorGate.print()
	} else {
		for _, g := range xorOutGates {
			if g.operation == 2 {
				xorOutGate = g
			}
			if g.operation == 0 {
				andOutGate = g
			}
		}
	}
	if len(carryOutGates) != 2 {
		carryGate.print()
	} else {
		for _, g := range carryOutGates {
			if g.operation == 2 {
				if xorOutGate == nil {
					xorOutGate = g
				}
				if xorOutGate.output != g.output {
					if xorOutGate.output != fmt.Sprintf("z%02d", bit) {
						xorGate.print()
					} else {
						carryGate.print()
					}
				} else {
					if xorOutGate.output != fmt.Sprintf("z%02d", bit) {
						xorOutGate.print()
					}
				}
			}
			if g.operation == 0 {
				andOutGate = g
			}
		}
	}

	if len(xorOutGates) != 2 && len(carryOutGates) != 2 {
		log.Println(xorOutGates)
		log.Println(carryOutGates)
		panic("AHHHHHHH")
	}

	andOutGates := gatesByInputs[andGate.output]
	andOutOutGates := gatesByInputs[andOutGate.output]
	var orGate *gate
	if len(andOutGates) != 1 || andOutGates[0].operation != 1 {
		andGate.print()
	} else {
		orGate = andOutGates[0]
	}
	if len(andOutOutGates) != 1 || andOutOutGates[0].operation != 1 {
		andOutGate.print()
	} else {
		if orGate == nil {
			orGate = andOutOutGates[0]
		}
		if orGate.output != andOutOutGates[0].output {
			panic("EEEEEEEEE")
		}
	}
	if orGate != nil {
		checkAdder(bit+1, orGate)
	}
}

func sol2() {
	_, gates := readInput()

	ends := buildGateMap(gates)

	slices.SortFunc(ends, func(a, b *gate) int {
		return -strings.Compare(a.output, b.output)
	})

	checkAdder(0, nil)

}
