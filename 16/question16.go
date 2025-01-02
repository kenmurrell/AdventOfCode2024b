package main

import (
	"bufio"
	"fmt"
	"math"
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

func load(filename string) (Position, Position, *map[Position]int) {
	file, _ := os.Open(filename)
	defer file.Close()
	maze := make(map[Position]int)
	scanner := bufio.NewScanner(file)
	var start Position
	var end Position
	y := -1
	for scanner.Scan() {
		y += 1
		line := scanner.Text()
		for x, c := range strings.Split(line, "") {
			p := Position{x, y}
			switch c {
			case ".":
				maze[p] = -1
			case "S":
				maze[p] = -1
				start = p
			case "E":
				maze[p] = -1
				end = p
			}
		}
	}
	return start, end, &maze
}

func Part1(filename string) int {
	start, end, maze := load(filename)
	found, count := FindMinimumScore(&start, &end, &East, 0, maze)
	if !found {
		panic("No end found")
	}
	return count
}

func FindMinimumScore(
	currentPos,
	targetPos,
	lastDir *Position,
	score int,
	maze *map[Position]int) (bool, int) {

	if currentPos.Equal(targetPos) {
		return true, score
	} else if lastScore, isvalid := (*maze)[*currentPos]; !isvalid || (lastScore != -1 && lastScore < score) {
		return false, 0
	} else {
		(*maze)[*currentPos] = score
	}

	var minScore int = 0
	var found bool = false
	for _, move := range moves {
		//find next move
		nextPos := currentPos.Add(move)

		var nextScore int
		if move.Equal(lastDir) {
			nextScore = score + 1
		} else if move.IsOpp(lastDir) {
			continue // don't go backwards
		} else {
			nextScore = score + 1001
		}

		if r, endScore := FindMinimumScore(nextPos, targetPos, move, nextScore, maze); r {
			found = true
			if minScore == 0 || minScore > endScore {
				minScore = endScore
			}
		}
	}
	return found, minScore
}

func Part2(filename string) int {
	start, end, maze := load(filename)
	_, paths := TrackBestPath(&start, &end, &East, 0, maze)
	return len(*paths)
}

func TrackBestPath(
	currentPos,
	targetPos,
	lastDir *Position,
	score int,
	maze *map[Position]int) (int, *map[Position]empty) {
	if currentPos.Equal(targetPos) {
		path := make(map[Position]empty)
		path[*currentPos] = empty{}
		return score, &path
	} else if lastScore, isvalid := (*maze)[*currentPos]; !isvalid || (lastScore != -1 && lastScore+1000 /*just awful*/ < score) {
		return 0, nil
	} else {
		(*maze)[*currentPos] = score
	}
	var minScore int = math.MaxInt32
	var bestpath *map[Position]empty
	for _, move := range moves {
		//find next move
		nextPos := currentPos.Add(move)
		var nextScore int
		if move.Equal(lastDir) {
			nextScore = score + 1
		} else if move.IsOpp(lastDir) {
			continue // don't go backwards
		} else {
			nextScore = score + 1001
		}

		if endScore, path := TrackBestPath(nextPos, targetPos, move, nextScore, maze); path != nil {
			(*path)[*currentPos] = empty{}
			if endScore < minScore {
				minScore = endScore
				bestpath = path
			} else if endScore == minScore {
				Merge(bestpath, path)
			}
		}
	}
	return minScore, bestpath
}

func Merge(A *map[Position]empty, B *map[Position]empty) {
	for p, _ := range *B {
		(*A)[p] = empty{}
	}
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 85480)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 518)\n", ans2)
}
