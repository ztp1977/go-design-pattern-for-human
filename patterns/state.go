package patterns

import (
	"github.com/k0kubun/pp"
	"strings"
)

type (
	State struct{}

	iWriteState interface {
		write(s string)
	}
	upperCase   struct{}
	lowerCase   struct{}
	commonPrint struct{}

	textEditor struct {
		state iWriteState
	}
)

func (m *textEditor) setState(state iWriteState) {
	m.state = state
}

func (m textEditor) Type(words string) {
	m.state.write(words)
}

// implements iWriteState
func (upperCase) write(s string) {
	pp.Printf("%s\n", strings.ToUpper(s))
}
func (lowerCase) write(s string) {
	pp.Printf("%s\n", strings.ToLower(s))
}

func (commonPrint) write(s string) {
	pp.Printf("%s\n", s)
}

func (State) Do() {

	desc.SetDesc("State(状态)", "", "", "", "", "")
	desc.print()

	editor := &textEditor{
		state: new(commonPrint),
	}

	editor.Type("First Line.")

	editor.setState(new(upperCase))
	editor.Type("Second Line.")

	editor.setState(new(lowerCase))
	editor.Type("Third Line")

}
