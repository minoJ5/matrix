package main

import (
	"fmt"
	mat "testmatrix/matrix"
)

func main() {

	var m mat.Matrix
	m.MakeMatrix(
		mat.Row{1, 2, 3},
		mat.Row{8, 5, 2},
		mat.Row{4, 0, 2},
	)
	//m.Print()
	var n mat.Matrix
	n.MakeMatrix(
		mat.Row{1, 2},
		mat.Row{8, 5},
		mat.Row{4, 0},
	)
	v := mat.Vector{100, 7, 9}
	p, err := mat.Product[mat.Vector](&m, &v)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("p: %v\n", p)

	var w int = 2
	d, err := mat.Product[mat.Matrix](&m, &w)

	fmt.Println(d == nil, err)

}
