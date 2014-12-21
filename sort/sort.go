package main

import "math"
import "sort"
import "fmt"

func Compar(as, bs *[], a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}

/*
FIXME: sort.go:16: internal compiler error: fault
func Sort(slice *[], compar func(as, bs *[], a, b int) int8) {
	for i := 0; i < len(*slice); i++ {
	for j := 0; j < i; j++ {
		if compar(slice, slice, i, j) > 0 {
			x := (*slice)[i]
			(*slice)[i] = (*slice)[j]
			(*slice)[j] = x
		}
	}}
}
*/
// BEGIN OF BOILERPLATE
func ComparΞfloat32(as, bs *[]float32, a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}

func SortΞfloat32(slice *[]float32, compar func(as, bs *[]float32, a, b int) int8) {
	for i := 0; i < len(*slice); i++ {
	for j := 0; j < i; j++ {
		if compar(slice, slice, i, j) > 0 {
			x := (*slice)[i]
			(*slice)[i] = (*slice)[j]
			(*slice)[j] = x
		}
	}}
}


// END OF BOILERPLATE


func main() {
	var a = []float32{818,128,39,153,643}
	var b = []float32{9,8,7,6,5,4,3,1,2}

	SortΞfloat32(&a, ComparΞfloat32)
	SortΞfloat32(&b, ComparΞfloat32)


	x := fmt.Sprint(a)
	y := fmt.Sprint(b)

	if x != "[818 643 153 128 39]" {
		panic("a")
	}
	if y != "[9 8 7 6 5 4 3 2 1]" {
		panic("b")
	}

}
