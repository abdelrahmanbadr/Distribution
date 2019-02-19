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
	studentsPerCoach := sudentsCount / len(self.Coaches) // use RoundDivide
	studentsIndex := 0

	for _, coach := range self.Coaches {

		for i := 0; i < studentsPerCoach; i++ {
			student := self.Students[studentsIndex]
			coach.Students = append(coach.Students, student)
			student.Coaches = append(student.Coaches, coach)
			studentsIndex++
		}
	}
	remainStudents := sudentsCount - studentsIndex
	//if remainStudents > 0
	for i := 0; i < remainStudents; i++ {
		coach := self.Coaches[i]
		self.AssignStudentForCoach(self.Students[studentsIndex], coach)
		studentsIndex++
	}
}

//problem 2
func (self *DistributionAlgorithm) FairDistribution() {
	sort.Sort(sort.Interface(self.Coaches))
	coachesNumber := len(self.Coaches)
	studentsCount := self.CountCoachesStudents() + len(self.Students)

	studentsPerCoach := RoundDivide(studentsCount, coachesNumber)

	//counter for coaches
	unSaturatedCoachesCount := 0

	for _, coach := range self.Coaches {

		if coach.GetStudentsCount() >= studentsPerCoach {
			studentsCount -= coach.GetStudentsCount()
			continue
		}
		unSaturatedCoachesCount++
	}

	studentsPerCoach = RoundDivide(studentsCount, unSaturatedCoachesCount)
	studentsIndex := 0
	for j := 0; j < unSaturatedCoachesCount; j++ {
		coach := self.Coaches[j]
		studentsCountPerCoach := studentsPerCoach - coach.GetStudentsCount()

		//last key will take the change
		if j == unSaturatedCoachesCount-1 {
			studentsCountPerCoach = len(self.Students) - studentsIndex
		}

		for i := 0; i < studentsCountPerCoach; i++ {
			self.AssignStudentForCoach(self.Students[studentsIndex], coach)
			studentsIndex++
		}
	}

}

func RoundDivide(num1, num2 int) int {

	return int(math.Round(float64(num1) / float64(num2)))
}
func (self *DistributionAlgorithm) AssignStudentForCoach(student *models.Student, coach *models.Coach) {
	coach.AddStudent(student)
	student.AppendCoach(coach)
}

func (self *DistributionAlgorithm) CountCoachesStudents() int {
	studentCount := 0
	for _, coach := range self.Coaches {
		studentCount += coach.GetStudentsCount()
	}
	return studentCount
}
