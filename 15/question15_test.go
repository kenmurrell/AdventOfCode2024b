package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 10092
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
