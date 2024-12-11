package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type empty struct{}

type Coordinate struct {
	Y int
	X int
}

func load(filename string) [][]int {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var tmap [][]int
	for scanner.Scan() {
		line := scanner.Text()
		iline := make([]int, len(line))
		for i, c := range strings.Split(line, "") {
			val, _ := strconv.Atoi(c)
			iline[i] = val
		}
		tmap = append(tmap, iline)
	}
	return tmap
}

func FollowPath(tmap *[][]int, visited *map[Coordinate]empty, y int, x int) int {
	val := (*tmap)[y][x]
	if visited != nil {
		c := Coordinate{y, x}
		if _, ok := (*visited)[c]; ok {
			return 0
		} else {
			(*visited)[c] = empty{}
		}
	}
	if val == 9 {
		return 1
	}
	local_count := 0
	if y+1 < len(*tmap) && (*tmap)[y+1][x] == val+1 {
		local_count += FollowPath(tmap, visited, y+1, x)
	}
	if y > 0 && (*tmap)[y-1][x] == val+1 {
		local_count += FollowPath(tmap, visited, y-1, x)
	}
	if x+1 < len((*tmap)[0]) && (*tmap)[y][x+1] == val+1 {
		local_count += FollowPath(tmap, visited, y, x+1)
	}
	if x > 0 && (*tmap)[y][x-1] == val+1 {
		local_count += FollowPath(tmap, visited, y, x-1)
	}
	return local_count
}

func Part1(filename string) int {
	tmap := load(filename)
	count := 0
	for y := 0; y < len(tmap); y++ {
		for x := 0; x < len(tmap[0]); x++ {
			if tmap[y][x] == 0 {
				visited := make(map[Coordinate]empty)
				count += FollowPath(&tmap, &visited, y, x)
			}
		}
	}
	return count
}

func Part2(filename string) int {
	tmap := load(filename)
	count := 0
	for y := 0; y < len(tmap); y++ {
		for x := 0; x < len(tmap[0]); x++ {
			if tmap[y][x] == 0 {
				count += FollowPath(&tmap, nil, y, x)
			}
		}
	}
	return count
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 512)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 1045)\n", ans2)
}
