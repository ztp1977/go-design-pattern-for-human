package patterns

import (
	"github.com/k0kubun/pp"
)

type (
	ProtoType struct{}

	sheep struct {
		name     string
		category string
	}
)

func NewSheep(name, category string) *sheep {
	return &sheep{
		name:     name,
		category: category,
	}
}

func (m *sheep) setname(name string) {
	m.name = name
}

func (m *sheep) getname() string {
	return m.name
}

func (m *sheep) setCategory(category string) {
	m.category = category
}

func (m *sheep) getCategory() string {
	return m.category
}

func (m sheep) Clone() sheep {
	return sheep{m.name, m.category}
}

func (m ProtoType) Do() {

	desc.SetDesc("ProtoType(原型)", "JS中超级重要的概念，这里只是clone，毕竟clone更快", "", "", "", "")
	desc.print()

	s := NewSheep("mimi", "Mountain Sheep")
	pp.Println(s)
	s2 := s.Clone()
	s2.setCategory("sheep")
	pp.Println(s2)

}
