package main

import (
	"bufio"
	"os"
	"strconv"
)

type Pair struct {
	x int
	y int
}

func addNeighbors(lr int, lc int, i int, j int) []Pair {
	s := make([]Pair, 0)
	if i > 0 {
		s = append(s, Pair{i-1, j})
		if j > 0 {
			s = append(s, Pair{i-1, j-1})
		}
		if j < lc {
			s = append(s, Pair{i-1, j+1})
		}
	}
	if i < lr {
		s = append(s, Pair{i+1, j})
		if j > 0 {
			s = append(s, Pair{i+1, j-1})
		}
		if j < lc {
			s = append(s, Pair{i+1, j+1})
		}
	}
	if j > 0 {
		s = append(s, Pair{i, j-1})
	}
	if j < lc {
		s = append(s, Pair{i, j+1})
	}
	return s
}

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	octopi := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		rowT := scanner.Text()
		for _, c := range rowT {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		octopi = append(octopi, row)
	}
	flashes := 0
	for step := 0; step < 100; step++ {
		for i, row := range octopi {
			for j, _ := range row {
				octopi[i][j]++
				if octopi[i][j] == 10 {
					flashes++
					//inc all neighbors by 1
					pairs := addNeighbors(len(octopi)-1, len(octopi[0])-1, i, j)
					for len(pairs) > 0 {
						p := pairs[0]
						pairs = append(pairs[1:])
						octopi[p.x][p.y]++
						if octopi[p.x][p.y] == 10 {
							flashes++
							pairs = append(pairs, addNeighbors(len(octopi)-1, len(octopi[0])-1, p.x, p.y)...)
						}
					}
				}
			}
		}
		for i, row := range octopi {
			for j, _ := range row {
				if octopi[i][j] >= 10 {
					octopi[i][j] = 0
				}
			}
		}
	}
	return flashes
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	octopi := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		rowT := scanner.Text()
		for _, c := range rowT {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		octopi = append(octopi, row)
	}
	for step := 1; ; step++ {
		flashes := 0
		for i, row := range octopi {
			for j, _ := range row {
				octopi[i][j]++
				if octopi[i][j] == 10 {
					flashes++
					//inc all neighbors by 1
					pairs := addNeighbors(len(octopi)-1, len(octopi[0])-1, i, j)
					for len(pairs) > 0 {
						p := pairs[0]
						pairs = append(pairs[1:])
						octopi[p.x][p.y]++
						if octopi[p.x][p.y] == 10 {
							flashes++
							pairs = append(pairs, addNeighbors(len(octopi)-1, len(octopi[0])-1, p.x, p.y)...)
						}
					}
				}
			}
		}
		for i, row := range octopi {
			for j, _ := range row {
				if octopi[i][j] >= 10 {
					octopi[i][j] = 0
				}
			}
		}
		if flashes == len(octopi) * len(octopi[0]) {
			return step
		}
	}
	return -1
}

func main() {
	println(partOne())
	println(partTwo())
	return
}