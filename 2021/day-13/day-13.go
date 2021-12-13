package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	x int
	y int
}

type Inst struct {
	axis string
	val int
}

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := make([]Pair, 0)
	inst := make([]Inst, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			line := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			points = append(points, Pair{x, y})
		} else if strings.Contains(line, "fold along") {
			line := strings.Split(strings.Split(scanner.Text(), " ")[2], "=")
			v, _ := strconv.Atoi(line[1])
			inst = append(inst, Inst{line[0], v})
		} else {
			continue
		}
	}
	fold := inst[0]
	for i, point := range points {
		if fold.axis == "x" {
			if fold.val < point.x {
				points[i].x = fold.val - (point.x - fold.val)
			}
		} else if fold.axis == "y" {
			if fold.val < point.y {
				points[i].y = fold.val - (point.y - fold.val)
			}
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x < points[j].x {
			return true
		}
		if points[i].x > points[j].x {
			return false
		}
		return points[i].y < points[j].y
	})
	count := 1
	ind := 1
	for ind < len(points) {
		if points[ind].x != points[ind-1].x || points[ind].y != points[ind-1].y {
			count++
		}
		ind++
	}
	return count
}

func partTwo() {

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := make([]Pair, 0)
	inst := make([]Inst, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			line := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			points = append(points, Pair{x, y})
		} else if strings.Contains(line, "fold along") {
			line := strings.Split(strings.Split(scanner.Text(), " ")[2], "=")
			v, _ := strconv.Atoi(line[1])
			inst = append(inst, Inst{line[0], v})
		} else {
			continue
		}
	}
	maxX := 0
	maxY := 0
	for j, fold := range inst {
		for i, point := range points {
			if fold.axis == "x" {
				if fold.val < point.x {
					points[i].x = fold.val - (point.x - fold.val)
				}
			} else if fold.axis == "y" {
				if fold.val < point.y {
					points[i].y = fold.val - (point.y - fold.val)
				}
			}
			if j == len(inst)-1 {
				if point.x > maxX {
					maxX = point.x
				}
				if point.y > maxY {
					maxY = point.y
				}
			}
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x < points[j].x {
			return true
		}
		if points[i].x > points[j].x {
			return false
		}
		return points[i].y < points[j].y
	})
	board := make([][]rune, 0)
	for i := 0; i <= maxY; i++ {
		row := make([]rune, 0)
		for j := 0; j <= maxX; j++ {
			row = append(row, '.')
		}
		board = append(board, row)
	}
	for _, point := range points {
		board[point.y][point.x] = '#'
	}
	for _, row := range board {
		for _, c := range row {
			print(string(c))
		}
		println()
	}
}

func main() {
	println(partOne())
	partTwo()
	return
}
