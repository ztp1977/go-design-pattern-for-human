package patterns

import "github.com/k0kubun/pp"

type (
	Builder struct{}

	burger struct {
		size    int
		cheese  bool
		lettuce bool
		tomato  bool
	}

	burgerBuilder struct {
		size    int
		cheese  bool
		lettuce bool
		tomato  bool
	}
)

func NewBurger(b burgerBuilder) burger {
	return burger{
		size:    b.size,
		cheese:  b.cheese,
		lettuce: b.lettuce,
		tomato:  b.tomato,
	}
}

func NewBugerBuild(size int) *burgerBuilder {
	return &burgerBuilder{size: size}
}

func (m burgerBuilder) build() burger {
	return NewBurger(m)
}

func (m *burgerBuilder) addChees() *burgerBuilder {
	m.cheese = true
	return m
}

func (m *burgerBuilder) addLettuce() *burgerBuilder {
	m.lettuce = true
	return m
}

func (m *burgerBuilder) addTomato() *burgerBuilder {
	m.tomato = true
	return m
}

func (Builder) Do() {

	desc.SetDesc("Builder", "每个函数要返回实例本身", "修改属性的各种函数", "初始化的过程", "一个链状的生产流程", "通过函数的不同，可以制造不同的商品")
	desc.print()

	// 创建的过程明示化
	bg := NewBugerBuild(14).
		addChees().
		addLettuce().
		addTomato().build()
	pp.Println(bg)
}
