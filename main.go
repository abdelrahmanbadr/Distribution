package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"student-distribution/models"
	"student-distribution/services"
)

func main() {
	//6,7,8,12
	coaches := fakeCoaches(3)
	students := fakeStudents(6)

	distributionAlgorithm := services.NewDistributionAlgorithm(students, coaches)
	distributionAlgorithm.FairDistribution()

	fmt.Println(coaches[0].GetStudentsCount())
	fmt.Println(coaches[1].GetStudentsCount())
	fmt.Println(coaches[2].GetStudentsCount())

}
func fakeStudents(studentNumber int) models.Students {
	students := []*models.Student{}
	for i := 1; i <= studentNumber; i++ {
		students = append(students, &models.Student{Name: fake.FirstName()})
	}
	return students
}
func fakeCoaches(coachesNumber int) models.Coaches {
	coaches := []*models.Coach{}
	for i := 1; i <= coachesNumber; i++ {
		coaches = append(coaches, &models.Coach{Name: fake.FirstName()})
	}
	return coaches
}
