package patterns

import (
	"fmt"
)

type (
	Desc struct {
		Name    string // 模式名称
		Fix     string // 固定的部分
		Open    string // 开放的接口
		Hide    string // 隐藏的实现
		Usage   string // 使用场景
		Comment string // 评价
	}
)

var desc Desc

func (m *Desc) SetDesc(name, fix, open, hide, usage, comment string) {
	m.Name = name
	m.Fix = fix
	m.Open = open
	m.Hide = hide
	m.Usage = usage
	m.Comment = comment
}

func (m *Desc) print() {
	fmt.Println("=======")
	fmt.Println("模式名: ", m.Name)
	fmt.Println("固定的部分: ", m.Fix)
	fmt.Println("开放的接口: ", m.Open)
	fmt.Println("隐藏的实现: ", m.Hide)
	fmt.Println("使用场景: ", m.Usage)
	fmt.Println("评价: ", m.Comment)
	fmt.Println("=======")
}
