package services

import (
	"math"
	"sort"
	"student-distribution/models"
)

type DistributionAlgorithm struct {
	Students models.Students
	Coaches  models.Coaches
}

//problem 1
func (self *DistributionAlgorithm) BasicDistribution() {

	sudentsCount := len(self.Students)
	studentsPerCoach := sudentsCount / len(self.Coaches)
	studentIndex := 0

	for _, coach := range self.Coaches {

		for i := 0; i < studentsPerCoach; i++ {
			student := self.Students[studentIndex]
			coach.Students = append(coach.Students, student)
			student.Coaches = append(student.Coaches, coach)
			studentIndex++
		}
	}
	remainStudents := sudentsCount - studentIndex
	//if remainStudents > 0
	for i := 0; i < remainStudents; i++ {
		coach := self.Coaches[i]
		student := self.Students[studentIndex]
		coach.Students = append(coach.Students, student)
		student.Coaches = append(student.Coaches, coach)
		studentIndex++
	}
}

//problem 2
func (self *DistributionAlgorithm) FairDistribution() {
	coachesNumber := len(self.Coaches)
	studentsNumber := len(self.Students)

	sort.Sort(sort.Interface(self.Coaches))
	for _, coach := range self.Coaches {
		studentsNumber += coach.GetStudentsCount()
	}

	division := Divide(studentsNumber, coachesNumber)

	counter := 0

	for _, coach := range self.Coaches {

		if coach.GetStudentsCount() >= division {
			studentsNumber -= coach.GetStudentsCount()
			continue
		}
		counter++
	}
	//remain studentsNumber
	division = Divide(studentsNumber, counter)
	studentsIndex := 0
	for j := 0; j < counter; j++ {
		studentsNumberForCoach := division - self.Coaches[j].GetStudentsCount()

		//last key will take the change
		if j == counter-1 {
			studentsNumberForCoach = len(self.Students) - studentsIndex
		}

		for i := 0; i < studentsNumberForCoach; i++ {
			self.Coaches[j].Students = append(self.Coaches[j].Students, self.Students[studentsIndex])
			studentsIndex++
		}
	}

}

func Divide(num1, num2 int) int {

	return int(math.Round(float64(num1) / float64(num2)))
}
