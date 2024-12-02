package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Load(filename string) [][]int {
	file, _ := os.Open(filename)
	defer file.Close()
	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str_row := strings.Fields(scanner.Text())
		int_row := make([]int, len(str_row))
		for i, _ := range str_row {
			c, _ := strconv.Atoi(str_row[i])
			int_row[i] = c
		}
		data = append(data, int_row)
	}
	return data
}

func CheckSafe(row []int) bool {
	var isIncrease bool
	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i-1]
		if i == 1 {
			isIncrease = diff > 0
		}
		if diff == 0 {
			return false
		}
		if (isIncrease && diff < 0) || (!isIncrease && diff > 0) {
			return false
		}
		if math.Abs(float64(diff)) > 3 {
			return false
		}
	}
	return true
}

func Part1(filename string) int {
	data := Load(filename)
	var sum int = 0
	for r, _ := range data {
		if CheckSafe(data[r]) {
			sum += 1
		}
	}

	return sum
}

func Part2(filename string) int {
	data := Load(filename)
	var sum int = 0
	for r, _ := range data {
		myrow := data[r]
		if CheckSafe(myrow) {
			sum += 1
		} else {
			var i int = 0
		recheck:
			testrow := make([]int, len(myrow))
			copy(testrow, myrow)
			testrow = append(testrow[:i], testrow[i+1:]...)
			if CheckSafe(testrow) {
				sum += 1
			} else if i < len(myrow)-1 {
				i += 1
				goto recheck
			}
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
