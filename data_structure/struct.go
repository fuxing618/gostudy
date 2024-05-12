package data_structure

type People struct {
	Name string
	Age int
}

func (p *People) SetName(name string, age int) {
	p.Name = name
	p.Age = age
}

func (p *People) getName() string {
	return p.Name
}
