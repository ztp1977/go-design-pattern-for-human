package patterns

import "github.com/k0kubun/pp"

type (
	// Vistor pattern lets you add further operations to objects without having to modify them
	Visitor struct{}

	Monkey struct{}
	Loin   struct{}

	iVisitor interface {
		accept(o iOperation)
	}

	iOperation interface {
		visitMonkey(monkey Monkey)
		visitLion(loin Loin)
	}

	speak struct{}
	jump  struct{}
)

// implements iVisitor for animal
func (m Loin) accept(o iOperation) {
	o.visitLion(m)
}

func (m Loin) roar() {
	pp.Println("Loin Roaaar!")
}

func (m Monkey) accept(o iOperation) {
	o.visitMonkey(m)
}

func (m Monkey) shout() {
	pp.Println("Monkey!! Ooh oo aa ")
}

// implements iOperation
func (m speak) visitMonkey(monkey Monkey) {
	monkey.shout()
}

func (m speak) visitLion(loin Loin) {
	loin.roar()
}

func (m jump) visitMonkey(monkey Monkey) {
	pp.Println("Jumped 20 feet high! on to the tree!")
}

func (m jump) visitLion(loin Loin) {
	pp.Println("Jumped 7 feet! Back on the ground!")
}

func (Visitor) Do() {
	desc.SetDesc("Visitor(访问者)", "动物的种类基本固定，动作需要添加", "执行动作的接口(accept)", "动作的实现", "", "")
	desc.print()

	// open[I/F] is accept
	// 方便追加动作， 但是追加访问者的时候比较麻烦
	money := new(Monkey)
	loin := new(Loin)
	speak := new(speak)
	jump := new(jump)

	money.accept(speak)
	money.accept(jump)
	loin.accept(speak)
	loin.accept(jump)
}
