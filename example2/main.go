package main

import "github.com/jonnyorman/pubsubmit"

type Model struct {
	Prop1 string
	Prop2 int
}

func main() {
	pubsubmit.RunTyped[Model]()
}
