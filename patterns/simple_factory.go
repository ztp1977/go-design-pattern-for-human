package patterns

import (
	"fmt"
	"reflect"
)

type doorTypes int

const (
	wooden doorTypes = iota + 1
	iron
)

type SimpleFactory struct{}

type (
	Door interface {
		getWidth() float32
		getHeight() float32
	}
	WoodenDoor struct {
		width  float32
		height float32
	}
	IronDoor struct {
		width  float32
		height float32
	}
)

func (m WoodenDoor) getWidth() float32 {
	return m.width
}

func (m WoodenDoor) getHeight() float32 {
	return m.height
}

func (m IronDoor) getWidth() float32 {
	return m.width * 100
}

func (m IronDoor) getHeight() float32 {
	return m.height * 100
}

func makeDoor(w, h float32, t doorTypes) Door {
	switch t {
	case wooden:
		return &WoodenDoor{w, h}
	case iron:
		return &IronDoor{w, h}
	}
	return nil
}

func (m SimpleFactory) Do() {

	desc.SetDesc("SimpleFactory()", "", "", "", "", "")
	desc.print()

	d1 := makeDoor(1, 2, wooden)
	d2 := makeDoor(2, 3, iron)
	fmt.Println(reflect.TypeOf(m), "ドアを作成して、サイズを呼べることを意識しかしない")
	fmt.Printf("%0.1f, %0.1f, %v\n", d1.getWidth(), d1.getWidth(), d1)
	fmt.Printf("%0.1f, %0.1f, %v", d2.getWidth(), d2.getWidth(), d2)
}
