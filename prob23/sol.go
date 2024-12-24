package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

type connection struct {
	pca string
	pcb string
}

func readInput() (conns []connection) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := scanner.Text()

		splits := strings.Split(val, "-")
		conn := connection{splits[0], splits[1]}
		conns = append(conns, conn)
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

type computer struct {
	name        string
	connections map[string]*computer
}

func buildGraph(conns []connection) map[string]*computer {
	computerMap := make(map[string]*computer)

	for _, conn := range conns {
		var comp1, comp2 *computer
		if comp1 = computerMap[conn.pca]; comp1 == nil {
			comp1 = &computer{name: conn.pca, connections: make(map[string]*computer)}
			computerMap[conn.pca] = comp1
		}
		if comp2 = computerMap[conn.pcb]; comp2 == nil {
			comp2 = &computer{name: conn.pcb, connections: make(map[string]*computer)}
			computerMap[conn.pcb] = comp2
		}
		comp1.connections[comp2.name] = comp2
		comp2.connections[comp1.name] = comp1
	}
	return computerMap
}

func sol1() {
	conns := readInput()

	computerMap := buildGraph(conns)

	graffmap := make(map[string]bool)

	running := 0
	for _, comp := range computerMap {
		if comp.name[0] == 't' {
			for _, nei1 := range comp.connections {
				for _, nei2 := range comp.connections {
					if nei1.name == nei2.name {
						continue
					}
					if nei1.connections[nei2.name] != nil {
						names := []string{comp.name, nei1.name, nei2.name}
						slices.SortFunc(names, func(a, b string) int {
							return strings.Compare(a, b)
						})
						index := strings.Join(names, "")
						if graffmap[index] {
							continue
						} else {
							graffmap[index] = true
						}
						running++
					}
				}
			}
		}
	}
	log.Println(running)
}

type computerNode struct {
	name     string
	children []*computerNode
}

func (cn *computerNode) tryFit(comp *computer) bool {
	if comp.connections[cn.name] == nil {
		return false
	} else {
		for _, child := range cn.children {
			child.tryFit(comp)
		}
		cn.children = append(cn.children, &computerNode{name: comp.name})
		return true
	}
}

func (cn *computerNode) largestGroup() []string {
	var longest []string
	for _, child := range cn.children {
		s := child.largestGroup()
		if len(s) > len(longest) {
			longest = s
		}
	}
	longest = append(longest, cn.name)
	return longest
}

func sol2() {
	conns := readInput()

	computerMap := buildGraph(conns)

	var trees []*computerNode
	count := 0
	for _, comp := range computerMap {
		count++
		found := false
		for _, tree := range trees {
			if tree.tryFit(comp) {
				found = true
				break
			}
		}
		if !found {
			trees = append(trees, &computerNode{name: comp.name})
		}
	}

	var longest []string
	for _, tree := range trees {
		s := tree.largestGroup()
		if len(s) > len(longest) {
			longest = s
		}
	}
	log.Println(len(longest))
	slices.SortFunc(longest, func(a, b string) int {
		return strings.Compare(a, b)
	})
	log.Println(strings.Join(longest, ","))
}
