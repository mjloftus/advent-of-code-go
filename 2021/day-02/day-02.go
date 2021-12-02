package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partOne() int {
	input, _ := ioutil.ReadFile("input.txt")
	commands := strings.Split(string(input), "\n")

	xPos := 0
	yPos := 0
	for _, command := range commands {
		command := strings.Split(command, " ")
		if len(command) != 2 {
			break
		}
		val, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			xPos += val
		case "up":
			yPos -= val
		case "down":
			yPos += val
		}
	}
	return xPos * yPos
}

func partTwo() int {
	input, _ := ioutil.ReadFile("input.txt")
	commands := strings.Split(string(input), "\n")

	xPos := 0
	yPos := 0
	aim := 0
	for _, command := range commands {
		command := strings.Split(command, " ")
		if len(command) != 2 {
			break
		}
		val, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			xPos += val
			yPos += val * aim
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}
	return xPos * yPos
}

func main() {
	fmt.Printf("part 1: %d\n", partOne())
	fmt.Printf("part 2: %d", partTwo())
	return
}