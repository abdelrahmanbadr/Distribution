package models

type Coach struct {
	Name     string     `json:name`
	Students []*Student `json:students`
}

func (coach *Coach) AppendStudent(student *Student) {
	//a = append(a, 4)
}
