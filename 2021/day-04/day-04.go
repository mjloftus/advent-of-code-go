package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partOne() int64 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bingoNums := strings.Split(scanner.Text(), ",")
	var bingoBoards [][][]int64
	rows := make([][]int64, 5)
	cols := make([][]int64, 5)
	// diags := make([][]int64, 2)
	bingoBoard := make([][]int64, 0)
	var row = 0
	scanner.Scan()
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		for ind, val := range vals {
			if val == "" {
				vals = append(vals[:ind], vals[ind+1:]...)
			}
		}
		if len(vals) == 0 {
			for _, i := range rows {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			}
			for _, i := range cols {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			}
			/* for _, i := range diags {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			} */
			bingoBoards = append(bingoBoards, bingoBoard)
			rows = make([][]int64, 5)
			cols = make([][]int64, 5)
			bingoBoard = make([][]int64, 0)
			//diags = make([][]int64, 2)
			row = 0
			continue
		}
		for col, val := range vals {
			num, _ := strconv.ParseInt(val, 10, 32)
			rows[row] = append(rows[row], num)
			cols[col] = append(cols[col], num)
			/* if row == col {
				diags[0] = append(diags[0], num)
			}
			if row + col == 4 {
				diags[1] = append(diags[1], num)
			} */
		}
		row = row + 1
	}

	for _, bingoNum := range bingoNums {
		num, _ := strconv.ParseInt(bingoNum, 10, 32)
		println(num)
		for _, board := range bingoBoards {
			for _, group := range board {
				for ind, val := range group {
					if val == num {
						group[ind] = -1
					}
				}
			}
			for _, group := range board {
				for ind, val := range group {
					if val != -1 {
						break
					}
					if ind == 4 {
						m := make(map[int64]bool)
						m[-1] = true
						var sum int64 = 0
						println("sum")
						for _, win := range board {
							for _, winV := range win {
								print(winV, " ")
								if !m[winV] {
									m[winV] = true
									println(winV)
									sum = sum + winV
								}
							}
							println()
						}
						println(sum)
						println(num)
						return sum * num
					}
				}
			}
		}
	}
	return 5
}

func partTwo() int64 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bingoNums := strings.Split(scanner.Text(), ",")
	var bingoBoards [][][]int64
	rows := make([][]int64, 5)
	cols := make([][]int64, 5)
	// diags := make([][]int64, 2)
	bingoBoard := make([][]int64, 0)
	var row = 0
	scanner.Scan()
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		for ind, val := range vals {
			if val == "" {
				vals = append(vals[:ind], vals[ind+1:]...)
			}
		}
		if len(vals) == 0 {
			for _, i := range rows {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			}
			for _, i := range cols {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			}
			/* for _, i := range diags {
				for _, j := range i {
					print(j, " ")
				}
				println()
				bingoBoard = append(bingoBoard, i)
			} */
			bingoBoards = append(bingoBoards, bingoBoard)
			rows = make([][]int64, 5)
			cols = make([][]int64, 5)
			bingoBoard = make([][]int64, 0)
			//diags = make([][]int64, 2)
			row = 0
			continue
		}
		for col, val := range vals {
			num, _ := strconv.ParseInt(val, 10, 32)
			rows[row] = append(rows[row], num)
			cols[col] = append(cols[col], num)
			/* if row == col {
				diags[0] = append(diags[0], num)
			}
			if row + col == 4 {
				diags[1] = append(diags[1], num)
			} */
		}
		row = row + 1
	}

	for _, bingoNum := range bingoNums {
		num, _ := strconv.ParseInt(bingoNum, 10, 32)
		println(num)
		var boardsToRemove []int
		for boardInd, board := range bingoBoards {
			for _, group := range board {
				for ind, val := range group {
					if val == num {
						group[ind] = -1
					}
				}
			}
			for _, group := range board {
				for ind, val := range group {
					if val != -1 {
						break
					}
					if ind == 4 {
						if len(bingoBoards) == 1 {
							m := make(map[int64]bool)
							m[-1] = true
							var sum int64 = 0
							println("sum")
							for _, win := range board {
								for _, winV := range win {
									print(winV, " ")
									if !m[winV] {
										m[winV] = true
										println(winV)
										sum = sum + winV
									}
								}
								println()
							}
							println(sum)
							println(num)
							return sum * num
						} else {
							boardsToRemove = append(boardsToRemove, boardInd)
						}
					}
				}
			}
		}
		sort.Slice(boardsToRemove, func(i, j int) bool {
			return boardsToRemove[i] > boardsToRemove[j]
		})
		for _, i := range boardsToRemove {
			println("good")
			bingoBoards = append(bingoBoards[:i], bingoBoards[i+1:]...)
		}
	}
	return 5
}

func main() {
	println(partOne())
	println(partTwo())
	return
}