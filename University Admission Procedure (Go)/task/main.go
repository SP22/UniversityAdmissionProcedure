package main

import (
	"fmt"
	"sort"
)

// Applicant struct to store applicant details
type Applicant struct {
	FirstName string
	LastName  string
	GPA       float64
	FullName  string
}

func main() {
	var N, M int

	// Read the number of applicants and number of accepted applicants
	fmt.Scan(&N)
	fmt.Scan(&M)

	applicants := make([]Applicant, N)

	// Read N lines of applicant data
	for i := 0; i < N; i++ {
		var firstName, lastName string
		var gpa float64
		fmt.Scan(&firstName, &lastName, &gpa)

		// Store applicant details
		applicants[i] = Applicant{
			firstName,
			lastName,
			gpa,
			firstName + " " + lastName,
		}
	}

	// Sort applicants by GPA in descending order
	// If GPA is the same, sort by full name in ascending order
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].GPA == applicants[j].GPA {
			return applicants[i].FullName < applicants[j].FullName
		}
		return applicants[i].GPA > applicants[j].GPA
	})

	// Print the successful applicants
	fmt.Println("Successful applicants:")
	for i := 0; i < M; i++ {
		fmt.Println(applicants[i].FullName)
	}
}
