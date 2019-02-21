package services

import (
	"sort"
	"student-distribution/common"
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

	studentsIndex := 0
	for i := 0; i < self.StudentsCount; i++ {
		coachIndex := i % self.CoachesCount
		coach := self.Coaches[coachIndex]
		self.AssignStudentForCoach(self.Students[studentsIndex], coach)
		studentsIndex++
	}
}

//problem 2
func (self *distributionAlgorithm) FairDistribution() {

	//if coaches numbers = 1 or if first coach students + incomming students =< second coach so return
	if self.CoachesCount == 1 || (self.CoachesCount > 2 && (self.Coaches[0].GetStudentsCount()+self.StudentsCount) < self.Coaches[1].GetStudentsCount()) {
		//first coach will take all students
		self.CoachesCount = 1
		self.BasicDistribution()
		return
	}
	studentsCount := self.CountCoachesStudents() + self.StudentsCount

	studentsPerCoach := common.RoundDivide(studentsCount, self.CoachesCount)

	//counter for coaches
	unSaturatedCoachesCount := 0

	for _, coach := range self.Coaches {

		if coach.GetStudentsCount() >= studentsPerCoach {
			studentsCount -= coach.GetStudentsCount()
			continue
		}
		unSaturatedCoachesCount++
	}

	studentsPerCoach = common.RoundDivide(studentsCount, unSaturatedCoachesCount)

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
