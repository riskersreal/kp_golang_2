package main

import "fmt"

type Student struct {
	name  string
	score float32
}

type sliceStudent struct {
	students []Student
}

func avgCount(total float32, count int) float32 {
	return total / float32(count)
}

func (s *sliceStudent) avgStudents() {
	var totalScore, avgScore float32
	var count int = 0

	for _, p := range s.students {
		count++
		totalScore += p.score
	}
	avgScore = avgCount(totalScore, count)
	fmt.Printf("Average Score: %.2f\n", avgScore)
}

func (s *sliceStudent) minStudents() {
	var minName string = ""
	var minScore float32 = 100.00

	for _, p := range s.students {
		if minScore > p.score {
			minName = p.name
			minScore = p.score
		}
	}
	fmt.Printf("Min Score of Students: %s (%.0f)\n", minName, minScore)
}

func (s *sliceStudent) maxStudents() {
	var maxName string = ""
	var maxScore float32 = 0.00

	for _, p := range s.students {
		if maxScore < p.score {
			maxName = p.name
			maxScore = p.score
		}
	}
	fmt.Printf("Max Score of Students: %s (%.0f)\n", maxName, maxScore)
}

func main() {
	var mainS sliceStudent
	var studentCount int = 5

	fmt.Println("Input")
	for i := 0; i < studentCount; i++ {
		var inputStudent Student
		fmt.Print("Input", i+1, "Student's Name ")
		fmt.Scan(&inputStudent.name)
		fmt.Print("Input", i+1, "Student's Score")
		fmt.Scan(&inputStudent.score)
		mainS.students = append(mainS.students, inputStudent)
	}
	mainS.avgStudents()
	mainS.minStudents()
	mainS.maxStudents()
}
