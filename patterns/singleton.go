package patterns

import "github.com/k0kubun/pp"

type (
	Singleton struct{}

	president struct {
		name string
	}
)

var instance *president

func (Singleton) Do() {

	desc.SetDesc("Singleton(单例)", "只创建一次", "getInstance", "instance", "", "用于可以复用的实例")
	desc.print()

	a := getInstance()
	a.name = "a"
	b := getInstance()
	b.name = "b"

	pp.Println(a, b)
}

func getInstance() *president {
	if instance == nil {
		instance = new(president)
	}
	return instance
}
