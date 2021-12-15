package main

import (
	"bufio"
	"math/big"
	"os"
	"strings"
)

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()
	chains := make(map[string]string)
	for scanner.Scan() {
		chain := strings.Split(scanner.Text(), " -> ")
		chains[chain[0]] = chain[1]
	}
	for i := 0; i < 10; i++ {
		newPolymer := ""
		for j, c := range polymer {
			newPolymer += string(c)
			if j < len(polymer)-1 {
				newPolymer += chains[string(polymer[j])+string(polymer[j+1])]
			}
		}
		polymer = newPolymer
	}
	count := make(map[rune]int)
	for _, c := range polymer {
		count[c]++
	}
	min := -1
	max := -1
	for _, v := range count {
		if max == -1 || v > max {
			max = v
		}
		if min == -1 || v < min {
			min = v
		}
	}
	return max - min
}

func partTwo() *big.Int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()
	chains := make(map[string]string)
	for scanner.Scan() {
		chain := strings.Split(scanner.Text(), " -> ")
		chains[chain[0]] = chain[1]
	}
	pairCount := make(map[string]big.Int)
	for i , _ := range polymer {
		if i == len(polymer)-1 {
			break
		}
		curVal := pairCount[string(polymer[i]) + string(polymer[i+1])]
		curVal.Add(&curVal, big.NewInt(1))
		pairCount[string(polymer[i]) + string(polymer[i+1])] = curVal
	}
	for i := 0; i < 40; i++ {
		newPairCount := make(map[string]big.Int)
		for k, v := range pairCount {
			curVal1 := newPairCount[string(k[0]) + chains[k]]
			curVal1.Add(&curVal1, &v)
			newPairCount[string(k[0]) + chains[k]] = curVal1
			curVal2 := newPairCount[chains[k] + string(k[1])]
			curVal2.Add(&curVal2, &v)
			newPairCount[chains[k] + string(k[1])] = curVal2
		}
		pairCount = newPairCount
	}
	count := make(map[rune]big.Int)
	for k, v := range pairCount {
		curVal1 := count[rune(k[0])]
		curVal1.Add(&curVal1, &v)
		count[rune(k[0])] = curVal1
		curVal2 := count[rune(k[1])]
		curVal2.Add(&curVal2, &v)
		count[rune(k[1])] = curVal2
	}
	for k, _ := range count {
		curVal := count[k]
		curVal.Div(&curVal, big.NewInt(2))
		count[k] = curVal
	}
	curVal := count[rune(polymer[0])]
	curVal.Add(&curVal, big.NewInt(1))
	count[rune(polymer[0])] = curVal
	curVal2 := count[rune(polymer[len(polymer)-1])]
	curVal2.Add(&curVal2, big.NewInt(1))
	count[rune(polymer[len(polymer)-1])] = curVal2
	if polymer[0] == polymer[len(polymer)-1] {
		curVal := count[rune(polymer[0])]
		curVal.Sub(&curVal, big.NewInt(1))
		count[rune(polymer[0])] = curVal
	}
	min := big.NewInt(-1)
	max := big.NewInt(-1)
	for _, v := range count {
		if max.Cmp(big.NewInt(-1)) == 0 || max.Cmp(&v) == -1 {
			max.Set(&v)
		}
		if min.Cmp(big.NewInt(-1)) == 0 || min.Cmp(&v) == 1 {
			min.Set(&v)
		}
	}
	res := big.NewInt(0)
	res.Sub(max, min)
	return res
}

func main() {
	println(partOne())
	println(partTwo().Text(10))
	return
}