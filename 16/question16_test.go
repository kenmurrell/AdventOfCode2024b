package main

import "testing"

func TestPartOne1(t *testing.T) {
	ans := 7036
	rTest := Part1("test1.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne2(t *testing.T) {
	ans := 11048
	rTest := Part1("test2.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo1(t *testing.T) {
	ans := 45
	rTest := Part2("test1.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo2(t *testing.T) {
	ans := 64
	rTest := Part2("test2.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
