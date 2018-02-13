package patterns

import "github.com/k0kubun/pp"

type (
	Command  struct{}
	iCommand interface {
		execute()
		undo()
		redo()
	}

	bulb    struct{}
	turnoff struct {
		bulb *bulb
	}
	turnon struct {
		bulb *bulb
	}
	remoteControl struct{}
)

func (bulb) TurnOn() {
	pp.Println("bulb has been it.")
}

func (bulb) TurnOff() {
	pp.Println("Darkness.")
}

func newTurnOff(b *bulb) *turnoff {
	return &turnoff{
		bulb: b,
	}
}

func newTurnOn(b *bulb) *turnon {
	return &turnon{
		bulb: b,
	}
}

func (m turnon) execute() {
	m.bulb.TurnOn()
}

func (m turnon) undo() {
	m.bulb.TurnOff()
}

func (m turnon) redo() {
	m.execute()
}

func (m turnoff) execute() {
	m.bulb.TurnOff()
}

func (m turnoff) undo() {
	m.bulb.TurnOn()
}

func (m turnoff) redo() {
	m.execute()
}

func (remoteControl) submit(command iCommand) {
	command.execute()
}

func (Command) Do() {
	// 有点儿像浏览器的哪个undo键

	bulb := new(bulb)
	turnOn := newTurnOn(bulb)
	turnOff := newTurnOff(bulb)
	remote := new(remoteControl)
	remote.submit(turnOn)
	remote.submit(turnOff)
	remote.submit(turnOff)
	remote.submit(turnOn)
}
