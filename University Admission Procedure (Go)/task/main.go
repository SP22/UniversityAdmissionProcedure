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

	// Print the acceptance message
	fmt.Println("Congratulations, you are accepted!")
}
