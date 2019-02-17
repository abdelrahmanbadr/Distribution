package services

import (
	"fmt"
	"sort"
	"student-distribution/models"
)

type DistributionAlgorithm struct {
	Students []*models.Student
	Coaches  []*models.Coach
}

//problem 1
func (self *DistributionAlgorithm) BasicDistribution() {
	studentdCount := len(self.Students)
	coachesCount := len(self.Coaches)
	studentsPerCoach := studentdCount / coachesCount
	coachIndex := 0
	for key, student := range self.Students {

		if key > 0 && key%studentsPerCoach == 0 && coachIndex < coachesCount-1 {
			coachIndex++;
		}
		student.Coaches = append(student.Coaches, self.Coaches[coachIndex])
		self.Coaches[coachIndex].Students = append(self.Coaches[coachIndex].Students, student)

	}
}

//problem 2
func (self *DistributionAlgorithm) FairDistribution() {
	//1- make all the coaches have the same count of students first
	mp := make([]int, len(self.Coaches))
	largestStudentsCount := 0
	for key, coach := range self.Coaches {
		mp[key] = coach.GetStudentsCount()
		if mp[key] > largestStudentsCount {
			largestStudentsCount = mp[key]
		}
	}
	sort.Sort(sort.IntSlice(mp))
	s := MakeFairDistribution(mp, make([]int, len(mp)), len(self.Students))
	fmt.Println(s)
	//for _, coach := range self.Coaches {
	//	//a = append(a[:i], a[i+1:]...)
	//
	//}

}

//arr should be sorted array
//@todo refactore this shit
func MakeFairDistribution(arr []int, result []int, number int) []int {

	equalSmallestNumbers := 1
	diffrenceBetweenSmallestNumberAndTheNextOne := 0
	//secondSmallestNumber := arr[0]
	//mp := make([]int, len(arr))

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			diffrenceBetweenSmallestNumberAndTheNextOne = arr[i] - arr[i-1]
			break
		}
		equalSmallestNumbers++

	}

	if diffrenceBetweenSmallestNumberAndTheNextOne == 0 {

		return BasicDistribution(arr, result, number)
	}
	// if all of them are equal will diffrenceBetweenSmallestNumberAndTheNextOne = 0

	if diffrenceBetweenSmallestNumberAndTheNextOne*equalSmallestNumbers >= number {
		// in this case number will be divided between equalSmallestNumbers
		val := number / equalSmallestNumbers
		for i := 0; i < equalSmallestNumbers; i++ {
			result[i] = result[i] + val
			number = number - val
		}
		if number > 0 {
			result[0] = result[0] + number
		}
		return result
	}
	//else

	for i := 0; i < equalSmallestNumbers; i++ {

		result[i] = result[i] + diffrenceBetweenSmallestNumberAndTheNextOne
		arr[i] = arr[i] + result[i]
		number = number - diffrenceBetweenSmallestNumberAndTheNextOne
	}
	//fmt.Println(equalSmallestNumbers)
	//fmt.Println(diffrenceBetweenSmallestNumberAndTheNextOne)
	//return result
	if number > 0 {
		MakeFairDistribution(arr, result, number)
	}

	return result
}

//array members are equal
func BasicDistribution(array []int, result []int, number int) []int {

	arraylength := len(array)
	eq := number / arraylength

	for i := 0; i < arraylength; i++ {
		result[i] = result[i] + eq
		number = number - eq
	}

	if number > 0 {
		for i := 0; i < number; i++ {

			result[i]++
		}
	}

	return result
}
