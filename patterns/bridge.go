package patterns

import (
	"github.com/k0kubun/pp"
)

// 换皮换皮了

type (
	Bridge struct{}

	webPage interface {
		getContent() string
	}

	about struct {
		theme theme
	}
	careers struct {
		theme theme
	}

	theme interface {
		getColor() string
	}
	darkTheme  struct{}
	lightTheme struct{}
)

func NewAboutPage(theme theme) webPage {
	return &about{
		theme: theme,
	}
}

func NewCareersPage(theme theme) webPage {
	return &careers{
		theme: theme,
	}
}

func (m *about) getContent() string {
	return "About page in " + m.theme.getColor()
}

func (m *careers) getContent() string {
	return "Careers page in " + m.theme.getColor()
}

func (darkTheme) getColor() string {
	return "Dark Black"
}

func (lightTheme) getColor() string {
	return "Light White"
}

func (m Bridge) Do() {

	desc.SetDesc("Bridge", "插入点是已知的", "显示页面的接口", "不同主题的实现", "网页的主题或者皮肤", "量产的时候有用")
	desc.print()

	dt := new(darkTheme)
	lt := new(lightTheme)

	pp.Println(NewAboutPage(dt).getContent())
	pp.Println(NewCareersPage(lt).getContent())
}
