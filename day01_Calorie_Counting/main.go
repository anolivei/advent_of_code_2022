package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
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
		for j = j; dats2[j] != ""; j++ {
				x, _ := strconv.Atoi(dats2[j])
				sum += x
		}
		ret = append(ret, sum)
	}
	r := 0
	for i := 0; i < len(ret); i++ {
		if ret[i] > r {
			r = ret[i]
		}
	}
	fmt.Println(ret, r)
}