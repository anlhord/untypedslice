package main

func test1(i [], j []) int {
	return 1
}

func test2(i,j []) int {
	return 2
}

type foo struct {
	Val []
}

const x = 3
const y = 3

// FIXME
const z = []{1,2,3}

func main() {

	var a int		
	var b byte		
	var c []		
	var d []byte		
	var e []		

	var f []		
	var g []byte		
	var h []	

// FIXME
//	i := []{1,2,3}

	_ = a
	_ = b
	_ = c
	_ = d
	_ = e
	_ = f
	_ = g
	_ = h
}
