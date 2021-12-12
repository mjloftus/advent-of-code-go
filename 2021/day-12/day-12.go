package main

import (
	"bufio"
	"os"
	"strings"
)

func search(caveMap map[string][]string, dfs []string, val int) int {
	count := 0
	// go to val'th connection from curPlace in dfs stack
	nextCave := caveMap[dfs[len(dfs)-1]][val]
	if nextCave == "start" {
		count += 0
	} else if nextCave == "end" {
		count += 1
	} else if nextCave[0] < 97 { // can always go to uppercase cave
		newDfs := append(dfs, nextCave)
		count += search(caveMap, newDfs, 0)
	} else { // lowercase cave
		for i, cave := range dfs {
			if cave == nextCave {
				break // skip because already visited lowercase
			} else if i == len(dfs)-1 {
				newDfs := append(dfs, nextCave)
				count += search(caveMap, newDfs, 0)
			}
		}
	}

	if val < len(caveMap[dfs[len(dfs)-1]])-1 {
		count += search(caveMap, dfs, val+1)
	} else {
		return count
	}
	return count
}

func partOne() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	caveToCave := make(map[string][]string)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "-")
		caveToCave[t[0]] = append(caveToCave[t[0]], t[1])
		caveToCave[t[1]] = append(caveToCave[t[1]], t[0])
	}
	dfs := make([]string, 0)
	dfs = append(dfs, "start")
	return search(caveToCave, dfs, 0)
}

func searchTwo(caveMap map[string][]string, dfs []string, val int, visitCount map[string]int, visitedTwice bool) int {
	count := 0
	// go to val'th connection from curPlace in dfs stack
	nextCave := caveMap[dfs[len(dfs)-1]][val]
	if nextCave == "start" {
		count += 0
	} else if nextCave == "end" {
		count += 1
	} else if nextCave[0] < 97 { // can always go to uppercase cave
		newDfs := append(dfs, nextCave)
		count += searchTwo(caveMap, newDfs, 0, visitCount, visitedTwice)
	} else { // lowercase cave
		if visitCount[nextCave] == 0 || (!visitedTwice && visitCount[nextCave] == 1) {
			newVisitCount := make(map[string]int)
			for k, v := range visitCount {
				newVisitCount[k] = v
			}
			newVisitCount[nextCave]++

			newDfs := append(dfs, nextCave)
			count += searchTwo(caveMap, newDfs, 0, newVisitCount, visitedTwice || newVisitCount[nextCave] == 2)
		}
	}

	if val < len(caveMap[dfs[len(dfs)-1]])-1 {
		count += searchTwo(caveMap, dfs, val+1, visitCount, visitedTwice)
	} else {
		return count
	}
	return count
}

func partTwo() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	caveToCave := make(map[string][]string)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), "-")
		caveToCave[t[0]] = append(caveToCave[t[0]], t[1])
		caveToCave[t[1]] = append(caveToCave[t[1]], t[0])
	}
	dfs := make([]string, 0)
	dfs = append(dfs, "start")
	visitCount := make(map[string]int)
	return searchTwo(caveToCave, dfs, 0, visitCount, false)
}

func main() {
	println(partOne())
	println(partTwo())
	return
}