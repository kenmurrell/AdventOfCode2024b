package main

import "testing"

func TestPartOne0(t *testing.T) {
	ans := 140
	rTest := Part1("test0.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne00(t *testing.T) {
	ans := 772
	rTest := Part1("test00.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne(t *testing.T) {
	ans := 1930
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}