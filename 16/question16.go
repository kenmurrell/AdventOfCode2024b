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

var East Position = Position{1, 0}
var West Position = Position{-1, 0}
var South Position = Position{0, 1}
var North Position = Position{0, -1}
var moves []*Position = []*Position{&East, &West, &South, &North}

type empty struct{}

func (p1 *Position) Add(p2 *Position) *Position {
	return &Position{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Position) Equal(p2 *Position) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (p1 *Position) IsOpp(p2 *Position) bool {
	return p1.X+p2.X == 0 && p1.Y+p2.Y == 0
}

func load(filename string) (Position, *map[Position]bool) {
	file, _ := os.Open(filename)
	defer file.Close()
	maze := make(map[Position]bool)
	scanner := bufio.NewScanner(file)
	var start Position
	y := -1
	for scanner.Scan() {
		y += 1
		line := scanner.Text()
		for x, c := range strings.Split(line, "") {
			p := Position{x, y}
			switch c {
			case ".":
				maze[p] = false
			case "S":
				start = p
			case "E":
				maze[p] = true
			}
		}
	}
	return start, &maze
}

func Part1(filename string) int {
	start, maze := load(filename)
	visited := make(map[Position]empty)
	tracker := make(chan empty)
	go func() {
		c := 0
		for range tracker {
			c += 1
			fmt.Printf("%d paths found\n", c)
		}
	}()
	count, _ := Traverse(&start, &East, maze, &visited, tracker)
	close(tracker)
	return count
}

func Traverse(
	currentPos *Position,
	lastDir *Position,
	maze *map[Position]bool,
	visited *map[Position]empty,
	tracker chan<- empty) (int, bool) {
	(*visited)[*currentPos] = empty{}
	minScore := -1
	foundEnd := false
	for _, move := range moves {
		//find next move
		nextPos := currentPos.Add(move)
		//check if visited
		if _, ok := (*visited)[*nextPos]; ok {
			continue
		}
		if isEnd, canMove := (*maze)[*nextPos]; canMove {
			// find score
			var score int
			if move.Equal(lastDir) {
				score = 1
			} else if !move.IsOpp(lastDir) {
				score = 1001
			}

			if isEnd {
				tracker <- empty{}
				return score, true
			}
			nextVisited := make(map[Position]empty)
			for k, v := range *visited {
				nextVisited[k] = v
			}
			if endScore, found := Traverse(nextPos, move, maze, &nextVisited, tracker); found {
				foundEnd = true
				if endScore+score < minScore || minScore == -1 {
					minScore = endScore + score
				}
			}
		}
	}
	return minScore, foundEnd
}

func Part2(filename string) int {
	return 0
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected )\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected )\n", ans2)
}
