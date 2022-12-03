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
	dats := string(dat)
	dats2 := strings.Split(dats, "\n")
	for j := 0; j < len(dats2); j++ {
		sum := 0
		for dats2[j] != "" {
				x, err := strconv.Atoi(dats2[j])
				if err != nil {
					panic(err)
				}
				sum += x
				j++
		}
		ret = append(ret, sum)
	}
	sort.Ints(ret)
	ret1 := ret[len(ret) - 1]
	ret2 := ret[len(ret) - 1] + ret[len(ret) - 2] + ret[len(ret) - 3]
	fmt.Println("--- Part One --- ", ret1)
	fmt.Println("--- Part Two --- ", ret2)

}