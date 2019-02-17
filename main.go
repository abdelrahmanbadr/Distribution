package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"student-distribution/models"
	"student-distribution/services"
)

func main() {
	DistributionAlgorithm := services.DistributionAlgorithm{}
	DistributionAlgorithm.Coaches = fakeCoaches(3)
	DistributionAlgorithm.Students = fakeStudents(9)

	DistributionAlgorithm.BasicDistribution()
	fmt.Println(DistributionAlgorithm.Coaches[0].Students)
}
func fakeStudents(studentNumber int) []*models.Student {
	students := []*models.Student{}
	for i := 1; i <= studentNumber; i++ {
		students = append(students, &models.Student{Name: fake.FullName()})
	}
	return students
}
func fakeCoaches(coachesNumber int) []*models.Coach {
	coaches := []*models.Coach{}
	for i := 1; i <= coachesNumber; i++ {
		coaches = append(coaches, &models.Coach{Name: fake.FullName()})
	}
	return coaches
}
