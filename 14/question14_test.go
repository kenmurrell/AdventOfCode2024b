package main

import "testing"

func TestPartOne(t *testing.T) {
	ans := 12
	rTest := Part1("test.txt", 11, 7)
	if rTest != ans {
		t.Errorf("Got %d; want %d", rTest, ans)
	}
}