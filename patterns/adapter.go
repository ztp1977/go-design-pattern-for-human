package patterns

import "github.com/k0kubun/pp"

type (
	Adapter struct{}

	lion interface {
		roar() string
	}

	africanLion struct{}
	asiaLion    struct{}

	hunter struct{}

	wildDog struct{}

	wildDogAdapter struct {
		dog wildDog
	}
)

func (m africanLion) roar() string {
	return "waao~~~"
}

func (m asiaLion) roar() string {
	return "ao~~~"
}

func (m wildDog) bark() string {
	return "wang wang !!"
}

func NewWildDogAdapter(dog wildDog) *wildDogAdapter {
	return &wildDogAdapter{
		dog: dog,
	}
}

func (m wildDogAdapter) roar() string {
	return m.dog.bark()
}

func (hunter) hunt(l lion) {
	pp.Println("kill lion: ", l.roar())
}

func (m Adapter) Do() {

	desc.SetDesc("Adapter", "旧的接口不变", "狮子的接口", "野狗的接口", "利用现有接口对应新事物", "一个转换的逻辑")
	desc.print()

	pp.Println(new(africanLion).roar())
	pp.Println(new(asiaLion).roar())

	// 这个模式有点儿用， 做一个接口， 可以让现有的class用
	adapter := NewWildDogAdapter(wildDog{})
	pp.Println(adapter.roar())
	hunter := new(hunter)
	hunter.hunt(adapter)
}
