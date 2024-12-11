package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	ID int
}

func (i *Item) isEmpty() bool {
	return i.ID == -1
}

func CheckSum(fs []Item) int {
	checksum := 0
	for i, _ := range fs {
		if !fs[i].isEmpty() {
			checksum += i * fs[i].ID
		}
	}
	return checksum
}

func load(filename string) []Item {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fs []Item
	id_iter := 0
	isEmpty := false
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), "") {
			i, _ := strconv.Atoi(c)
			var item Item
			if isEmpty {
				item = Item{-1}
			} else {
				item = Item{id_iter}
				id_iter += 1
			}
			for t := 0; t < i; t++ {
				fs = append(fs, item)
			}
			isEmpty = !isEmpty
		}
	}
	return fs
}

func Part1(filename string) int {
	fs := load(filename)
	first, last := 0, len(fs)-1
	for {
		if first >= last {
			break
		}
		if !fs[first].isEmpty() {
			first += 1
			continue
		} else if fs[last].isEmpty() {
			last -= 1
			continue
		} else {
			empty := fs[first]
			fs[first] = fs[last]
			fs[last] = empty
			first += 1
			last -= 1
		}
	}
	return CheckSum(fs)
}

func FindFirstEmptyRegion(fs []Item, idx int, size int) int {
	first, count := -1, 0
	for i := 0; i < idx; i++ {
		if fs[i].isEmpty() {
			if first < 0 {
				first = i
			}
			count += 1
			if count == size {
				return first
			}
		} else {
			first = -1
			count = 0
		}
	}
	return -1
}

func Swap(fs *[]Item, first int, last int, count int) {
	for i := 0; i < count; i++ {
		(*fs)[first+i], (*fs)[last-i] = (*fs)[last-i], (*fs)[first+i]
	}
}

func Part2(filename string) int {
	fs := load(filename)
	item, idx := Item{-1}, -1
	for i := len(fs) - 1; i >= 0; i-- {
		if fs[i].ID == item.ID {
			continue
		} else {
			if !item.isEmpty() {
				emptyOffset := FindFirstEmptyRegion(fs, idx, idx-i)
				if emptyOffset >= 0 {
					Swap(&fs, emptyOffset, idx, idx-i)
				}
			}
			item = fs[i]
			idx = i
		}
	}

	return CheckSum(fs)
}

func main() {
	ans1 := Part1("data.txt")
	fmt.Printf("ANSWER ONE: %d (expected 6291146824486)\n", ans1)
	ans2 := Part2("data.txt")
	fmt.Printf("ANSWER TWO: %d (expected 6307279963620)\n", ans2)
}
