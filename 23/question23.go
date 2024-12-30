package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type empty struct{}

type Graph map[string][]string

func load(filename string) (Graph, map[string]empty) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tpcs := make(map[string]empty)
	graph := make(map[string][]string)
	for scanner.Scan() {
		pcs := strings.Split(scanner.Text(), "-")

		if _, ok := graph[pcs[0]]; !ok {
			graph[pcs[0]] = []string{}
		}
		if _, ok := graph[pcs[1]]; !ok {
			graph[pcs[1]] = []string{}
		}
		graph[pcs[0]] = append(graph[pcs[0]], pcs[1])
		graph[pcs[1]] = append(graph[pcs[1]], pcs[0])

		if strings.HasPrefix(pcs[0], "t") {
			tpcs[pcs[0]] = empty{}
		}
		if strings.HasPrefix(pcs[1], "t") {
			tpcs[pcs[1]] = empty{}
		}
	}
	return graph, tpcs
}

func Part1(filename string) int {
	graph, tpcs := load(filename)
	uniquesets := make(map[string]empty)
	for tpc, _ := range tpcs {
		for _, pc2 := range graph[tpc] {
			for _, pc3 := range graph[pc2] {
				for _, pc4 := range graph[pc3] {
					if pc4 == tpc {
						key := []string{tpc, pc2, pc3}
						sort.Strings(key)
						uniquesets[strings.Join(key, "-")] = empty{}
					}
				}
			}
		}
	}
	return len(uniquesets)
}

func Part2(filename string) string {
	graph, _ := load(filename)
	pcList := make([]string, 0, len(graph))
	for pc := range graph {
		pcList = append(pcList, pc)
	}
	var maxSubGraph []string
	BK_algo(graph, []string{}, pcList, []string{}, &maxSubGraph)

	sort.Strings(maxSubGraph)
	return strings.Join(maxSubGraph, ",")
}

// had to look this up on wikipedia :/
func BK_algo(graph Graph, Res, Potnl, Excl []string, maxSubGraph *[]string) {
	if len(Potnl) == 0 && len(Excl) == 0 {
		if len(Res) > len(*maxSubGraph) {
			*maxSubGraph = append([]string{}, Res...)
		}
		return
	}

	pCopy := append([]string{}, Potnl...)
	for _, v := range pCopy {
		newR := append(Res, v)
		newP := intersect(Potnl, graph[v])
		newX := intersect(Excl, graph[v])

		BK_algo(graph, newR, newP, newX, maxSubGraph)

		Potnl = remove(Potnl, v)
		Excl = append(Excl, v)
	}
}

func intersect(a, b []string) []string {
	set := make(map[string]bool)
	for _, val := range b {
		set[val] = true
	}

	var result []string
	for _, val := range a {
		if set[val] {
			result = append(result, val)
		}
	}
	return result
}

func remove(slice []string, elem string) []string {
	for i, v := range slice {
		if v == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 1302)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %s (expected `cb,df,fo,ho,kk,nw,ox,pq,rt,sf,tq,wi,xz`)\n", ans2)
}
