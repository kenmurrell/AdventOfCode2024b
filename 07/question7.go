package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func load(filename string) [][]int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var equations [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var eq []int
		for _, c := range strings.Split(line, " ") {
			i, _ := strconv.Atoi(strings.Trim(c, ":"))
			eq = append(eq, i)
		}
		equations = append(equations, eq)
	}
	return equations
}

func Part1(filename string) int {
	eq := load(filename)
	for _, e := range eq {
		target := e[0]
		sum := e[1]
		for i, _ := range e[2:] {

		}
	}
}

func Part2(filename string) int {
	return 0
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 4890)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 1835)\n", ans2)
}
