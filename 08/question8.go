package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	X int
	Y int
}

type empty struct{}

func (c1 *Coordinate) Subtract(c2 Coordinate) (int, int) {
	return c1.X - c2.X, c1.Y - c2.Y
}

func (c1 *Coordinate) AntiNode(c2 Coordinate, multiplier int) Coordinate {
	x, y := c1.Subtract(c2)
	return Coordinate{c1.X + (x * multiplier), c1.Y + (y * multiplier)}
}

func (c1 *Coordinate) WithinBounds(max_x int, max_y int) bool {
	return c1.X >= 0 && c1.Y >= 0 && c1.X < max_x && c1.Y < max_y
}

func load(filename string) (map[rune][]Coordinate, int, int) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	array := make(map[rune][]Coordinate)
	y := 0
	max_x := 0
	for scanner.Scan() {
		max_x = len(scanner.Text())
		for x, c := range []rune(scanner.Text()) {
			if c == '.' {
				continue
			}
			if _, ok := array[c]; !ok {
				array[c] = []Coordinate{{x, y}}
			} else {
				array[c] = append(array[c], Coordinate{x, y})
			}
		}
		y += 1
	}
	return array, max_x, y
}

func Part1(filename string) int {
	antinodes := make(map[Coordinate]empty)
	array, max_x, max_y := load(filename)
	for _, v := range array {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				a1 := v[i].AntiNode(v[j], 1)
				a2 := v[j].AntiNode(v[i], 1)
				if a1.WithinBounds(max_x, max_y) {
					antinodes[a1] = empty{}
				}
				if a2.WithinBounds(max_x, max_y) {
					antinodes[a2] = empty{}
				}
			}
		}
	}
	return len(antinodes)
}

func CreateAndTestAntinodes(
	visited *map[Coordinate]empty,
	antennae1 Coordinate,
	antennae2 Coordinate,
	max_x int,
	max_y int) {
	m := 1
	//cheap fix
	(*visited)[antennae2] = empty{}
	for {
		antinode := antennae1.AntiNode(antennae2, m)
		if antinode.WithinBounds(max_x, max_y) {
			(*visited)[antinode] = empty{}
			m += 1
		} else {
			break
		}
	}
}

func Part2(filename string) int {
	antinodes := make(map[Coordinate]empty)
	array, max_x, max_y := load(filename)
	for _, v := range array {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				CreateAndTestAntinodes(&antinodes, v[i], v[j], max_x, max_y)
				CreateAndTestAntinodes(&antinodes, v[j], v[i], max_x, max_y)
			}
		}
	}
	return len(antinodes)
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 295)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 1034)\n", ans2)
}
