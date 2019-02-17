package models

type Student struct {
	Name    string   `json:name`
	Coaches []*Coach `json:coaches`
}

func (student *Student) AppendCoach(coach *Coach) {

}
func (student *Student) GetCoachesCount() int {
	return len(student.Coaches)
}
