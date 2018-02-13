package patterns

import "github.com/k0kubun/pp"

type (
	AbstractFactory struct {
		*Desc
	}
	iPet interface {
		speak() string
		String() string
	}

	iFactory interface {
		getFood() string
		getPet() iPet
	}

	petShop struct {
		petFactory iFactory
	}

	dog struct{}
	cat struct{}

	dogFactory struct{}
	catFactory struct{}
)

func (m *petShop) showPet() {
	pet := m.petFactory.getPet()
	pp.Println("This is a lovely ", pet.String())
	pp.Println("It says ", pet.speak())
	pp.Println("It eats", m.petFactory.getFood())
}

// implement iPet
func (p *dog) speak() string {
	return "woaf"
}

func (p *dog) String() string {
	return "dog"
}

func (p *cat) speak() string {
	return "meow"
}

func (p *cat) String() string {
	return "cat"
}

// implement iFactory
func (p *dogFactory) getFood() string {
	return "dog food"
}

func (p *dogFactory) getPet() iPet {
	return new(dog)
}

func (p *catFactory) getFood() string {
	return "cat food"
}

func (p *catFactory) getPet() iPet {
	return new(cat)
}

func (m AbstractFactory) Do() {

	desc.SetDesc("AbstractFactory", "工厂和产品的对应关系是不变的", "生成和加工产品", "生成产品的过程", "", "用处好像不大")
	desc.print()

	list := []iFactory{new(dogFactory), new(catFactory)}
	shop := petShop{}
	for _, factory := range list {
		shop.petFactory = factory
		shop.showPet()
	}
}
