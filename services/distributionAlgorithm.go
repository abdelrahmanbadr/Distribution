package services

import (
	"math"
	"sort"
	"student-distribution/models"
)

type distributionAlgorithm struct {
	Students      models.Students
	Coaches       models.Coaches
	CoachesCount  int
	StudentsCount int
}

func NewDistributionAlgorithm(students models.Students, coaches models.Coaches) *distributionAlgorithm {
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

	studentsPerCoach := RoundDivide(self.StudentsCount, self.CoachesCount)
	studentsIndex := 0

	for _, coach := range self.Coaches {

		for i := 0; i < studentsPerCoach; i++ {
			student := self.Students[studentsIndex]
			coach.Students = append(coach.Students, student)
			student.Coaches = append(student.Coaches, coach)
			studentsIndex++
		}
	}
	remainStudents := self.StudentsCount - studentsIndex
	//if remainStudents > 0
	for i := 0; i < remainStudents; i++ {
		coach := self.Coaches[i]
		self.AssignStudentForCoach(self.Students[studentsIndex], coach)
		studentsIndex++
	}
}

//problem 2
func (self *distributionAlgorithm) FairDistribution() {

	//if coaches numbers = 1
	//or if first coach students + incomming students =< second coach so return
	if self.CoachesCount == 1 || (self.CoachesCount > 2 && (self.Coaches[0].GetStudentsCount()+self.StudentsCount) < self.Coaches[1].GetStudentsCount()) {
		//first coach will take all students

		self.AssignAllStudentsForFirstCoach()

		return
	}
	studentsCount := self.CountCoachesStudents() + self.StudentsCount

	studentsPerCoach := RoundDivide(studentsCount, self.CoachesCount)

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
func FloorDivide(num1, num2 int) int {

	return int(math.Floor(float64(num1) / float64(num2)))
}
func (self *distributionAlgorithm) AssignStudentForCoach(student *models.Student, coach *models.Coach) {
	coach.AddStudent(student)
	student.AppendCoach(coach)
}

func (self *distributionAlgorithm) CountCoachesStudents() int {
	studentCount := 0
	for _, coach := range self.Coaches {
		studentCount += coach.GetStudentsCount()
	}
	return studentCount
}

func (self *distributionAlgorithm) AssignAllStudentsForFirstCoach() {

	coach := self.Coaches[0]
	for _, student := range self.Students {
		coach.Students = append(coach.Students, student)
		student.Coaches = append(student.Coaches, coach)
	}

}
