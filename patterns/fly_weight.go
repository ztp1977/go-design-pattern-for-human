package patterns

import "github.com/k0kubun/pp"

type (
	FlyWeight struct{}

	karaTea struct{}

	teaMaker struct {
		availableTea map[string]*karaTea
	}

	teaShop struct {
		orders   map[int]*karaTea
		teaMaker *teaMaker
	}
)

func newTeaMaker() *teaMaker {
	return &teaMaker{
		availableTea: make(map[string]*karaTea),
	}
}

func newTeaShop(teaMaker *teaMaker) *teaShop {
	return &teaShop{
		orders:   make(map[int]*karaTea),
		teaMaker: newTeaMaker(),
	}
}

func (m *teaShop) takeOrder(teaType string, table int) {
	m.orders[table] = m.teaMaker.make(teaType)
}

func (m *teaShop) serve() {
	for i, v := range m.orders {
		pp.Printf("Serving tea to table# %s, %+v \n", i, v)
	}
}

func (m *teaMaker) make(name string) *karaTea {
	_, ok := m.availableTea[name]
	if !ok {
		m.availableTea[name] = new(karaTea)
	}
	return m.availableTea[name]
}

func (FlyWeight) Do() {

	desc.SetDesc("FlyWeight(轻量)", "不重复创建实例，用空间换时间", "", "", "", "")
	desc.print()

	// 把有的instance存起来，留着以后用
	teaMaker := new(teaMaker)
	shop := newTeaShop(teaMaker)

	shop.takeOrder("less sugar", 1)
	shop.takeOrder("more milk", 4)
	shop.takeOrder("without suger", 3)
	shop.serve()
}
