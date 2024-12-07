package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type empty struct {
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) IsOutside(max_x int, max_y int) bool {
	return c.X >= max_x || c.X < 0 || c.Y >= max_y || c.Y < 0
}

type Direction int

const (
	Up = iota
	Right
	Down
	Left
)

func Next(dir Direction, c *Coordinate) Coordinate {
	switch dir {
	case Up:
		return Coordinate{c.X, c.Y - 1}
	case Down:
		return Coordinate{c.X, c.Y + 1}
	case Left:
		return Coordinate{c.X - 1, c.Y}
	case Right:
		return Coordinate{c.X + 1, c.Y}
	}
	panic("oh no")
}

func travel(max_x int, max_y int, startPos Coordinate, startDir Direction, walls map[Coordinate]empty) map[Coordinate]empty {
	visited := make(map[Coordinate]empty)
	visited[startPos] = empty{}
	for {
		next := Next(startDir, &startPos)
		if _, ok := walls[next]; ok {
			startDir = (startDir + 1) % 4
			continue
		} else if next.IsOutside(max_x, max_y) {
			return visited
		}
		startPos = next
		visited[startPos] = empty{}
	}
}

func load(filename string) (int, int, Coordinate, Direction, map[Coordinate]empty) {
	file, _ := os.Open(filename)
	defer file.Close()
	walls := make(map[Coordinate]empty)
	scanner := bufio.NewScanner(file)
	max_x := 0
	max_y := 0
	var startDir Direction
	var startPos Coordinate
	for scanner.Scan() {
		line := scanner.Text()
		max_x = len(line)
		for x, c := range strings.Split(line, "") {
			switch c {
			case ".":
				continue
			case "#":
				walls[Coordinate{x, max_y}] = empty{}
			default:
				startPos = Coordinate{x, max_y}
				switch c {
				case "^":
					startDir = Up
				case ">":
					startDir = Right
				case "v":
					startDir = Down
				case "<":
					startDir = Left
				default:
					panic("oh nooo")
				}
			}
		}
		max_y += 1
	}
	fmt.Printf("Bounds: X(0,%d) Y(0,%d)\nStart Direction: %d\nStart Position: (%d, %d)\nWalls: %d\n", max_x, max_y, startDir, startPos.X, startPos.Y, len(walls))
	return max_x, max_y, startPos, startDir, walls
}

func Part1(filename string) int {
	max_x, max_y, startPos, startDir, walls := load(filename)
	visted := travel(max_x, max_y, startPos, startDir, walls)
	return len(visted)
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
