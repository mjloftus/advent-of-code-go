package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func partOne(days int) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lanternFish := make([]int, 0)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	for _, val := range vals {
		v, _ := strconv.Atoi(val)
		lanternFish = append(lanternFish, v)
	}
	for i := 0; i < days; i++ {
		newFish := 0
		for ind, _ := range lanternFish {
			lanternFish[ind]--
			if lanternFish[ind] == -1 {
				newFish++
				lanternFish[ind] = 6
			}
		}
		for j := 0; j < newFish; j++ {
			lanternFish = append(lanternFish, 8)
		}
	}
	return len(lanternFish)
}

func partTwo(days int) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lanternFish := make([]int, 0)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	for _, val := range vals {
		v, _ := strconv.Atoi(val)
		lanternFish = append(lanternFish, v)
	}
	m := make(map[int]int)
	for i := days-1; i > days-9; i-- {
		m[i] = 0
	}
	for i := days-9; i > 0; i-- {
		m[i] = 1 + (days - (i + 9)) / 7
		for j := i + 9; j < days; j += 7 {
			m[i] += m[j]
		}
	}
	total := len(lanternFish)
	for _, fish := range lanternFish {
		total += 1 + (days - fish - 1) / 7
		for i := fish+1; i < days; i += 7 {
			total += m[i]
		}
	}
	return total
}

func main() {
	println(partOne(80))
	println(partTwo(256))
	return
}