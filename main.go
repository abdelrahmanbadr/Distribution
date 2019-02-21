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

	coaches[0].Students = fakeStudents(10)
	coaches[1].Students = fakeStudents(14)
	coaches[2].Students = fakeStudents(3)
	//coaches[3].Students = fakeStudents(0)
	distributionAlgorithm := services.NewDistributionAlgorithm(students, coaches)
	distributionAlgorithm.FairDistribution()

	fmt.Println(coaches[0].GetStudentsCount())
	fmt.Println(coaches[1].GetStudentsCount())
	fmt.Println(coaches[2].GetStudentsCount())
	//fmt.Println(coaches[3].GetStudentsCount())

	//arr := []int{0, 0, 0}
	//
	//sort.Sort(sort.IntSlice(arr))

	//x := services.BasicDistribution(arr, result, 7)
	//fmt.Println(x)
	//coaches := fakeCoaches(3)
	//students := fakeStudents(10)
	//DistributionAlgorithm := services.DistributionAlgorithm{}
	//DistributionAlgorithm.Coaches = coaches
	//DistributionAlgorithm.Students = students
	//
	////
	//DistributionAlgorithm.FairDistribution()
	//fmt.Println(len(DistributionAlgorithm.Coaches[0].Students))
	//fmt.Println(len(DistributionAlgorithm.Coaches[1].Students))
	//fmt.Println(len(DistributionAlgorithm.Coaches[2].Students))
	//fmt.Println(students[0].Coaches[0].Name)
	//fmt.Println(students[1].Coaches[0].Name)
	//fmt.Println(students[2].Coaches[0].Name)
	//fmt.Println(students[3].Coaches[0].Name)
	//fmt.Println(students[4].Coaches[0].Name)
	//fmt.Println(students[5].Coaches[0].Name)
	//fmt.Println(students[6].Coaches[0].Name)
	//fmt.Println(students[7].Coaches[0].Name)
	//fmt.Println(students[8].Coaches[0].Name)
	//fmt.Println(students[9].Coaches[0].Name)

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
