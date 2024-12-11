package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type empty struct {
}

type Coordinate struct {
	X int
	Y int
}

type Velocity struct {
	C Coordinate
	D Direction
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

func checkloop(
	max_x int,
	max_y int,
	startPos Coordinate,
	startDir Direction,
	walls *map[Coordinate]empty) bool {
	visited := make(map[Velocity]empty)
	visited[Velocity{startPos, startDir}] = empty{}
	for {
		next := Next(startDir, &startPos)
		if _, ok := (*walls)[next]; ok {
			startDir = (startDir + 1) % 4
			continue
		} else if next.IsOutside(max_x, max_y) {
			return false
		}
		startPos = next
		vel := Velocity{startPos, startDir}
		if _, ok := visited[vel]; ok {
			return true
		} else {
			visited[vel] = empty{}
		}
	}
}

func copy(m1 *map[Coordinate]empty) *map[Coordinate]empty {
	m2 := make(map[Coordinate]empty)
	for k, v := range *m1 {
		m2[k] = v
	}
	return &m2
}

// screw it, lets just bruteforce it
func Part2(filename string) (int, int) {
	max_x, max_y, startPos, startDir, walls := load(filename)
	start := time.Now()
	count := 0
	in := make(chan *map[Coordinate]empty)
	out := make(chan bool)
	var wg sync.WaitGroup
	for _ = range 10 {
		wg.Add(1)
		go func(wallChan <-chan *map[Coordinate]empty, out chan<- bool) {
			for w := range in {
				out <- checkloop(max_x, max_y, startPos, startDir, w)
			}
			wg.Done()
		}(in, out)
	}
	go func() {
		for b := range out {
			if b {
				count += 1
			}
		}
	}()
	for i := range max_x * max_y {
		if i%1000 == 0 {
			fmt.Printf("%d/%d, count:%d\n", i, max_x*max_y, count)
		}
		x := i % max_x
		y := int(i / max_x)
		c := Coordinate{x, y}
		if _, ok := walls[c]; !ok {
			w2 := copy(&walls)
			(*w2)[c] = empty{}
			in <- w2
		}
	}
	close(in)
	wg.Wait()
	close(out)
	return count, int(time.Since(start).Abs().Seconds())
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 4890)\n", ans1)
	ans2, s := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (%d seconds) (expected 1995)\n", ans2, s)
}
