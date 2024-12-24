package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Robot struct {
	Pos *Coordinate
	Vel *Coordinate
}

func CreateRobot(posX int, posY int, velx int, vely int) *Robot {
	pos := Coordinate{posX, posY}
	vel := Coordinate{velx, vely}
	return &Robot{&pos, &vel}
}

func (r *Robot) Tick(boundX int, boundY int) {
	newX := (boundX + r.Pos.X + r.Vel.X) % boundX
	newY := (boundY + r.Pos.Y + r.Vel.Y) % boundY
	r.Pos = &Coordinate{newX, newY}
}

func load(filename string) []*Robot {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var robots []*Robot
	pat := regexp.MustCompile(`p\=(\d+)\,(\d+) v\=(-?\d+)\,(-?\d+)`)
	for scanner.Scan() {
		matches := pat.FindAllStringSubmatch(scanner.Text(), -1)
		posX, _ := strconv.Atoi(matches[0][1])
		posY, _ := strconv.Atoi(matches[0][2])
		velX, _ := strconv.Atoi(matches[0][3])
		velY, _ := strconv.Atoi(matches[0][4])
		robots = append(robots, CreateRobot(posX, posY, velX, velY))
	}
	return robots
}

func Part1(filename string, boundX int, boundY int) int {
	ticks := 100
	robots := load(filename)
	for _ = range ticks {
		for _, r := range robots {
			r.Tick(boundX, boundY)
		}
	}
	quadrants := []int{0,0,0,0}
	for _, r := range robots {
		id := 0
		middleX := (boundX /2)
		middleY := (boundY /2)
		if middleX == r.Pos.X || middleY == r.Pos.Y{
			continue 
		}
		if middleX < r.Pos.X {
			id += 1
		}
		if middleY < r.Pos.Y {
			id += 2
		}
		quadrants[id] += 1
	}

	sum := 1
	for i, _ := range quadrants {
		sum *= quadrants[i]
	}
	return sum
}

func Part2(filename string, boundX int, boundY int, ticks int, limit int ) {
	robots := load(filename)
	for t := range ticks {
		for _, r := range robots {
			r.Tick(boundX, boundY)
		}
		// shitty code incoming...
		var m [][]int
		for i:=0; i<=101; i++ {
			var row []int
			for j :=0; j<103; j++ {
				row = append(row, 0)
			}
			m = append(m, row)
		}
		for _, r := range robots {
			m[r.Pos.X][r.Pos.Y] += 1
		}
		var s strings.Builder
		count := 0
		for i:=0; i<len(m); i++ {
			for j:=0; j<len(m[0]); j++ {
				
				if m[i][j] > 0 {
					if i > 0 && i < len(m)-1 && j > 0 && j < len(m[0])-1 {
						if m[i+1][j] > 0 {
							count+=1
						}
						if m[i-1][j] > 0 {
							count+=1
						}
						if m[i][j+1] > 0 {
							count+=1
						}
						if m[i][j-1] > 0 {
							count+=1
						}
					}
					s.WriteRune('X')
				} else {
					s.WriteRune('.')
				}
			}
			s.WriteString("\n")
		}
		if count > limit {
			fmt.Println(s.String())
			fmt.Println(count)
			fmt.Printf("%d\n", t+1)
		}

	}
	
}

func main() {
	ticks, _ := strconv.Atoi(os.Args[1])
	limit, _ := strconv.Atoi(os.Args[2])
	ans1 := Part1("data.txt", 101, 103)
	fmt.Printf("ANSWER ONE: %d (expected 225552000)\n", ans1)
	Part2("data.txt", 101, 103, ticks, limit)
}
