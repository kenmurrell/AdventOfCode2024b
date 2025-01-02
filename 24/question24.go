package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	A  string
	B  string
	C  string
	OP string
}

func (t *Task) tostring() string {
	return fmt.Sprintf("%s-%s-%s-%s", t.A, t.OP, t.B, t.C)
}

func load(filename string) (*[]Task, *map[string]int) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wires := make(map[string]int)
	tasks := []Task{}
	initial_values := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			initial_values = false
			continue
		}
		if initial_values {
			vals := strings.Split(line, ": ")
			val, _ := strconv.Atoi(vals[1])
			wires[vals[0]] = val
		} else {
			vals := strings.Split(line, " ")
			tasks = append(tasks, Task{A: vals[0], OP: vals[1], B: vals[2], C: vals[4]})
		}

	}
	return &tasks, &wires
}

func Part1(filename string) int {
	tasks, wires := load(filename)
	completed := make(map[string]bool)
	iter := 0
	for {
		if len(completed) == len(*tasks) {
			break
		}
		task := (*tasks)[iter%len(*tasks)]
		iter += 1
		if completed[task.tostring()] {
			continue
		}

		A_val, ok1 := (*wires)[task.A]
		B_val, ok2 := (*wires)[task.B]
		if !ok1 || !ok2 {
			continue
		}
		var c_val int
		switch task.OP {
		case "AND":
			c_val = A_val & B_val
		case "XOR":
			c_val = A_val ^ B_val
		case "OR":
			c_val = A_val | B_val
		default:
			fmt.Printf("err: %d\n", c_val)
		}
		(*wires)[task.C] = c_val
		completed[task.tostring()] = true
	}
	var num int
	for w, val := range *wires {
		if strings.HasPrefix(w, "z") {
			offset, _ := strconv.Atoi(w[1:])
			num += val << offset
		}
	}
	return num
}

func Part2(filename string) int {
	return 0
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 49574189473968)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected )", ans2)
}
