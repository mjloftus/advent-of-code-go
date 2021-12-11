package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]rune, 0)
		doBreak := 0
		for _, c := range line {
			if strings.Contains("([{<", string(c)) {
				stack = append(stack, c)
			} else if strings.Contains(")]}>", string(c)) {
				switch c {
				case ')':
					if stack[len(stack)-1] == '(' {
						stack = append(stack[:len(stack)-1])
					} else {
						println(line, c)
						score += 3
						doBreak = 1
					}
				case ']':
					if stack[len(stack)-1] == '[' {
						stack = append(stack[:len(stack)-1])
					} else {
						println(line, c)
						score += 57
						doBreak = 1
					}
				case '}':
					if stack[len(stack)-1] == '{' {
						stack = append(stack[:len(stack)-1])
					} else {
						println(line, c)
						score += 1197
						doBreak = 1
					}
				case '>':
					if stack[len(stack)-1] == '<' {
						stack = append(stack[:len(stack)-1])
					} else {
						println(line, c)
						score += 25137
						doBreak = 1
					}
				}
				if doBreak == 1 {
					break
				}
			}
		}
	}
	return score
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scores := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		stack := make([]rune, 0)
		doBreak := 0
		for _, c := range line {
			if strings.Contains("([{<", string(c)) {
				stack = append(stack, c)
			} else if strings.Contains(")]}>", string(c)) {
				switch c {
				case ')':
					if stack[len(stack)-1] == '(' {
						stack = append(stack[:len(stack)-1])
					} else {
						doBreak = 1
					}
				case ']':
					if stack[len(stack)-1] == '[' {
						stack = append(stack[:len(stack)-1])
					} else {
						doBreak = 1
					}
				case '}':
					if stack[len(stack)-1] == '{' {
						stack = append(stack[:len(stack)-1])
					} else {
						doBreak = 1
					}
				case '>':
					if stack[len(stack)-1] == '<' {
						stack = append(stack[:len(stack)-1])
					} else {
						doBreak = 1
					}
				}
				if doBreak == 1 {
					break
				}
			}
		}
		if doBreak == 1 {
			continue
		}
		lineScore := 0
		for i := len(stack)-1; i >= 0; i-- {
			lineScore *= 5
			switch stack[i] {
			case '(':
				lineScore += 1
			case '[':
				lineScore += 2
			case '{':
				lineScore += 3
			case '<':
				lineScore += 4
			}
		}
		scores = append(scores, lineScore)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	println(partOne())
	println(partTwo())
	return
}