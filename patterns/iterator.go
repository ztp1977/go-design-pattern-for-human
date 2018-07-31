package patterns

import (
	"github.com/k0kubun/pp"
)

type (
	Iterator struct{}

	LoopAble interface {
		getNext() bool
	}
)

type station struct {
	frequency float32
}

type stationList struct {
	current int
	list    []*station
}

var intData = []int{
	1, 2, 3, 4, 5, 6, 7,
}

type intStatefulIterator struct {
	current int
	data    []int
}

func (it *intStatefulIterator) value() int {
	return it.data[it.current]
}

func (it *intStatefulIterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}
	return true
}

func NewIntStatefullIterator(data []int) *intStatefulIterator {
	return &intStatefulIterator{data: data, current: -1}
}

func (Iterator) Do() {

	desc.SetDesc("Iterator(迭代)", "必须是一个集合", "Next, Cap", "", "让任意的集合可以像数组一样运算", "")
	desc.print()

	// 想象一个数组的循环子，满足找到自己， 并且知道是否有下一个以便结束循环

	st := newStationList()
	st.add(newRedioStation(89))
	st.add(newRedioStation(90))
	st.add(newRedioStation(93))

	for _, v := range st.list {
		pp.Println(v.getFrequency())
	}
	st.RemoveStation(newRedioStation(90))

	pp.Println(st)

	// ちゃんとgoの書き方すると
	var sum = 0
	cb := func(val int) {
		sum += val
	}
	// with callback
	IntCallbackIterator(cb)
	pp.Println("callback: ", sum)

	// with channel buffer
	sum = 0
	for val := range IntChannelIterator() {
		sum += val
	}
	pp.Println("channel: ", sum)

	// with closure
	sum = 0
	var val int
	for it, hasNext := IntClosureIterator(); hasNext; {
		val, hasNext = it()
		//pp.Println(val, hasNext)
		sum += val
	}
	pp.Println("closure iterator: ", sum)

	// with stateful iterators
	sum = 0
	it := NewIntStatefullIterator(intData)
	for it.Next() {
		sum += it.value()
	}
	pp.Println("stateful iterator: ", sum)

	// with channel
	// Do like this
	//for item := range Foo.Iterator() {
	//	pp.Println(item)
	//}
}

func IntClosureIterator() (func() (int, bool), bool) {
	var idx = 0
	var dataLen = len(intData)
	return func() (int, bool) {
		prevIdx := idx
		idx++
		return intData[prevIdx], idx < dataLen
	}, idx < dataLen
}

func IntChannelIterator() <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range intData {
			ch <- val
		}
		close(ch)
	}()
	return ch
}

func IntCallbackIterator(cb func(int)) {
	for _, val := range intData {
		cb(val)
	}
}

// それぞれのcreatorが必要
func newRedioStation(feq float32) *station {
	return &station{
		frequency: feq,
	}
}

func (m station) getFrequency() float32 {
	return m.frequency
}

func (m *stationList) add(station *station) {
	m.list = append(m.list, station)
}

func (m *stationList) RemoveStation(s *station) {

	list := make([]*station, 0)
	for _, v := range m.list {
		if v.frequency != s.frequency {
			list = append(list, v)
		}
	}
	m.list = list
}

func newStationList() *stationList {
	return &stationList{
		list: make([]*station, 0),
	}
}

func (m *stationList) getNext() bool {
	m.current++
	return m.current < len(m.list)
}
