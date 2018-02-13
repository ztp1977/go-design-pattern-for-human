package patterns

import "github.com/k0kubun/pp"

type (
	Facade         struct{}
	computer       struct{}
	computerSwitch struct {
		computer *computer
	}
)

func (computer *computer) getElectricShock() {
	pp.Println("打开电源!")
}
func (computer *computer) makeSound() {
	pp.Println("Beep beep!")
}
func (computer *computer) showLoaddingScreen() {
	pp.Println("显示Loading...")
}
func (computer *computer) bam() {
	pp.Println("都准备好了")
}
func (computer *computer) closeEverything() {
	pp.Println("关闭电脑了")
}
func (computer *computer) sooth() {
	pp.Println("sooth")
}
func (computer *computer) pullCurrent() {
	pp.Println("current")
}

func (m computerSwitch) trunOff() {
	m.computer.getElectricShock()
	m.computer.makeSound()
	m.computer.showLoaddingScreen()
	m.computer.bam()
}

func (m computerSwitch) trunOn() {
	m.computer.closeEverything()
	m.computer.pullCurrent()
	m.computer.sooth()
}

func (Facade) Do() {

	desc.SetDesc("Facade(门脸儿)", "公开最简单的接口", "只暴露需要的接口", "复杂的步骤", "", "")
	desc.print()

	// 复杂的处理，对外只是一个简单的接口
	c := &computerSwitch{new(computer)}
	c.trunOn()
	c.trunOff()
}
