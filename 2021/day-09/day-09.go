package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var vals [][]int
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, n := range line {
			val, _ := strconv.Atoi(string(n))
			row = append(row, val)
		}
		vals = append(vals, row)
	}
	risk := 0
	for i, row := range vals {
		for j, n := range row {
			if i != 0 && vals[i-1][j] <= vals[i][j] {
				continue
			}
			if i <= len(vals)-2 && vals[i+1][j] <= vals[i][j] {
				continue
			}
			if j != 0 && vals[i][j-1] <= vals[i][j] {
				continue
			}
			if j <= len(vals[0])-2 && vals[i][j+1] <= vals[i][j] {
				continue
			}
			risk += 1 + n
		}
	}
	return risk
}

type Pair struct {
	a int
	b int
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var vals [][]int
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, n := range line {
			val, _ := strconv.Atoi(string(n))
			row = append(row, val)
		}
		vals = append(vals, row)
	}
	var basinSizes []int
	for i, row := range vals {
		for j, _ := range row {
			if vals[i][j] == -1 {
				continue
			}
			if i != 0 && vals[i-1][j] <= vals[i][j] {
				continue
			}
			if i <= len(vals)-2 && vals[i+1][j] <= vals[i][j] {
				continue
			}
			if j != 0 && vals[i][j-1] <= vals[i][j] {
				continue
			}
			if j <= len(vals[0])-2 && vals[i][j+1] <= vals[i][j] {
				continue
			}
			basinSize := 0
			var queue []Pair
			queue = append(queue, Pair{i, j})
			for len(queue) > 0 {
				a := queue[0].a
				b := queue[0].b
				queue = queue[1:]
				if vals[a][b] == -1 {
					continue
				}
				basinSize++
				vals[a][b] = -1
				if a != 0 && vals[a-1][b] != 9 {
					queue = append(queue, Pair{a-1, b})
				}
				if a <= len(vals)-2 && vals[a+1][b] != 9 {
					queue = append(queue, Pair{a+1, b})
				}
				if b != 0 && vals[a][b-1] != 9 {
					queue = append(queue, Pair{a, b-1})
				}
				if b <= len(vals[0])-2 && vals[a][b+1] != 9 {
					queue = append(queue, Pair{a, b+1})
				}
			}
			basinSizes = append(basinSizes, basinSize)
		}
	}
	sort.Ints(basinSizes)
	l := len(basinSizes)
	res := basinSizes[l-1] * basinSizes[l-2] * basinSizes[l-3]
	return res
}

func main() {
	println(partOne())
	println(partTwo())
	return
}
