package main

import "github.com/RGood/go-collection-functions/pkg/set"

type Foo interface {
	comparable
	Yell()
}

type bar struct{}

func (bar *bar) Yell() {
	println("bar")
}

func main() {
	bars := set.NewOrderedSet[Foo]()
}
