package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func load(filename string) []uint64 {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var secrets []uint64
	for scanner.Scan() {
		line := scanner.Text()
		in, _ := strconv.Atoi(line)
		secrets = append(secrets, uint64(in))
	}
	return secrets
}

func mix(n1, n2 uint64) uint64 {
	return n1 ^ n2
}

func prune(n1 uint64) uint64 {
	return n1 % 16777216
}

func Part1(filename string) uint64 {
	secrets := load(filename)
	gens := 2000
	for i, secret := range secrets {
		s := secret
		for _ = range gens {
			// Step 1: multiply by 64
			s = prune(mix(s, s*64))

			// Step 2: divide by 32
			s = prune(mix(s, s/32))

			// Step 3: divide by 32
			s = prune(mix(s, s*2048))
		}
		secrets[i] = s
	}
	var sum uint64 = 0
	for _, secret := range secrets {
		sum += secret
	}
	return sum
}

func Part2(filename string, gens int) int {
	secrets := load(filename)
	srank := make(map[[4]int]int)
	for _, secret := range secrets {
		visited := make(map[[4]int]bool)
		s := secret
		tiers := make([]uint64, gens+1)
		tiers[0] = s
		for i := 0; i < gens; i++ {
			s = prune(mix(s, s*64))
			s = prune(mix(s, s/32))
			s = prune(mix(s, s*2048))
			tiers[i+1] = s
		}
		for i := 0; i <= len(tiers)-5; i++ {
			sequence := tiers[i : i+5]
			key := [4]int{
				diff(sequence[1], sequence[0]),
				diff(sequence[2], sequence[1]),
				diff(sequence[3], sequence[2]),
				diff(sequence[4], sequence[3]),
			}
			if !visited[key] {
				srank[key] += int(sequence[4] % 10)
				visited[key] = true
			}
		}

	}
	var sbest int = 0
	for _, count := range srank {
		if count > sbest {
			sbest = count
		}
	}
	return sbest
}

func diff(n2, n1 uint64) int {
	return int((n2 % 10) - (n1 % 10))
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 17965282217)\n", ans1)
	ans2 := Part2("data.txt", 2001)
	fmt.Printf("ANSWER TWO: %d (expected 2152)\n", ans2)
}
