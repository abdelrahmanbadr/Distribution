package models

type Students []*Student
type Student struct {
	Name    string   `json:name`
	Coaches []*Coach `json:coaches`
}

func (student *Student) AppendCoach(coach *Coach) {

}
func (student *Student) GetCoachesCount() int {
	return len(student.Coaches)
}

func (student *Student) AddStudent(coach *Coach) {
	student.Coaches = append(student.Coaches, coach)
}
