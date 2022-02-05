package main

import "fmt"

type a struct {
	b *b
}

type b struct {
	c int
}

func main() {
	x := a{&b{1}}
	y := a{}
	fmt.Println(x.b.c)
	if y.b != nil {
		fmt.Println(y.b.c)
	} else {
		fmt.Println("C is null")
	}
}
