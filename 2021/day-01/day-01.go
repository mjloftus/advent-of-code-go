package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func partOne() int {
	input, _ := ioutil.ReadFile("input.txt")
	depths := strings.Split(string(input), "\n")

	increases := 0
	for ind, val := range depths {
		if ind == 0 {
			continue
		}
		curr, _ := strconv.Atoi(val)
		prev, _ := strconv.Atoi(depths[ind-1])
		if curr > prev {
			increases++
		}
	}
	return increases
}

func partTwo() int {
	input, _ := ioutil.ReadFile("input.txt")
	depths := strings.Split(string(input), "\n")

	increases := 0
	sum := 0
	for i := 0; i < 3; i++ {
		val, _ := strconv.Atoi(depths[i])
		sum += val
	}
	for ind, val := range depths {
		if ind < 3 {
			continue
		}
		d, _ := strconv.Atoi(depths[ind-3])
		v, _ := strconv.Atoi(val)
		if v > d {
			increases++
		}
		sum = sum - d + v
	}
	return increases
}

func main() {
	println(partOne())
	println(partTwo())
	return
}