package main

import (
	"fmt"
)

func main() {
	var score1, score2, score3 int

	// Read three integer inputs
	fmt.Scan(&score1)
	fmt.Scan(&score2)
	fmt.Scan(&score3)

	// Calculate the mean score
	meanScore := float64(score1+score2+score3) / 3

	// Print the mean score
	fmt.Println(meanScore)

	// Conditional check for acceptance or rejection
	if meanScore >= 60.0 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}
