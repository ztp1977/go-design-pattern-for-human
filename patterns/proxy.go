package patterns

import "fmt"

type (
	Proxy struct{}

	iDoor interface {
		open()
		close()
	}

	security struct {
		door iDoor // 这个部分可以换
	}

	labDoor struct{}
)

func newSecurity(door iDoor) *security {
	return &security{
		door: door,
	}
}

func (security) auth(password string) bool {
	return password == "1234"
}

func (m security) open(password string) {
	// 添加一个新功能
	if m.auth(password) {
		m.door.open()
	} else { // 处理新的Error
		fmt.Printf("Big no!, It ain't passible.\n")
	}
}

// 关门保持原来的动作
func (m security) close() {
	m.door.close()
}

func (labDoor) open() {
	fmt.Println("Opening lab door.")
}

func (labDoor) close() {
	fmt.Println("Close lab door.")
}

func (Proxy) Do() {

	desc.SetDesc("Proxy()", "也是一个Wrap, 但是和原类有着共同的方法", "", "", "", "")
	desc.print()

	door := new(labDoor)
	door.open()
	door.close()

	// 做普通的门加一把锁
	d := newSecurity(door)
	d.open("I guest!!")
	d.open("1234")
	d.close()
}
