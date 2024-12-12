package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 55312
	rTest := Part1("test.txt")
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}