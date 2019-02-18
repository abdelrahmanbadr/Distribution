package models

type Coaches []*Coach
type Coach struct {
	Name     string     `json:name`
	Students []*Student `json:students`
}

func (coach *Coach) AppendStudent(student *Student) {
	//a = append(a, 4)
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
