package patterns

import "github.com/k0kubun/pp"

type (
	Observer  struct{}
	iObserver interface {
		update(string)
	}

	iSubject interface {
		addObserver(observer iObserver)
		notifyObservers()
	}

	teacher struct {
		Name string
	}

	student struct {
		Name      string
		observers []iObserver
	}
)

// implements iObserver
func (t *teacher) update(msg string) {
	pp.Printf("%s catches %s\n", t.Name, msg)
}

// implements iSubject
func (s *student) addObserver(obs iObserver) {
	// オブサーバーを追加
	s.observers = append(s.observers, obs)
}

func (s *student) notifyObservers() {
	// 通知するイベントを全員送る。
	for _, v := range s.observers {
		v.update("hello from " + s.Name)
	}
}

func (Observer) Do() {

	desc.SetDesc("Observer(观察者)", "分离了事件发动者和事件处理者的依赖关系", "添加观察者，事件启动", "观察者动作", "", "这个也是一个超级重要的模式，客户端的基于event的实现，都和这个有关系")
	desc.print()

	t1 := &teacher{"Teacher Saith"}
	t2 := &teacher{"Teacher Tatiana"}

	subjects := make([]iSubject, 0)

	// add observer teacher
	s1 := &student{"Student Taro", make([]iObserver, 0)}
	s1.addObserver(t1)
	s1.addObserver(t2)

	subjects = append(subjects, s1)
	for _, s := range subjects {
		// send notify to teacher
		s.notifyObservers()
	}
}
