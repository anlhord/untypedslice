package main

/*
FIXME: internal compiler error: fault

func Compar(as, bs *[], a, b int) int8 {
	if (*as)[a] < (*bs)[b] {
		return -1
	}
	return int8((*as)[a] - (*bs)[b])
}*/
// BEGIN OF BOILERPLATE
func ComparΞint(as, bs *[]int, a, b int) int8 {
	if (*as)[a] < (*bs)[b] {
		return -1
	}
	return int8((*as)[a] - (*bs)[b])
}
func ComparΞuint(as, bs *[]uint, a, b int) int8 {
	if (*as)[a] < (*bs)[b] {
		return -1
	}
	return int8((*as)[a] - (*bs)[b])
}

// END OF BOILERPLATE

func main() {
	var a = []int{7, 5, 7, 999}
	var b = []uint{7, 5, 7, 999}

	if ComparΞint(&a, &a, 0, 1) <= 0 {
		panic("miscompared a 0, 1\n")
	}

	if ComparΞint(&a, &a, 0, 2) != 0 {
		panic("miscompared a 0, 2\n")
	}
	if ComparΞint(&a, &a, 0, 3) >= 0 {
		panic("miscompared a 0, 3\n")
	}

	if ComparΞuint(&b, &b, 0, 1) <= 0 {
		panic("miscompared 0, 1\n")
	}

	if ComparΞuint(&b, &b, 0, 2) != 0 {
		panic("miscompared 0, 2\n")
	}
	if ComparΞuint(&b, &b, 0, 3) >= 0 {
		panic("miscompared 0, 3\n")
	}

}
