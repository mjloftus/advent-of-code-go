package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	x, y int
}

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	m := make(map[pair]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		line2 := strings.Split(line[1], " ")
		x1, _ := strconv.Atoi(line[0])
		y1, _ := strconv.Atoi(line2[0])
		x2, _ := strconv.Atoi(line2[2])
		y2, _ := strconv.Atoi(line[2])
		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				if y1 > y2 {
					for i := y1; i >= y2; i-- {
						m[pair{x1, i}]++
					}
				} else {
					for i := y2; i >= y1; i-- {
						m[pair{x1, i}]++
					}
				}
			} else {
				if x1 > x2 {
					for i := x1; i >= x2; i-- {
						m[pair{i, y1}]++
					}
				} else {
					for i := x2; i >= x1; i-- {
						m[pair{i, y1}]++
					}
				}
			}
		}
	}
	var moreThanOne = 0
	for i := range m {
		if m[i] >= 2 {
			moreThanOne++
		}
	}
	return moreThanOne
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	m := make(map[pair]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		line2 := strings.Split(line[1], " ")
		x1, _ := strconv.Atoi(line[0])
		y1, _ := strconv.Atoi(line2[0])
		x2, _ := strconv.Atoi(line2[2])
		y2, _ := strconv.Atoi(line[2])
		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				if y1 > y2 {
					for i := y1; i >= y2; i-- {
						m[pair{x1, i}]++
					}
				} else {
					for i := y2; i >= y1; i-- {
						m[pair{x1, i}]++
					}
				}
			} else {
				if x1 > x2 {
					for i := x1; i >= x2; i-- {
						m[pair{i, y1}]++
					}
				} else {
					for i := x2; i >= x1; i-- {
						m[pair{i, y1}]++
					}
				}
			}
		} else if math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2)) {
			println("here1")
			if x1 > x2 {
				println("here2")
				if y1 > y2 {
					println("here3")
					for i := x1; i >= x2; i-- {
						m[pair{i,y1 - (x1 - i)}]++
					}
				} else {
					println("here4")
					for i := x1; i >= x2; i-- {
						m[pair{i,y1 + (x1 - i)}]++
					}
				}
			} else {
				if y1 > y2 {
					for i := x2; i >= x1; i-- {
						m[pair{i,y2 + (x2 - i)}]++
					}
				} else {
					for i := x2; i >= x1; i-- {
						m[pair{i,y2 - (x2 - i)}]++
					}
				}
			}
		}
	}
	var moreThanOne = 0
	for i := range m {
		if m[i] >= 2 {
			moreThanOne++
		}
	}
	return moreThanOne
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}