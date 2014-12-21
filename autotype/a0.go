// automatic type inference for a local variable
package main

import (
//	"fmt"
)

func semi() string {
//	var xa [] = []int{1,3,3,7}
//	var xb [] = []string{"13","37"}
//	var xc [] = []byte("l33t")
//	var xd [] = []float64{1.3,3.7}

////////////
	var xa []int	= []int{1,3,3,7}
	var xb []string	= []string{"13","37"}
	var xc []byte	= []byte("l33t")
	var xd []float64	= []float64{1.3,3.7}
////////////

	ia := []int{1,3,3,7}
	ib := []string{"13","37"}
	id := []float64{1.3,3.7}

//	return fmt.Sprintf("%v %v %v %v  %v %v %v %v",
//		xa,xb,xc,xd, ia,ib,id)
	return ""
}

func manual() string {
	var xa []int	= []int{1,3,3,7}
	var xb []string	= []string{"13","37"}
	var xc []byte	= []byte("l33t")
	var xd []float64	= []float64{1.3,3.7}

	ia := []int{1,3,3,7}
	ib := []string{"13","37"}
	id := []float64{1.3,3.7}

//	return fmt.Sprintf("%v %v %v %v  %v %v %v %v",
//		xa,xb,xc,xd, ia,ib,id)
	return ""
}

func auto() string {
//	var xa [] = []{1,3,3,7}
//	var xb [] = []{"13","37"}
//	var xc [] = []byte("l33t")
//	var xd [] = []{1.3,3.7}
////////////
	var xa []int	= []int{1,3,3,7}
	var xb []string	= []string{"13","37"}
	var xc []byte	= []byte("l33t")
	var xd []float64	= []float64{1.3,3.7}
////////////
//	ia := []{1,3,3,7}
//	ib := []{"13","37"}
//	id := []{1.3,3.7}
////////////
	ia := []int{1,3,3,7}
	ib := []string{"13","37"}
	id := []float64{1.3,3.7}
////////////
//	return fmt.Sprintf("%v %v %v %v  %v %v %v %v",
//		xa,xb,xc,xd, ia,ib,id)
	return ""
}

func main() {
	a := auto()
	s := semi()
	m := manual()

	if a != m {
		panic("Error auto.")
	}

	if s != m {
		panic("Error semi.")
	}
}
