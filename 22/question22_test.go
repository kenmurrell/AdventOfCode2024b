package main

import "testing"

func TestPartOne(t *testing.T) {
	var ans uint64 = 37327623
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo0(t *testing.T) {
	var ans int = 6
	rTest := Part2("test0.txt", 10)
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	var ans int = 23
	rTest := Part2("test2.txt", 2001)
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
