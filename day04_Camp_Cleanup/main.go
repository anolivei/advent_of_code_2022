package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createArray(s []string) []int {
	var ret []int
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	for i := start; i <= end; i++ {
		ret = append(ret, i)
	}
	return ret
}

func partOne(first, second []int) int {
	exists := make(map[int]bool)
	if len(first) > len(second) {
		for _, value := range first {
			exists[value] = true
		}
		for _, value := range second {
			if !exists[value] {
				return 0
			}
		}
	} else {
		for _, value := range second {
			exists[value] = true
		}
		for _, value := range first {
			if !exists[value] {
				return 0
			}
		}
	}
	return 1
}

func partTwo(first, second []int) int {
	exists := make(map[int]bool)
	if len(first) > len(second) {
		for _, value := range first {
			exists[value] = true
		}
		for _, value := range second {
			if exists[value] {
				return 1
			}
		}
	} else {
		for _, value := range second {
			exists[value] = true
		}
		for _, value := range first {
			if exists[value] {
				return 1
			}
		}
	}
	return 0
}

func main() {
	var ret1 int
	var ret2 int
	var s1 []int
	var s2 []int

	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n")
	for i := 0; i < len(datString) - 1; i++ {
		ret := strings.Split(string(datString[i]), ",")
		s1 = createArray(strings.Split(string(ret[0]), "-"))
		s2 = createArray(strings.Split(string(ret[1]), "-"))
		ret1 += partOne(s1, s2)
		ret2 += partTwo(s1, s2)
	}
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
