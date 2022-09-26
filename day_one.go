package main

func main() {

	// Uncomment once you are ready to test your code.
	// answer := getAnswer("day_one.txt", 1)

	// fmt.Printf("Output: %d\n", answer)
}

func getAnswer(filePath string, part int) int {
	// Open the file. Check https://pkg.go.dev/os 

	// Check for any errors. https://go.dev/blog/error-handling-and-go 

	// We want to close the file at the end of this function. F
	// For that we can use the 'defer' keyword. https://go.dev/tour/flowcontrol/12

	// Read in the contents of the file
	
	// Check for errors again. Should we refactor this code duplication?

	// We need a way of tracking how many blocks we should travel and in which direction. 
	// We are told that the city we are in is based on a grid system and so we can only 
	// turn in 90 degree increments, or put differently, North, South, East, West. Lets
	// create a 2D array to capture the 4 different directions we can travel.
	// https://gobyexample.com/arrays
	// https://gobyexample.com/slices

	// *** COME BACK TO THIS ONCE YOU HAVE COMPLETED PART 1 ***
	// We need a way of tracking where we have already visited. One way of doing this is 
	// to create a map where the key is an array of our current coordinates and the value
	// is a boolean where false indicates we haven't visited the current location before
	// true indicates we have. Create a map with the structure mentioned above. Don't 
	// forget to mark our start position as having already been visited.
	// https://gobyexample.com/maps

	// Create a variable of type int to track of which way we are facing.
	// Create variables to track of our position in the grid.
	// https://gobyexample.com/variables

	// We need to split our instructions up separately and remove any whitespace. 
	// https://pkg.go.dev/strings 
	// Now its time to analyse the instructions. For this we will use a for loop. 
	// https://gobyexample.com/for & https://gobyexample.com/range
	
	// Iterate over every movement instruction
	// Retrieve the first character of a single instruction
	// Retrieve the number of blocks of a single instruction numSteps, 

	// Use if/else block to change the direction we are facing. 
	// Dont forget to consider what happens if we turn >= 360 degrees

	// We are now facing the correct direction. We will use another for loop to move
	// one block at a time and track where we are in on the grid.

	// *** COME BACK TO THIS ONCE YOU HAVE COMPLETED PART 1 ***	
	// Check if we have visited our current location before
	// If we haven't visited before, update the map

	// Finally calculate the total distance in blocks and return
	return 0
}
