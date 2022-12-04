package main

import (
	"fmt"
	"os"
	"strings"
)

func won(ret []string) int {
	if (ret[0] == "A" && ret[1] == "Y") ||
		(ret[0] == "B" && ret[1] == "Z") ||
		(ret[0] == "C" && ret[1] == "X") {
		return 6
	}
	return 0
}

func draw(ret []string) int {
	if (ret[0] == "A" && ret[1] == "X") ||
		(ret[0] == "B" && ret[1] == "Y") ||
		(ret[0] == "C" && ret[1] == "Z") {
		return 3
	}
	return 0
}

func shapeSelected(ret []string) int {
	if ret[1] == "X" {
		return 1
	}
	if ret[1] == "Y" {
		return 2
	}
	return 3
}

func main() {
	var ret [][]string
	var ret1 int
	var ret2 int
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n")
	for i := 0; i < len(datString) - 1; i++ {
		dup := strings.Split(string(datString[i]), " ")
		ret1 += shapeSelected(dup) + draw(dup) + won(dup)
		ret = append(ret, dup)
	}
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
