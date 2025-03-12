package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Applicant struct to store applicant details
type Applicant struct {
	FirstName   string
	LastName    string
	GPA         float64
	Preferences []string
	FullName    string
}

func main() {
	var N int

	// Read the number of accepted applicants
	fmt.Scan(&N)

	file, err := os.Open("applicants.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var applicants []Applicant
	departments := []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}
	deptAdmissions := make(map[string][]Applicant)

	scanner := bufio.NewScanner(file)

	// Read applicant data
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		gpa, _ := strconv.ParseFloat(fields[2], 64)
		preferences := fields[3:]

		// Store applicant details
		applicants = append(applicants, Applicant{
			fields[0],
			fields[1],
			gpa,
			preferences,
			fields[0] + " " + fields[1],
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort applicants by GPA (descending) and FullName (ascending if GPA is equal)
	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].GPA == applicants[j].GPA {
			return applicants[i].FullName < applicants[j].FullName
		}
		return applicants[i].GPA > applicants[j].GPA
	})

	// Process applications based on preferences
	for _, prefLevel := range []int{0, 1, 2} {
		for i := 0; i < len(applicants); i++ {
			applicant := &applicants[i]
			if len(applicant.Preferences) > prefLevel {
				dept := applicant.Preferences[prefLevel]
				if len(deptAdmissions[dept]) < N {
					deptAdmissions[dept] = append(deptAdmissions[dept], *applicant)
					applicants = append(applicants[:i], applicants[i+1:]...) // Remove admitted applicant
					i--                                                      // Adjust index after removal
				}
			}
		}
	}

	// Sort admitted applicants within each department
	for _, dept := range departments {
		sort.Slice(deptAdmissions[dept], func(i, j int) bool {
			if deptAdmissions[dept][i].GPA == deptAdmissions[dept][j].GPA {
				return deptAdmissions[dept][i].FullName < deptAdmissions[dept][j].FullName
			}
			return deptAdmissions[dept][i].GPA > deptAdmissions[dept][j].GPA
		})
	}

	// Print the successful applicants
	for _, dept := range departments {
		fmt.Println(dept)
		for _, applicant := range deptAdmissions[dept] {
			fmt.Printf("%s %.2f\n", applicant.FullName, applicant.GPA)
		}
		fmt.Println()
	}
}
