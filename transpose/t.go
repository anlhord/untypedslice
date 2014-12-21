package main

import "fmt"
/*
FIXME: t.go:10: tempname called with nil type
func Transpose(matrix [], stride int) (o [], news int) {
	for i := 0; i < stride; i++ {
		for j := 0; j < len(matrix); j+=stride {
			o = append(o, matrix[i+j])
		}
	}
	return o, len(matrix) / stride
}
*/
// BEGIN OF BOILERPLATE
func TransposeΞfloat32(matrix []float32, stride int) (o []float32, news int) {
	for i := 0; i < stride; i++ {
		for j := 0; j < len(matrix); j+=stride {
			o = append(o, matrix[i+j])
		}
	}
	return o, len(matrix) / stride
}
// END OF BOILERPLATE

func main() {
	i := []float32 { 0, 1, 2, 3,
			 4, 5, 6, 7,
			 8, 9,10,11,
			12,13,14,15,
			16,17,18,19}

	j, k := TransposeΞfloat32(i, 4)
	x := fmt.Sprintf("%v  %v ", j, k)
	if x != "[0 4 8 12 16 1 5 9 13 17 2 6 10 14 18 3 7 11 15 19]  5 " {
		panic("Transpose" + x)
	}
}
