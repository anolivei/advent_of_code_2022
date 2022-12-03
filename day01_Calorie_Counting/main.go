package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	var ret []int
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	datString := strings.Split(string(dat), "\n")
	for i := 0; i < len(datString); i++ {
		sum := 0
		for datString[i] != "" {
				n, err := strconv.Atoi(datString[i])
				if err != nil {
					panic(err)
				}
				sum += n
				i++
		}
		ret = append(ret, sum)
	}
	sort.Ints(ret)
	ret1 := ret[len(ret) - 1]
	ret2 := ret[len(ret) - 1] + ret[len(ret) - 2] + ret[len(ret) - 3]
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)
}
