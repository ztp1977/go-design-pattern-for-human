package main

import (
	"flag"
	"github.com/k0kubun/pp"
	"go-design-pattern-for-human/patterns"
)

type Runner interface {
	Do()
}

var Runners = map[string]Runner{
	"SimpleFactory":         patterns.SimpleFactory{},
	"FactoryMethod":         patterns.FactoryMethod{},
	"AbstractFactory":       patterns.AbstractFactory{},
	"Builder":               patterns.Builder{},
	"ProtoType":             patterns.ProtoType{},
	"Singleton":             patterns.Singleton{},
	"Adapter":               patterns.Adapter{},
	"Bridge":                patterns.Bridge{},
	"Composite":             patterns.Composite{},
	"Decorator":             patterns.Decorator{},
	"Facade":                patterns.Facade{},
	"FlyWeight":             patterns.FlyWeight{},
	"Proxy":                 patterns.Proxy{},
	"ChainOfResponsibility": patterns.ChainOfResponsibility{},
	"Command":               patterns.Command{},
	"Iterator":              patterns.Iterator{},
	"Mediator":              patterns.Mediator{},
	"Memento":               patterns.Memento{},
	"Observer":              patterns.Observer{},
	"Visitor":               patterns.Visitor{},
	"Strategy":              patterns.Strategy{},
	"State":                 patterns.State{},
	"TemplateMethod":        patterns.TemplateMethod{},
}

func main() {
	typePtr := flag.String("p", "nothing", "patternName please")
	flag.Parse()
	if r, ok := Runners[*typePtr]; !ok {
		for i := range Runners {
			pp.Println(i) // print all pattern
		}
	} else {
		r.Do()
	}
}
