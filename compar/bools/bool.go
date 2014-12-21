package main

func Compar(as, bs *[], a, b int) (c int8) {
	if (*as)[a] {
		c++
	}
	if (*bs)[b] {
		c--
	}
	return c
}
// BEGIN OF BOILERPLATE
func ComparΞbool(as, bs *[]bool, a, b int) (c int8) {
	if (*as)[a] {
		c++
	}
	if (*bs)[b] {
		c--
	}
	return c
}

// END OF BOILERPLATE

func main() {
	var a = []bool{false, true}

	if ComparΞbool(&a, &a, 0, 0) != 0 {
		panic("miscompared a 0, 0\n")
	}
	if ComparΞbool(&a, &a, 1, 1) != 0 {
		panic("miscompared a 1, 1\n")
	}
	if ComparΞbool(&a, &a, 0, 1) >= 0 {
		panic("miscompared a 0, 1\n")
	}
	if ComparΞbool(&a, &a, 1, 0) <= 0 {
		panic("miscompared a 1, 0\n")
	}

}

