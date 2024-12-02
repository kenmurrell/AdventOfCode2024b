package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ans := 2
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 4
	rTest := Part2("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
}
