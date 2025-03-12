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
	Scores      map[string]float64
	Preferences []string
	FullName    string
}

func (a *Applicant) GetScore(dept string) float64 {
	return a.Scores[departmentSubjects[dept]]
}

var departmentSubjects = map[string]string{
	"Biotech":     "Chemistry",
	"Chemistry":   "Chemistry",
	"Engineering": "Computer Science",
	"Mathematics": "Math",
	"Physics":     "Physics",
}

func main() {
	var N int
	fmt.Scan(&N)

	applicants, err := readApplicants("applicants.txt")
	if err != nil {
		fmt.Println("Error reading applicants:", err)
	}

	deptAdmissions := processAdmissions(applicants, N)
	sortAdmissions(deptAdmissions)
	printResults(deptAdmissions)
}

func readApplicants(filename string) ([]Applicant, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var applicants []Applicant

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		scores := map[string]float64{
			"Physics":          parseScore(fields[2]),
			"Chemistry":        parseScore(fields[3]),
			"Math":             parseScore(fields[4]),
			"Computer Science": parseScore(fields[5]),
		}
		preferences := fields[6:]

		// Store applicant details
		applicants = append(applicants, Applicant{
			FirstName:   fields[0],
			LastName:    fields[1],
			Scores:      scores,
			Preferences: preferences,
			FullName:    fields[0] + " " + fields[1],
		})
	}

	return applicants, scanner.Err()
}

func processAdmissions(applicants []Applicant, N int) map[string][]Applicant {
	deptAdmissions := make(map[string][]Applicant)

	for _, prefLevel := range []int{0, 1, 2} {
		sortApplicants(&applicants, prefLevel)
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
	return deptAdmissions
}

func sortApplicants(applicants *[]Applicant, prefLevel int) {
	sort.SliceStable(*applicants, func(i, j int) bool {
		deptI := (*applicants)[i].Preferences[prefLevel]
		deptJ := (*applicants)[j].Preferences[prefLevel]
		scoreI := (*applicants)[i].GetScore(deptI)
		scoreJ := (*applicants)[j].GetScore(deptJ)
		if scoreI == scoreJ {
			return (*applicants)[i].FullName < (*applicants)[j].FullName
		}
		return scoreI > scoreJ
	})
}

func sortAdmissions(deptAdmissions map[string][]Applicant) {
	for dept, applicants := range deptAdmissions {
		sort.SliceStable(applicants, func(i, j int) bool {
			if applicants[i].GetScore(dept) == applicants[j].GetScore(dept) {
				return applicants[i].FullName < applicants[j].FullName
			}
			return applicants[i].GetScore(dept) > applicants[j].GetScore(dept)
		})
	}
}

func printResults(deptAdmissions map[string][]Applicant) {
	for _, dept := range []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"} {
		fmt.Println(dept)
		for _, applicant := range deptAdmissions[dept] {
			fmt.Printf("%s %.1f\n", applicant.FullName, applicant.GetScore(dept))
		}
		fmt.Println()
	}
}

func parseScore(s string) float64 {
	score, _ := strconv.ParseFloat(s, 64)
	return score
}
