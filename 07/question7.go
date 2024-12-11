package main

import (
	"bufio"
	"fmt"
	"math"
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

func SimpleTestEquation(target int, nums []int) bool {
	combos := math.Pow(2, float64(len(nums)-1))
	for c := 0; c < int(combos); c++ {
		sum := nums[0]
		for i, n := range nums[1:] {
			if (1<<i)&c == 0 {
				sum += n
			} else {
				sum *= n
			}
		}
		if sum == target {
			return true
		}
	}
	return false
}

func Part1(filename string) int {
	eq := load(filename)
	total := 0
	for _, e := range eq {
		target := e[0]
		if SimpleTestEquation(target, e[1:]) {
			total += target
		}
	}
	return total
}

func RecursiveTestEquation(target int, total int, nums []int) bool {
	if len(nums) == 0 {
		return total == target
	}
	for i := 0; i < 3; i++ {
		newtotal := total
		switch i {
		case 0:
			newtotal += nums[0]
		case 1:
			newtotal *= nums[0]
		case 2:
			var digits int
			if nums[0] == 0 {
				digits = 1
			} else {
				digits = int(math.Log10(float64(nums[0])) + 1)
			}
			newtotal = (newtotal * int(math.Pow10(digits))) + nums[0]
		}
		if RecursiveTestEquation(target, newtotal, nums[1:]) {
			return true
		}
	}
	return false
}

func Part2(filename string) int64 {
	eq := load(filename)
	var total int64 = 0
	for _, e := range eq {
		target := e[0]
		if RecursiveTestEquation(target, e[1], e[2:]) {
			total += int64(target)
		}
	}
	return total
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 663613490587)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 110365987435001)\n", ans2)
}
