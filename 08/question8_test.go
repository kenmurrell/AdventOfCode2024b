package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 14
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo0(t *testing.T) {
	ansTest := 9
	rTest := Part2("test0.txt")
	if int(rTest) != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 34
	rTest := Part2("test.txt")
	if int(rTest) != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
}
