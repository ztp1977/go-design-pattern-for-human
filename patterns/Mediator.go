package patterns

import (
	"github.com/k0kubun/pp"
	"time"
)

type (
	Mediator struct{}

	chatroomMediator interface {
		showMessage(user User, msg string)
	}

	User struct {
		name string
		crm  chatroomMediator
	}

	chatroom struct {
	}
)

func newUser(name string, crm chatroomMediator) *User {
	return &User{
		name: name,
		crm:  crm,
	}
}

func (m User) send(message string) {
	m.crm.showMessage(m, message)
}

func (m User) getName() string {
	return m.name
}

// 共通処理で、ユーザの名前と時間を表示する
func (m *chatroom) showMessage(user User, message string) {
	pp.Printf("%s %s: %s\n", time.Now(), user.getName(), message)
}

func (Mediator) Do() {

	desc.SetDesc("Mediator(调停者)", "协调连个实例之间的关系", "", "", "", "")
	desc.print()

	// 仲介者
	mediator := new(chatroom)

	jam := newUser("Jam Doe", mediator)
	tom := newUser("Tom Dow", mediator)

	jam.send("Hi there!")
	tom.send("Hey!")
}
