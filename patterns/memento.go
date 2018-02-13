package patterns

import (
	"fmt"
	"github.com/k0kubun/pp"
)

type (
	// 履歴を記録する
	Memento struct{}

	EditorMemento struct {
		content string
	}

	Editor struct {
		content string
	}
)

func (m *Editor) Type(w string) {
	m.content = fmt.Sprintf("%s %s", m.content, w)
}

func (m *Editor) getContent() string {
	return m.content
}

func (m *Editor) save() *EditorMemento {
	return newEditorMemento(m.content)
}

func (m *Editor) restore(memento EditorMemento) {
	m.content = memento.getContent()
}

func newEditorMemento(content string) *EditorMemento {
	return &EditorMemento{
		content: content,
	}
}

func (m *EditorMemento) getContent() string {
	return m.content
}

func (Memento) Do() {

	desc.SetDesc("Memento(记忆)", "", "", "", "undo, redo", "")
	desc.print()

	// 使用场景， 特定的时间记忆， 然后可以取回记忆

	editor := new(Editor)
	editor.Type("This is the first sentence.")
	editor.Type("This is the second sentence.")

	saved := editor.save()
	editor.Type("and this is third.")
	pp.Println(editor.getContent())
	editor.restore(*saved)
	pp.Println(editor.getContent())

}
