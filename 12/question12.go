package main

import (
	"bufio"
	"fmt"
	"os"
)

type empty struct {}

type Coordinate struct {
	X int
	Y int
}

type Garden struct {
	garden *[][]rune
}

func (g *Garden) Get(c *Coordinate) rune {
	if c.Y >= 0 && c.Y < len(*g.garden) && c.X >= 0 && c.X < len((*g.garden)[0]) {
		return (*g.garden)[c.Y][c.X]
	} 
	return -1
}

func load(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var garden [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		iline := make([]rune, len(line))
		for i, c := range []rune(line) {
			iline[i] = c
		}
		garden = append(garden, iline)
	}
	return garden
}

func Traverse(
	garden *Garden,
	c Coordinate,
	v *map[Coordinate]empty) (int, int) {
	if _, ok := (*v)[c]; ok {
		return 0, 0
	}
	(*v)[c] = empty{}
	perimeter := 0 
	area := 1
	char := garden.Get(&c)
	checkAndAdd := func(c Coordinate){
		if new := garden.Get(&c); new == char{
			a, p := Traverse(garden, c, v)
			area, perimeter = area + a, perimeter + p
		} else {
			perimeter += 1
		}
	}
	
	checkAndAdd(Coordinate{c.X, c.Y-1})
	checkAndAdd(Coordinate{c.X+1, c.Y})
	checkAndAdd(Coordinate{c.X, c.Y+1})
	checkAndAdd(Coordinate{c.X-1, c.Y})
	return area, perimeter
}

func Part1(filename string) int {
	total := 0
	g := load(filename)
	garden := Garden{&g}
	visited := make(map[Coordinate]empty)
	for row := 0; row<len(g);row++ {
		for col := 0; col < len(g[0]); col ++ {
			area, perimeter := Traverse(&garden, Coordinate{row, col}, &visited)
			total += (area * perimeter)
		}
	}
	return total
}

func Part2(filename string) int {
	return 0
}


func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 1381056)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected )\n", ans2)
}
