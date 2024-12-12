package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func load(filename string) []int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stones []int
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), " ") {
			val, _ := strconv.Atoi(c)
			stones = append(stones, val)
		}
	}
	return stones
}

func Part1(filename string) int {
	stones := load(filename)
	blinks := 25
	for b := range blinks {
		for i:=0; i<len(stones); i++ {
			if stones[i] == 0 {
				stones[i] += 1
				continue
			} 
			var num int = stones[i]
			digits := 0
			for num > 0{
				num /= 10
				digits+=1
			}
			if digits % 2 == 0 {
				first := int(stones[i] / int(math.Pow10(digits/2)))
				second := stones[i] - (first*int(math.Pow10(digits/2)))
				stones = append(stones[:i+1], stones[i:]...)
				stones[i] = first
				stones[i+1] = second
				i += 1
				continue
			}
			stones[i] *= 2024
		}
		fmt.Printf("%d\n", b)
	}

	return len(stones)
}

func Part2(filename string) int {
	return 0
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 186996)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected )\n", ans2)
}
