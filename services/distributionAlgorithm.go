package services

import (
	"sort"
	"student-distribution/interfaces"
	"student-distribution/models"
)

type distributionAlgorithm struct {
	Students      models.Students
	Coaches       models.Coaches
	CoachesCount  int
	StudentsCount int
}

func NewDistributionAlgorithm(students models.Students, coaches models.Coaches) interfaces.DistributionInterface {
	sort.Sort(sort.Interface(coaches))
	return &distributionAlgorithm{
		Students:      students,
		Coaches:       coaches,
		CoachesCount:  len(coaches),
		StudentsCount: len(students),
	}
}

//problem 1
func (self *distributionAlgorithm) BasicDistribution() {

	coachIndex := 0
	for _, student := range self.Students {
		coach := self.Coaches[coachIndex]
		self.assignStudentForCoach(student, coach)
		if coachIndex == self.CoachesCount-1 {
			coachIndex = 0
			break
		}

		coachIndex++
	}

}

//problem 2
func (self *distributionAlgorithm) DistributeStudents() {

	coachIndex := 0

	for _, student := range self.Students {
		coach := self.Coaches[coachIndex]
		self.assignStudentForCoach(student, coach)
		if coachIndex == self.CoachesCount-1 {
			coachIndex = 0
			break
		}

		nextCoach := self.Coaches[coachIndex+1]
		if coach.GetStudentsCount() > nextCoach.GetStudentsCount() {
			coachIndex++
		} else if coachIndex != 0 {
			coachIndex--
		}

	}
}

func (self *distributionAlgorithm) assignStudentForCoach(student *models.Student, coach *models.Coach) {
	coach.AddStudent(student)
	student.AppendCoach(coach)
}
