package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 3749
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 11387
	rTest := Part2("test.txt")
	if int(rTest) != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
}
