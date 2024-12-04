package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var code []string = []string{"X", "M", "A", "S"}

func Part1(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}

	sum := 0
	for r := range len(data) {
		for c := range len(data[0]) {
			if data[r][c] == "X" {
				sum += Detect(data, r, c)
			}
		}
	}
	return sum
}

func Detect(data [][]string, r int, c int) int {
	sum := 0
	// Check Up
	sum += CheckDirection(data, r, c, -1, 0)
	// Check Up Right
	sum += CheckDirection(data, r, c, -1, 1)
	// Check Right
	sum += CheckDirection(data, r, c, 0, 1)
	// Check Down Right
	sum += CheckDirection(data, r, c, 1, 1)
	// Check Down
	sum += CheckDirection(data, r, c, 1, 0)
	// Check Down Left
	sum += CheckDirection(data, r, c, 1, -1)
	// Check Left
	sum += CheckDirection(data, r, c, 0, -1)
	// Check Up Left
	sum += CheckDirection(data, r, c, -1, -1)
	return sum
}

// r and c are either, -1, 0, or 1 depending on the direction here
func CheckDirection(data [][]string, r int, c int, rdir int, cdir int) int {
	for i := 1; i < len(code); i++ {
		i_r := r + (rdir * i)
		i_c := c + (cdir * i)
		if i_r >= len(data) || i_r < 0 || i_c >= len(data[0]) || i_c < 0 {
			return 0
		}
		if data[i_r][i_c] != code[i] {
			return 0
		}
	}
	return 1
}

func DetectX(data [][]string, r int, c int) int {
	get := func(i_r int, i_c int) string {
		if i_r < 0 || i_r >= len(data) || i_c < 0 || i_c >= len(data[0]) {
			return ""
		}
		return data[i_r][i_c]
	}
	tl := get(r-1, c-1)
	tr := get(r-1, c+1)
	bl := get(r+1, c-1)
	br := get(r+1, c+1)
	if tl == "" || tr == "" || bl == "" || br == "" {
		return 0
	}

	// M   M      S   S
	//   A    or    A
	// S   S      M   M
	// scenario 1
	if tl == tr && bl == br && ((tl == "S" && bl == "M") || (tl == "M" && bl == "S")) {
		return 1
	}

	// S   M      M   S
	//   A    or    A
	// S   M      M   S
	// scenario 2
	if tl == bl && tr == br && ((tl == "S" && tr == "M") || (tl == "M" && tr == "S")) {
		return 1
	}
	return 0
}

func Part2(filename string) int {
	file, _ := os.Open(filename)
	defer file.Close()
	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}

	sum := 0
	for r := range len(data) {
		for c := range len(data[0]) {
			if data[r][c] == "A" {
				sum += DetectX(data, r, c)
			}
		}
	}
	return sum
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 2434)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 1835)\n", ans2)
}
