package models

type Coaches []*Coach
type Coach struct {
	Name     string     `json:name`
	Students []*Student `json:students`
	Capacity int        `json:capacity`
}

func (coach *Coach) GetStudentsCount() int {
	return len(coach.Students)
}
func (c Coaches) Len() int {
	return len(c)
}
func (c Coaches) Less(i, j int) bool {
	return c[i].GetStudentsCount() < c[j].GetStudentsCount()
}

func (c Coaches) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (coach *Coach) AddStudent(student *Student) {
	coach.Students = append(coach.Students, student)
}
