package main

import "math/big"
/*
FIXME:ints.go:6: internal compiler error: walkexpr: switch 1 unknown op CALL l(6) tc(1)

func Compar(as, bs *[], a, b int) int8 {
	return int8((*as)[a].Cmp((*bs)[b]))
}
*/
// BEGIN OF BOILERPLATE
func ComparΞbigΞInt(as, bs *[]*big.Int, a, b int) int8 {
	return int8((*as)[a].Cmp((*bs)[b]))
}
func ComparΞbigΞRat(as, bs *[]*big.Rat, a, b int) int8 {
	return int8((*as)[a].Cmp((*bs)[b]))
}

// END OF BOILERPLATE

func main() {
	var a = []*big.Int{big.NewInt(7), big.NewInt(5), big.NewInt(7), big.NewInt(999)}
	var b = []*big.Rat{big.NewRat(7,1), big.NewRat(5,1), big.NewRat(7,1), big.NewRat(999,1)}

	if ComparΞbigΞInt(&a, &a, 0, 1) <= 0 {
		panic("miscompared a 0, 1\n")
	}

	if ComparΞbigΞInt(&a, &a, 0, 2) != 0 {
		panic("miscompared a 0, 2\n")
	}
	if ComparΞbigΞInt(&a, &a, 0, 3) >= 0 {
		panic("miscompared a 0, 3\n")
	}

	if ComparΞbigΞRat(&b, &b, 0, 1) <= 0 {
		panic("miscompared 0, 1\n")
	}

	if ComparΞbigΞRat(&b, &b, 0, 2) != 0 {
		panic("miscompared 0, 2\n")
	}
	if ComparΞbigΞRat(&b, &b, 0, 3) >= 0 {
		panic("miscompared 0, 3\n")
	}

}
