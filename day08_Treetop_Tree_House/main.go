package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./input_test.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")
	var trees [][]string
	for i := 0; i < len(input) - 1; i++ {
		trees = append(trees, strings.Split(input[i], ""))
		fmt.Println(trees)
	}
}
