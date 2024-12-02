package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func Part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	var left []int
	var right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(nums[0])
		left = append(left, num1)
		num2, _ := strconv.Atoi(nums[1])
		right = append(right, num2)
	}
	var wg sync.WaitGroup

	job := func(l []int) {
		defer wg.Done()
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
	}

	wg.Add(1)
	go job(left)
	wg.Add(1)
	go job(right)
	wg.Wait()
	sum := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		sum += int(math.Abs(float64(diff)))
	}
	return sum
}

func Part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	var left []int
	right := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(nums[0])
		left = append(left, num1)
		num2, _ := strconv.Atoi(nums[1])
		if _, ok := right[num2]; ok {
			right[num2] += 1
		} else {
			right[num2] = 1
		}
	}
	sum := 0
	for i := 0; i < len(left); i++ {
		if val, ok := right[left[i]]; ok {
			// fmt.Println(val)
			sum += left[i] * val
		}
	}
	return sum
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d\n", ans2)
}
