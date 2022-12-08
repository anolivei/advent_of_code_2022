package main

import (
	"fmt"
	"os"
)

func thereIsDuplicate(s string) (bool) {
	m := make(map[rune]uint, len(s))
	for _, r := range s {
		m[r]++
		if m[r] > 1 {
			return true
		}
	}
	return false
}

func solve(input string, n int) (int) {
	for i := 0; i < len(input) - n - 1; i++ {
		s := input[i : i + n]
		if !thereIsDuplicate(s) {
			return i + n
		}
	}
	return 0
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := string(dat[:len(dat) - 1])
	fmt.Println("--- Part One --- ", solve(datString, 4))
	fmt.Println("--- Part Two --- ", solve(datString, 14))
}
