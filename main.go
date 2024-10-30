package main

import "fmt"

// CountingLetter counts the unreadable letters in the given text.
func CountingLetter(text string) int {
	// Define the unreadable letters
	unreadableLetters := "rstzRSTZ"
	count := 0
	
	// Loop through each character in the text
	for _, char := range text {
		for _, unreadable := range unreadableLetters {
			if char == unreadable {
				count++
				break // Stop checking once a match is found
			}
		}
	}
	return count
}

// Main function for testing
func main() {
	fmt.Println(CountingLetter("Remaja muda yang berbakat")) // Expected output: 3 (for 'R', 'r', and 't')
	fmt.Println(CountingLetter("Zebra Zig Zag")) // Expected output: 4 (for 'Z', 'Z', 'Z', and 'r')
}
