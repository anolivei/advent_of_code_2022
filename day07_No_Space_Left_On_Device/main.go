package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	MAX_SIZE = 100000
	TOTAL_DISK_SPACE = 70000000
	SPACE_TO_UPDATE = 30000000
)

type directory struct {
	name string
	size int
	isFile bool
	subDir map[string] *directory
	upperDir *directory
}

func calcSize(dir directory) (int) {
	if dir.isFile {
		return dir.size
	}
	size := 0
	for _, d := range dir.subDir {
		size += calcSize(*d)
	}
	return size
}

func parseInput(input []string, currentDir *directory, dirs []*directory) ([]*directory) {
	for i := 0; i < len(input) - 1; i++ {
		line := strings.Split(input[i], " ")
		if line[1] == "cd" {
			if line[2] == ".." {
				currentDir = currentDir.upperDir
			} else if line[2] == "/" {
				currentDir = &directory{
										"/",
										0,
										false,
										make(map[string]*directory),
										nil}
				dirs = append(dirs, currentDir)
			} else {
				currentDir = currentDir.subDir[line[2]]
			}
		} else if line[0] == "dir" {
			currentDir.subDir[line[1]] = &directory{
													line[1],
													0,
													false,
													make(map[string]*directory),
													currentDir}
			dirs = append(dirs, currentDir.subDir[line[1]])
		} else if line[1] != "ls" {
			size, _ := strconv.Atoi(line[0])
			currentDir.subDir[line[1]] = &directory{
													line[1],
													size,
													true,
													nil,
													currentDir}
		}
	}
	return dirs
}

func partOne(dirs []*directory) {
	totalSize := 0
	for _, dir := range dirs {
		size := calcSize(*dir)
		if size <= MAX_SIZE {
			totalSize += size
		}
	}
	fmt.Println("--- Part One --- ", totalSize)
}

func partTwo(dirs []*directory) {
	rootSize := calcSize(*dirs[0])
	toFree:= SPACE_TO_UPDATE - (TOTAL_DISK_SPACE - rootSize)
	var toDelete int = rootSize
	for _, dir := range dirs {
		size := calcSize(*dir)
		if size > toFree && size < toDelete {
			toDelete = size
		}
	}
	fmt.Println("--- Part Two --- ", toDelete)
}

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(dat), "\n")
	var currentDir *directory
	dirs := []*directory{}
	dirs = parseInput(input, currentDir, dirs)
	partOne(dirs)
	partTwo(dirs)
}
