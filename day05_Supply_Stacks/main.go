package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	crate []string
}

func (s *stack) pushOne(crate string) {
	s.crate = append(s.crate, crate)
}

func (s *stack) popOne() (string) {
	ret := s.crate[len(s.crate) - 1]
	s.crate = s.crate[:len(s.crate) - 1]
	return ret
}

func (s *stack) pushTwo(crate []string) {
	s.crate = append(s.crate, crate...)
}

func (s *stack) popTwo(toMove int) ([]string) {
	ret := s.crate[len(s.crate) - toMove : len(s.crate)]
	s.crate = s.crate[:len(s.crate) - toMove]
	return ret
}

func (s *stack) addToBotton(crate string) {
	s.crate = append([]string{crate}, s.crate...)
}

func calcNumOfStacks(s string) (int) {
	n := strings.Split(s, " ")
	numStacks, _ := strconv.Atoi(n[len(n) - 2])
	return numStacks
}

func createStack(n int, s []string) ([]stack) {
	stacks := make([]stack, n)
	for i := 0; i < len(s) - 1; i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] != ' ' && s[i][j] != '[' && s[i][j] != ']' {
				stacks[j / 4].addToBotton(string(s[i][j]))
			}
		}
	}
	return stacks
}

func partOne(cmd []string, n int, s []string) {
	stacks := createStack(n, s)
	for i := 0; i < len(cmd) - 1; i++ {
		var toMove, from, to int
		fmt.Sscanf(cmd[i], "move %d from %d to %d", &toMove, &from, &to)
		for move := 0; move < toMove; move++ {
			stacks[to - 1].pushOne(stacks[from - 1].popOne())
		}
	}
	var ret1 string
	for _, s := range stacks{
		ret1 = ret1 + fmt.Sprint(string(s.popOne()))
	}
	fmt.Println("--- Part One --- ", ret1)
}

func partTwo(cmd []string, n int, s []string) {
	stacks := createStack(n, s)
	for i := 0; i < len(cmd) - 1; i++ {
		var toMove, from, to int
		fmt.Sscanf(cmd[i], "move %d from %d to %d", &toMove, &from, &to)
		stacks[to - 1].pushTwo(stacks[from - 1].popTwo(toMove))
	}
	var ret2 string
	for _, s := range stacks{
		ret2 = ret2 + fmt.Sprint(string(s.popOne()))
	}
	fmt.Println("--- Part Two --- ", ret2)
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n\n")
	numOfStacks := calcNumOfStacks(datString[0])
	stackString := strings.Split(datString[0], "\n")
	cmd := strings.Split(datString[1], "\n")

	partOne(cmd, numOfStacks, stackString)
	partTwo(cmd, numOfStacks, stackString)
}
