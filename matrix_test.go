package matrix

import (
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
	p, err := Product[Matrix](&m, &n)
	pc := Matrix{
		{29, 12},
		{56, 41},
		{12, 8},
	}

	if !reflect.DeepEqual(p, pc) || err != nil {
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
	q, err := Product[Matrix](&r, &s)
	if err == nil || q != nil {
		t.Error("Matrix product dimenstions check fail!")
	}

	v := Vector{100, 7, 9}
	k, err := Product[Vector](&m, &v)
	kc := Vector{141, 853, 418}

	if !reflect.DeepEqual(k, kc) || err != nil {
		t.Error("Matrix vector product fail!")
	}

	h, err := Product[Vector](&s, &v)
	if err == nil || h != nil {
		t.Error("Matrix vector product dimenstions check fail!")
	}

	var w int = 2
	d, err := Product[Matrix](&s, &w)

	if err == nil || d != nil {
		t.Error("Matrix product type check failed!")
	}
}
