package main

import "testing"

func TestPartOne000(t *testing.T) {
	ans := 2
	rTest := Part1("test000.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne00(t *testing.T) {
	ans := 4
	rTest := Part1("test00.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne0(t *testing.T) {
	ans := 3
	rTest := Part1("test0.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartOne(t *testing.T) {
	ans := 36
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}

func TestPartTwo(t *testing.T) {
	ansTest := 81
	rTest := Part2("test.txt")
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
}
