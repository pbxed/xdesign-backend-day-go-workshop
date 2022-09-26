package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := getAnswer("xdesign_day_one.txt", 2)

	fmt.Printf("Output: %d\n", answer)
}

func getAnswer(filePath string, part int) int {
	file, err := os.Open(filePath) // Open the file

	check(err) // check for errors opening the file

	defer file.Close() // Defer delays the execution of this line until we get to the end of this function
	// file.Close() closes the file so it can not be used by IO anymore.

	buf, err := os.ReadFile(filePath) // read in file
	check(err)                        // check for errors

	var directions = [][2]int{
		{-1, 0}, // North
		{0, 1},  // East
		{1, 0},  // South
		{0, -1}, // West
	}

	visited := map[[2]int]bool{ // A Map of Key: 2D slice containing slices of 2 digit ints with Value: boolean flag
		{0, 0}: true, // We mark the starting position as having already been visited
	}

	var directionIndex int // Keep track of which way we are facing
	var x, y int           // Keep track of our position in the grid

	for _, instruction := range strings.Split(strings.TrimSpace(string(buf)), ", ") { // Iterate over every movement instruction
		turn := string(instruction[0])               // Takes the first char of each instruction
		numSteps, _ := strconv.Atoi(instruction[1:]) // Takes the number of numSteps of each instruction

		if turn == "R" {
			directionIndex = (directionIndex + 1) % 4 // Modulo 4 as we only have 4 possible directions we can be facing
		} else if turn == "L" {
			directionIndex = (directionIndex + 3) % 4 // +3 instead of -1 so that we dont have to deal with negative numbers
		} else {
			panic("direction instruction other than Right or Left found: " + turn) // Throw error if our instructions contain something other than R or L turn instruction
		}

		for i := 0; i < numSteps; i++ { // Move one step at a time
			x += directions[directionIndex][0] // track where we are in on the grid
			y += directions[directionIndex][1]
			// check if we have visited before
			if visited[[2]int{x, y}] && part == 2 {
				return distance(x, y)
			}
			// if not visited before mark visited
			visited[[2]int{x, y}] = true
		}
	}
	return distance(x, y)
}

func distance(x, y int) int {
	// Assume starting position is (0,0)

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	// Distance given by |x|+|y|
	return x + y
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
