package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	A = "A"
	B = "B"
	C = "C"
	X = "X"
	Y = "Y"
	Z = "Z"
)

func won(ret []string) int {
	if (ret[0] == A && ret[1] == Y) ||
		(ret[0] == B && ret[1] == Z) ||
		(ret[0] == C && ret[1] == X) {
		return 6
	}
	return 0
}

func draw(ret []string) int {
	if (ret[0] == A && ret[1] == X) ||
		(ret[0] == B && ret[1] == Y) ||
		(ret[0] == C && ret[1] == Z) {
		return 3
	}
	return 0
}

func shapeSelected(ret []string) int {
	if ret[1] == X {
		return 1
	}
	if ret[1] == Y {
		return 2
	}
	return 3
}

func isRock(ret []string) int {
	if ret[0] == A && ret[1] == Y ||
		ret[0] == B && ret[1] == X ||
		ret[0] == C && ret[1] == Z {
			return 1
		}
	return 0
}

func isPaper(ret []string) int {
	if ret[0] == A && ret[1] == Z ||
		ret[0] == B && ret[1] == Y ||
		ret[0] == C && ret[1] == X {
			return 2
		}
	return 0
}

func isScissors(ret []string) int {
	if ret[0] == A && ret[1] == X ||
		ret[0] == B && ret[1] == Z ||
		ret[0] == C && ret[1] == Y {
			return 3
		}
	return 0
}

func wonOrDraw(ret []string) int {
	if ret[1] == Y {
		return 3
	}
	if ret[1] == Z {
		return 6
	}
	return 0
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
		dup := strings.Split(string(datString[i]), " ")
		ret1 += shapeSelected(dup) + draw(dup) + won(dup)
		ret2 += wonOrDraw(dup) + isRock(dup) + isPaper(dup) + isScissors(dup)
	}
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
