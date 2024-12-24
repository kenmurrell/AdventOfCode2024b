package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 480
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}
