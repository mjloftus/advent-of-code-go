package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	var pos []int
	min := 999
	max := -1
	for _, val := range vals{
		v, _ := strconv.Atoi(val)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		pos = append(pos, v)
	}
	minSum := -1
	for indI, i := range pos {
		total := 0
		for indJ, j := range pos {
			if indI == indJ {
				continue
			}
			total += int(math.Abs(float64(i - j)))
		}
		if total < minSum || minSum == -1 {
			minSum = total
		}
	}
	return minSum
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	var pos []int
	min := 999
	max := -1
	s := 0
	for _, val := range vals{
		v, _ := strconv.Atoi(val)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		pos = append(pos, v)
		s += v
	}
	minSum := -1
	distCache := make(map[int]int)
	for i := min; i <= max; i++ {
		total := 0
		for _, j := range pos {
			diff := int(math.Abs(float64(i - j)))
			if distCache[diff] != 0 {
				total += distCache[diff]
			} else {
				distCost := 0
				for k := 0; k <= diff; k++ {
					distCost += k
				}
				distCache[diff] = distCost
				total += distCache[diff]
			}
		}
		if total < minSum || minSum == -1 {
			minSum = total
		}
	}
	return minSum
}

func main() {
	println(partOne())
	println(partTwo())
}