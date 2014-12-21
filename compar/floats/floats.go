package main

import "math"

func Compar(as, bs *[], a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}

// BEGIN OF BOILERPLATE
func ComparΞfloat32(as, bs *[]float32, a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}
func ComparΞfloat64(as, bs *[]float64, a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}
// END OF BOILERPLATE

func main() {
	var a = []float32{7., 5., 7., 999., 99999., 3.}
	var b = []float64{7., 5., 7., 999., 99999., 3.}


	if(ComparΞfloat32(&a, &a, 0, 1) <= 0) {
		panic("miscompared 0, 1\n")
	}

	if(ComparΞfloat32(&a, &a, 0, 2) != 0) {
		panic("miscompared 0, 2\n")
	}
	if(ComparΞfloat32(&a, &a, 0, 3) >= 0) {
		panic("miscompared 0, 3\n")
	}


	if(ComparΞfloat64(&b, &b, 0, 1) <= 0) {
		panic("miscompared 0, 1\n")
	}

	if(ComparΞfloat64(&b, &b, 0, 2) != 0) {
		panic("miscompared 0, 2\n")
	}
	if(ComparΞfloat64(&b, &b, 0, 3) >= 0) {
		panic("miscompared 0, 3\n")
	}

}
