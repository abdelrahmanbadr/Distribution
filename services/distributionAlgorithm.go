package services

import (
	"fmt"
	"student-distribution/models"
)

type DistributionAlgorithm struct {
	Students []*models.Student
	Coaches  []*models.Coach
}

func (self *DistributionAlgorithm) BasicDistribution() {

	studentsPerCoach := len(self.Students) / len(self.Coaches)
	coachIndex := 0
	for key, student := range self.Students {

		if key > 0 && key%studentsPerCoach == 0 && coachIndex < len(self.Coaches)-1 {

			coachIndex++;
			fmt.Println(coachIndex)
		}
		student.Coaches = append(student.Coaches, self.Coaches[coachIndex])
		self.Coaches[coachIndex].Students = append(self.Coaches[coachIndex].Students, student)

	}

}
