package main
/*
tokey.go:8: internal compiler error: fault

// Key produces a Key for the slice i-th element
func Key(s *[], i int) float64 {
	if i > len(*s) || i < 0 {
		return 9999.
	}
	return float64((*s)[i])
}
*/
// Keys produces a Key for the first, last element of a sorted slice
func Keys(s *[], k func (s *[], i int) (float64)) (float64, float64) {
	return k(s, 0), k(s, len(*s)-1)
}

// BEGIN OF BOILERPLATE

func KeyΞfloat64(s *[]float64, i int) float64 {
	if i >= len(*s) || i < 0 {
		return 1.e99
	}
	return float64((*s)[i])
}


func KeysΞfloat64(s *[]float64, k func (s *[]float64, i int) (float64)) (float64, float64) {
	return k(s, 0), k(s, len(*s)-1)
}

// END OF BOILERPLATE


func main() {
	var a = []float64{}
	var b = []float64{7,8,13,25,66,1337}

	p, q := KeysΞfloat64(&a, KeyΞfloat64)
	r, s := KeysΞfloat64(&b, KeyΞfloat64)

	if(p != 1.e99 || q != 1.e99) {
		panic("keys empty\n")
	}
	if(r != 7. || s != 1337) {
		panic("keys nonempty\n")
	}
}
