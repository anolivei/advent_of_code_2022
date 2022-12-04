package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func decode(s string) rune {
	var num rune
	for _, r := range s {
		if unicode.IsUpper(r) {
			y := r - 38
			num = y
		}
		if unicode.IsLower(r) {
			z := r - 96
			num = z
		}
	}
	return num
}

func main() {
	var ret1 int
	var ret2 int
	var letter string

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n")
	for i := 0; i < len(datString) - 1; i++ {
		for j := 0; j < len(datString[i]) / 2; j++ {
			for k := len(datString[i]) / 2; k < len(datString[i]); k++ {
				if datString[i][j] == datString[i][k] {
					letter = string(datString[i][j])
					break
				}
		}}
		ret1 += int(decode(letter))
	}
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
