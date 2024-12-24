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
	X float64
	Y float64
}

type Machine struct {
	A     Coordinate
	B     Coordinate
	Prize Coordinate
}

func CreateCoordinateFromText(text string) Coordinate {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(text, 2)
	x, _ := strconv.Atoi(matches[0])
	y, _ := strconv.Atoi(matches[1])
	return Coordinate{float64(x), float64(y)}
}

func load(filename string) []Machine {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var machines []Machine
	var buttonA Coordinate
	var buttonB Coordinate
	var prize Coordinate
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			machines = append(machines, Machine{buttonA, buttonB, prize})
		}
		sections := strings.Split(line, ":")
		switch sections[0] {
		case "Button A":
			buttonA = CreateCoordinateFromText(sections[1])
		case "Button B":
			buttonB = CreateCoordinateFromText(sections[1])
		case "Prize":
			prize = CreateCoordinateFromText(sections[1])
		}
	}
	machines = append(machines, Machine{buttonA, buttonB, prize})
	return machines
}

func Part1(filename string) int {
	machines := load(filename)
	tokens := 0
	for _, m := range machines {
		buttonA := ((m.B.Y * m.Prize.X) - (m.B.X * m.Prize.Y)) / ((m.A.X * m.B.Y) - (m.B.X * m.A.Y))
		buttonB := (m.Prize.X - (m.A.X * buttonA)) / (m.B.X)
		buttonAInt := int(buttonA)
		buttonBInt := int(buttonB)
		if buttonA == float64(buttonAInt) && buttonB == float64(buttonBInt) {
			tokens += (buttonAInt * 3) + buttonBInt
		}
	}
	return tokens
}

func Part2(filename string) uint64 {
	machines := load(filename)
	var tokens uint64 = 0
	for _, m := range machines {
		prizeX := m.Prize.X + 10000000000000
		prizeY := m.Prize.Y + 10000000000000
		buttonA := ((m.B.Y * prizeX) - (m.B.X * prizeY)) / ((m.A.X * m.B.Y) - (m.B.X * m.A.Y))
		buttonB := (prizeX - ((m.A.X) * buttonA)) / (m.B.X)
		buttonAInt := uint64(buttonA)
		buttonBInt := uint64(buttonB)
		if buttonA == float64(buttonAInt) && buttonB == float64(buttonBInt) {
			tokens += (buttonAInt * 3) + buttonBInt
		}
	}
	return tokens
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 36870)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 78101482023732)\n", ans2)
}
