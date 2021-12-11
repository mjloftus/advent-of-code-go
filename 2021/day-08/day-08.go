package main

import (
	"bufio"
	"os"
	"strings"
)

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	//outputs := make([]string, 0)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		output := strings.Split(line, "|")[1]
		digits := strings.Split(strings.TrimSpace(output), " ")
		for _, digit := range digits {
			switch len(digit) {
			case 2:
				count += 1
			case 4:
				count += 1
			case 3:
				count += 1
			case 7:
				count += 1
			}
		}
	}
	return count
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		signals := strings.Split(strings.TrimSpace(strings.Split(line, "|")[0]), " ")
		digitToSignal := make(map[int]string)
		for _, signal := range signals {
			curDigit := -1
			switch len(signal) {
			case 2:
				curDigit = 1
			case 4:
				curDigit = 4
			case 3:
				curDigit = 7
			case 7:
				curDigit = 8
			}
			if curDigit != -1 {
				digitToSignal[curDigit] = signal
			}
		}
		for _, signal := range signals {
			curDigit := -1
			switch len(signal) {
			case 5: // 2, 3, 5
				simCount := 0
				if strings.Contains(signal, string(digitToSignal[1][0])) && strings.Contains(signal, string(digitToSignal[1][1])) {
					curDigit = 3
				} else {
					for _, c := range digitToSignal[4] {
						if strings.Contains(signal, string(c)) {
							simCount += 1
						}
					}
					if simCount == 3 {
						curDigit = 5
					} else {
						curDigit = 2
					}
				}
			case 6: // 0, 6, 9
				if strings.Contains(signal, string(digitToSignal[4][0])) && strings.Contains(signal, string(digitToSignal[4][1])) && strings.Contains(signal, string(digitToSignal[4][2])) && strings.Contains(signal, string(digitToSignal[4][3])) {
					curDigit = 9
				} else if strings.Contains(signal, string(digitToSignal[1][0])) && strings.Contains(signal, string(digitToSignal[1][1])) {
					curDigit = 0
				} else {
					curDigit = 6
				}
			}
			if curDigit != -1 {
				digitToSignal[curDigit] = signal
			}
		}
		outputs := strings.Split(strings.TrimSpace(strings.Split(line, "|")[1]), " ")
		val := 0
		mag := 1000
		for _, output := range outputs {
			for k, v := range digitToSignal {
				if len(output) == len(v) {
					for i, j := range v {
						if !strings.Contains(output, string(j)) {
							break
						}
						if i == len(output) - 1 {
							val += k * mag
							mag /= 10
							break
						}
					}
				}
			}
		}
		count += val
	}
	return count
}

func main() {
	println(partOne())
	println(partTwo())
	return
}
