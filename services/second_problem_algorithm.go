package services

//1,2,10 ---> number to be distributed 9

//way to solve it
// 1+2+10 + 9 = 22

//(1+2+9/2 = 6) so every one of them need to reach 6 except the last one take the change
//8 is greater than  10  -> no
//8 is greater than  1   -> yes  ->   = 5  need to reach 6
//8 is greater than  2   -> yes  ->  = 4   the remain
//----------------------------------------------------------------------

//6,6,7   --> number to be distributed 10

//way to solve it
// 6+7+8 + 10 = 31
//31 /3 = 11 or 10

//(31 / 3 = 10) so every one of them need to reach 10 except the last one take the change
//10 is greater than  6  ->  yes ->  4    //need to reach 10
//10 is greater than  7  ->  yes ->  3    //need to reach 10
//10 is greater than  8  ->  yes ->  3    //take the remain
//----------------------------------------------------------------------
// follow the next steps

//6,7,8,12 ---> number to be distributed 5

// 6+7+8+12+ 5 = 38 ,,,,   38 / 4 = 10
//10 is greater than  6 ->  yes
//10 is greater than  7  ->  yes
//10 is greater than  8  ->  yes
//10 is greater than  12  ->  no

//count of yes is 3

//(6+7+8+5/3 = 9) so every one of them need to reach 9 except the last one take the change

// 6   ->  3    //need to 3 reach 9
// 7   ->  2    //need to 3 reach 9
// 8   ->  0   //take the remain
//------------------------------------------------------
//0,0,0  --> number is 10
//10 / 3 ceil is 4 fuck shit
//type Distrib struct {
//	Students models.Students
//	Coaches  models.Coaches
//}
//
//func (self *Distrib) FairDistribution() {
//	coachesNumber := len(self.Coaches)
//	studentsNumber := len(self.Students)
//	totalNumber := studentsNumber
//	sort.Sort(sort.Interface(self.Coaches))
//	for _, coach := range self.Coaches {
//		totalNumber += coach.GetStudentsCount()
//	}
//	//fmt.Println(totalNumber)
//	division := int(math.Ceil(float64(totalNumber) / float64(coachesNumber)))
//
//	counter := 0
//	for _, coach := range self.Coaches {
//
//		if coach.GetStudentsCount() >= division {
//			totalNumber -= coach.GetStudentsCount()
//			continue
//		}
//		counter++
//	}
//	//fmt.Println(totalNumber)
//	division = int(math.Ceil(float64(totalNumber) / float64(counter)))
//	studentsIndex := 0
//	for j := 0; j < counter; j++ {
//		studentsNumberForCoach := division - self.Coaches[j].GetStudentsCount()
//
//		//last key will take the change
//		if j == counter-1 {
//			studentsNumberForCoach = studentsNumber - studentsIndex
//		}
//
//		for i := 0; i < studentsNumberForCoach; i++ {
//			self.Coaches[j].Students = append(self.Coaches[j].Students, self.Students[studentsIndex])
//			studentsIndex++
//		}
//
//	}
//
//}
