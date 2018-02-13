package patterns

import "github.com/k0kubun/pp"

type (
	// 这是一个使用手册，顺序是固定的
	TemplateMethod struct{}

	iBuilder interface {
		test()
		lint()
		assemble()
		deploy()
	}

	builder struct {
	}

	androidBuilder struct {
		*builder
	}
	iosBuilder struct {
		*builder
	}
)

func (androidBuilder) test() {
	pp.Println("android test...")
}

func (androidBuilder) lint() {

	pp.Println("android lint...")
}

func (androidBuilder) assemble() {

	pp.Println("android assemble...")
}

func (androidBuilder) deploy() {

	pp.Println("android deploy...")
}
func (iosBuilder) test() {

	pp.Println("ios test...")
}

func (iosBuilder) lint() {

	pp.Println("ios lint...")
}

func (iosBuilder) assemble() {

	pp.Println("ios assemble...")
}

func (iosBuilder) deploy() {

	pp.Println("ios deploy...")
}

func build(b iBuilder) {
	b.test()
	b.lint()
	b.assemble()
	b.deploy()
}

func (TemplateMethod) Do() {

	desc.SetDesc("TemplateMethod(临时方法)", "创建顺序不变", "各个步骤", "每个创建步骤的实际操作", "", "有操作手册的情况下")
	desc.print()

	iosApp := new(iosBuilder)
	androidApp := new(androidBuilder)
	build(iosApp)
	build(androidApp)
}
