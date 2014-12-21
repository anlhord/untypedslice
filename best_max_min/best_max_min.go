package main

import "math"

func Compar(as, bs *[], a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}

func ReverseCompar(as, bs *[], a, b int) int8 {
	return -Compar(as, bs, a, b)
}

func Best(s *[], better func(as, bs *[], a, b int) int8) (j int) {
	for i := range *s {
		if better(s,s,i,j) > 0 {
			j = i
		}
	}
	return j
}

// BEGIN OF BOILERPLATE
func ComparΞfloat32(as, bs *[]float32, a, b int) int8 {
	return int8(math.Mod(float64((*as)[a]-(*bs)[b]),128.))
}
func ReverseComparΞfloat32(as, bs *[]float32, a, b int) int8 {
	return -ComparΞfloat32(as, bs, a, b)
}


func BestΞfloat32(s *[]float32, better func(as, bs *[]float32, a, b int) int8) (j int) {
	for i := range *s {
		if better(s,s,i,j) > 0 {
			j = i
		}
	}
	return j
}

// END OF BOILERPLATE


func main() {
	var a = []float32{7., 5., 7., 999., 99999., 3.}
	var b = []float32{3,3,3,3,3,3,3,999,3,3,3,3,0,3,3,}


	if(BestΞfloat32(&a, ComparΞfloat32) != 4) {
		panic("mismaxed + a\n")
	}

	if(BestΞfloat32(&b, ComparΞfloat32) != 7) {
		panic("mismaxed + b\n")
	}

	if(BestΞfloat32(&a, ReverseComparΞfloat32) != 5) {
		panic("mismaxed - a\n")
	}

	if(BestΞfloat32(&b, ReverseComparΞfloat32) != 12) {
		panic("mismaxed - b\n")
	}


}
