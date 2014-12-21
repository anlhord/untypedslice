package main

import (
	"fmt"
)

func triple(a *[]) {
// FIXME
//	*a = append(*a, *a...)
//	*a = append(*a, *a...)
}
func main() {
	a := []byte("1337")
	b := []int{1,3,3,7}
	c := []interface{}{1, "3", 3.0, 7}

// FIXME
//	a = triple(a)
//	b = triple(b)
//	c = triple(c)

	fmt.Println("A:", a, "B:", b, "C:", c)
}
