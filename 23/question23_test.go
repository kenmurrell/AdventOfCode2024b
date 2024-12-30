package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 7
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ans := "co,de,ka,ta"
	rTest := Part2("test.txt")
	if rTest != ans {
		t.Errorf("Got %s; want %s", rTest, ans)
	}
}
