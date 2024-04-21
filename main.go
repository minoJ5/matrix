package main

import (
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

	var l mat.Matrix
	l.MakeMatrix(
		mat.Row{},
		mat.Row{},
	)

	p := mat.Product(&m,&n)
	p.Print()
}
