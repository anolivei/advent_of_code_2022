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

func partOne(s string) int {
	var letter string
	for i := 0; i < len(s) / 2; i++ {
		for j := len(s) / 2; j < len(s); j++ {
			if s[i] == s[j] {
				letter = string(s[i])
				break
			}
		}
	}
	return int(decode(letter))
}

func partTwo(s1, s2, s3 string) int {
	var letter string
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			for k:= 0; k < len(s3); k++ {
				if s1[i] == s2[j] && s1[i] == s3[k] {
					letter = string(s1[i])
					break
				}
			}
		}
	}
	return int(decode(letter))
}

func main() {
	var ret1 int
	var ret2 int

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n")
	for i := 0; i < len(datString) - 1; i++ {
		ret1 += partOne(datString[i])
	}
	for i := 0; i < len(datString) - 3; i += 3 {
		ret2 += partTwo(datString[i], datString[i + 1], datString[i + 2])
	}
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
