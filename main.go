package main

import (
	"github.com/icrowley/fake"
	"student-distribution/models"
	"student-distribution/services"
)

func main() {
	//arr := []int{0, 0, 0}
	//
	//sort.Sort(sort.IntSlice(arr))

	//x := services.BasicDistribution(arr, result, 7)
	//fmt.Println(x)
	coaches := fakeCoaches(3)
	students := fakeStudents(10)
	DistributionAlgorithm := services.DistributionAlgorithm{}
	DistributionAlgorithm.Coaches = coaches
	DistributionAlgorithm.Students = students
	DistributionAlgorithm.Coaches[2].Students = fakeStudents(6)
	DistributionAlgorithm.Coaches[1].Students = fakeStudents(7)
	DistributionAlgorithm.Coaches[0].Students = fakeStudents(7)

	//
	DistributionAlgorithm.FairDistribution()
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
func fakeStudents(studentNumber int) []*models.Student {
	students := []*models.Student{}
	for i := 1; i <= studentNumber; i++ {
		students = append(students, &models.Student{Name: fake.FirstName()})
	}
	return students
}
func fakeCoaches(coachesNumber int) []*models.Coach {
	coaches := []*models.Coach{}
	for i := 1; i <= coachesNumber; i++ {
		coaches = append(coaches, &models.Coach{Name: fake.FirstName()})
	}
	return coaches
}
