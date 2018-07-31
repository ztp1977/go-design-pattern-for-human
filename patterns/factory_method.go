package patterns

import (
	"github.com/k0kubun/pp"
)

type (
	interviewer interface {
		askQuestion() string
	}

	developer          struct{}
	communityExecutive struct{}

	manager interface {
		makeInterview() interviewer
	}

	hiringManager struct {
		m manager
	}

	developManager struct {
		hiringManager
	}
	marketingManager struct {
		hiringManager
	}

	FactoryMethod struct {
	}
)

func (developer) askQuestion() string {
	return "Asking questions about patterns"
}

func (communityExecutive) askQuestion() string {
	return "Asking questions about community building"
}

func (developManager) makeInterview() interviewer {
	return new(developer)
}

func (m developManager) interView() string {
	return m.takeInterview(m)
}

func (marketingManager) makeInterview() interviewer {
	return new(communityExecutive)
}

func (m marketingManager) interView() string {
	return m.takeInterview(m)
}

func (*hiringManager) takeInterview(m manager) string {
	return m.makeInterview().askQuestion()
}

func (m FactoryMethod) printMsg() {
	pp.Println("FactoryMethod, abstract factory みたいですが、goは継承できないので、このパターンがほぼ使えないのでは")
}

func (m FactoryMethod) Do() {

	desc.SetDesc("FactoryMethod()", "", "", "", "", "")
	desc.print()

	m.printMsg()

	devManager := new(developManager)
	pp.Println(devManager.interView())

	markManager := new(marketingManager)
	pp.Println(markManager.interView())
}
