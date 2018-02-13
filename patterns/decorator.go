package patterns

import "github.com/k0kubun/pp"

type (
	Decorator struct{}

	iCoffee interface {
		getCost() int64
		getDescription() string
	}

	simpleCoffee struct {
	}
	milkCoffee struct {
		c iCoffee
	}
	whipCoffee struct {
		c iCoffee
	}
	vanillaCoffee struct {
		c iCoffee
	}
)

func newMilkCoffee(coffee iCoffee) *milkCoffee {
	return &milkCoffee{c: coffee}
}

func newWhipCoffee(coffee iCoffee) *whipCoffee {
	return &whipCoffee{c: coffee}
}

func newVanillaCoffee(coffee iCoffee) *vanillaCoffee {
	return &vanillaCoffee{c: coffee}
}

func (m *simpleCoffee) getCost() int64 {
	return 10
}

func (m *simpleCoffee) getDescription() string {
	return "simple coffee"
}

func (m *milkCoffee) getCost() int64 {
	return m.c.getCost() + 10
}

func (m *milkCoffee) getDescription() string {
	return m.c.getDescription() + "+milk"
}

func (m *whipCoffee) getCost() int64 {
	return m.c.getCost() + 100
}

func (m *whipCoffee) getDescription() string {
	return m.c.getDescription() + "+whip"
}

func (m *vanillaCoffee) getCost() int64 {
	return m.c.getCost() + 1000
}

func (m *vanillaCoffee) getDescription() string {
	return m.c.getDescription() + "+vanilla"
}

func (Decorator) Do() {

	desc.SetDesc("Decorator(装饰)", "包裹起来生成另一个包裹", "初始化方法", "包裹起来的实例", "errors.Wrap就是这个模式", "通过不断修饰，在不毁坏旧逻辑的情况下，重写新逻辑")
	desc.print()

	// 也是一个链条， 有共同的I/F
	var coffee iCoffee
	coffee = new(simpleCoffee)
	pp.Println(coffee.getCost(), coffee.getDescription())

	coffee = newMilkCoffee(coffee)
	pp.Println(coffee.getCost(), coffee.getDescription())

	coffee = newWhipCoffee(coffee)
	pp.Println(coffee.getCost(), coffee.getDescription())

	coffee = newVanillaCoffee(coffee)
	pp.Println(coffee.getCost(), coffee.getDescription())
}
