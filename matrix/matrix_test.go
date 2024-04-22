package matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	var m Matrix
	m.MakeMatrix()
	if m != nil {
		t.Error("Empty matrix should be nil but: ", m)
	}
	var n Matrix
	n.MakeMatrix(Row{}, Row{})
	if n != nil {
		t.Error("Empty matrix should be nil but: ", m)
	}

	var l Matrix
	l.MakeMatrix(
		Row{1, 2},
		Row{8, 5},
	)
	lc := Matrix{
		{1, 2},
		{8, 5},
	}
	if !reflect.DeepEqual(lc, l) {
		t.Error("Create matrix failed!")
	}

	var o Matrix
	o.MakeMatrix(
		Row{1},
		Row{8, 5, 9},
	)
	oc := Matrix{
		{1, 0, 0},
		{8, 5, 9},
	}
	if !reflect.DeepEqual(oc, o) {
		t.Error("Create matrix failed: Fix Matrix failed!")
	}

	var p Matrix
	p.MakeMatrix(
		Row{},
		Row{8, 5, 9},
	)
	pc := Matrix{
		{0, 0, 0},
		{8, 5, 9},
	}
	if !reflect.DeepEqual(pc, p) {
		t.Error("Create matrix failed: Fix Matrix failed!")
	}
}

func TestProduct(t *testing.T) {
	var m Matrix
	m.MakeMatrix(
		Row{1, 2, 3},
		Row{8, 5, 2},
		Row{4, 0, 2},
	)

	var n Matrix
	n.MakeMatrix(
		Row{1, 2},
		Row{8, 5},
		Row{4, 0},
	)
	p, err := Product(&m,&n)
	pc := Matrix{
		{29, 12},
		{56, 41},
		{12, 8},
	}

	if !reflect.DeepEqual(p, pc) || err != nil{
		t.Error("Matrix product fail!")
	}

	var r Matrix
	r.MakeMatrix(
		Row{1, 2},
		Row{8, 5},
		Row{4, 0},
	)

	var s Matrix
	s.MakeMatrix(
		Row{1, 2},
		Row{8, 5},
		Row{4, 0},
	)
	q, err := Product(&r,&s)
	fmt.Println(q)
	if err == nil || q != nil {
		t.Error("Matrix product dimenstions check fail!")
	}

}
