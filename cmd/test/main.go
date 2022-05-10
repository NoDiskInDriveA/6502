package main

import "fmt"

type Base interface {
	B()
}

type Ext interface {
	Base
	E()
}

func main() {
	t := int8(-80)
	ut := uint8(t)
	fmt.Printf("%b %b", t, ut)
}
