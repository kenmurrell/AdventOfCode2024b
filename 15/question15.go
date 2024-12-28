package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

func (p1 *Position) Add(p2 Position) *Position {
	return &Position{p1.X + p2.X, p1.Y + p2.Y}
}

type Ob int

const (
	Wall = 1
	Box  = 2
)

func load(filename string) (Position, map[Position]Ob, []Position) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	floor := make(map[Position]Ob)
	var operations []Position
	var robot Position
	y := -1
	for scanner.Scan() {
		y += 1
		line := scanner.Text()
		for x, c := range strings.Split(line, "") {
			p := Position{x, y}
			switch c {
			case "#":
				floor[p] = Wall
			case "@":
				robot = p
			case "O":
				floor[p] = Box
			case "^":
				operations = append(operations, Position{0, -1})
			case "v":
				operations = append(operations, Position{0, 1})
			case ">":
				operations = append(operations, Position{1, 0})
			case "<":
				operations = append(operations, Position{-1, 0})
			}
		}
	}
	return robot, floor, operations
}

func MoveBox(box Position, move Position, floor *map[Position]Ob) bool {
	newPos := box.Add(move)
	if ob, ok := (*floor)[*newPos]; !ok {
		delete(*floor, box)
		(*floor)[*newPos] = Box
		return true
	} else if ob == Wall {
		return false
	} else if ob == Box && MoveBox(*newPos, move, floor) {
		delete(*floor, box)
		(*floor)[*newPos] = Box
		return true
	}
	return false
}

func Part1(filename string) int {
	robot, floor, operations := load(filename)
	for _, op := range operations {
		newPos := robot.Add(op)
		if ob, ok := floor[*newPos]; !ok {
			// move robot to empty space
			robot = *newPos
			continue
		} else if ob == Wall {
			// robot hits a wall
			continue
		} else if ob == Box && MoveBox(*newPos, op, &floor) {
			// robot moves a box
			robot = *newPos
		}
	}

	score := 0
	for pos, ob := range floor {
		if ob == Box {
			score += (pos.Y * 100) + pos.X
		}
	}
	return score
}

func Part2(filename string) int {
	return 0
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 186996)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 221683913164898)\n", ans2)
}
