package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() int64 {
	input, _ := ioutil.ReadFile("input.txt")
	binaries := strings.Split(string(input), "\n")
	var counts []int

	for ind, val := range binaries {
		for ind2, bit := range val {
			if ind == 0 {
				counts = append(counts, 0)
			}
			if bit == '1' {
				counts[ind2] += 1
			}
		}
	}
	var gammaS = ""
	var epsilonS = ""
	for _, val := range counts {
		if val <= len(binaries) / 2 {
			gammaS += "0"
			epsilonS += "1"
		} else {
			gammaS += "1"
			epsilonS += "0"
		}
	}
	gamma, _ := strconv.ParseInt(gammaS, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonS, 2, 64)
	return gamma * epsilon
}

func generateFreqString(binaries []string) string {
	var counts []int
	for ind, val := range binaries {
		for ind2, binary := range val {
			if ind == 0 {
				counts = append(counts, 0)
			}
			if binary == '1' {
				counts[ind2] += 1
			}
		}
	}
	var freqString = ""
	for _, val := range counts {
		if float32(val) < float32(len(binaries)) / 2.0 {
			freqString += "0"
		} else {
			freqString += "1"
		}
	}
	return freqString
}

func partTwo() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("err", err)
	}
	scanner := bufio.NewScanner(file)
	var binaries []string
	for scanner.Scan() {
		binaries = append(binaries, scanner.Text())
	}

	oxygenCands := make([]string, len(binaries))
	coCands := make([]string, len(binaries))
	copy(oxygenCands, binaries)
	copy(coCands, binaries)
	oxygenFreqString := generateFreqString(binaries)
	coFreqString := generateFreqString(binaries)
	for i := 0; i < len(binaries[0]); i++ {
		var newOxygenCands []string
		for _, cand := range oxygenCands {
			if len(oxygenCands) == 1 {
				break
			}
			if cand[i] == oxygenFreqString[i] {
				newOxygenCands = append(newOxygenCands, cand)
			}
		}
		if len(oxygenCands) != 1 {
			oxygenCands = newOxygenCands
			oxygenFreqString = generateFreqString(oxygenCands)
		}
		var newCoCands []string
		for _, cand := range coCands {
			if len(coCands) == 1 {
				break
			}
			if cand[i] != coFreqString[i] {
				newCoCands = append(newCoCands, cand)
			}
		}
		if len(coCands) != 1 {
			coCands = newCoCands
			coFreqString = generateFreqString(coCands)
		}
	}
	oxygenValue, _ := strconv.ParseInt(oxygenCands[0], 2, 32)
	coValue, _ := strconv.ParseInt(coCands[0], 2, 32)
	return int(oxygenValue * coValue)
}

func main() {
	println(partOne())
	println(partTwo())
	return
}