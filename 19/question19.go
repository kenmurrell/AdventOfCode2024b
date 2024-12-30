package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func load(filename string) ([]string, *map[string]bool) {
	file, _ := os.Open(filename)
	defer file.Close()
	towels := make(map[string]bool)
	var patterns []string
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			for _, t := range strings.Split(line, ", ") {
				towels[t] = true
			}
		} else if len(line) > 0 {
			patterns = append(patterns, line)
		}
		i++
	}
	return patterns, &towels
}

func Part1(filename string) int {
	patterns, towels := load(filename)
	count := 0
	for _, pattern := range patterns {
		if hasCombo(pattern, towels) {
			count += 1
		}
	}
	return count
}

func hasCombo(pattern string, towels *map[string]bool) bool {
	if len(pattern) == 0 {
		return true
	}
	for i := 1; i <= len(pattern); i++ {
		sub := pattern[:i]
		if (*towels)[sub] {
			if hasCombo(pattern[i:], towels) {
				return true
			}
		}
	}
	return false
}

func Part2(filename string) int {
	patterns, towels := load(filename)
	count := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		count += countCombos(pattern, towels, &cache)
	}

	return count
}

func countCombos(pattern string, towels *map[string]bool, cache *map[string]int) int {
	if len(pattern) == 0 {
		return 1
	}
	if c, ok := (*cache)[pattern]; ok {
		return c
	}
	count := 0
	for i := 1; i <= len(pattern); i++ {
		sub := pattern[:i]
		if _, ok := (*towels)[sub]; ok {
			count += countCombos(pattern[i:], towels, cache)
		}
	}
	(*cache)[pattern] = count
	return count
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 285)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 636483903099279)\n", ans2)
}
