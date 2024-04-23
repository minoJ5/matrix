package matrix

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
)

const (
	MatrixEmpty = iota + 1
	MissingRows
	MissingRowData
	MissingRowsAndData
	DataIsGood
)

type linearObject interface {
	Matrix | Vector
}

// Alias for a 1 dimensional array representing a vector
type Row []float64

// Semantics
type Vector Row

// Alias for a 2 dimentional array representing a matrix
type Matrix []Row

func (m *Matrix) allocateMatrix(r, c int) {
	*m = make(Matrix, r)
	for i := range *m {
		(*m)[i] = make(Row, c)
	}
}

// Fill the matrix with data from rows.
// Set to nil if input is invalid, i. e., the matrix is empty.
// Otherwise fill the gaps with zeros.
func (m *Matrix) MakeMatrix(rows ...Row) {
	if rows == nil {
		*m = nil
	} else {
		*m = make(Matrix, len(rows))
		for r, row := range rows {
			(*m)[r] = make(Row, len(row))
			copy((*m)[r], row)
		}
	}
	q := checkIntegrity(*m)
	if q == MatrixEmpty {
		log.Printf("Critical! Cannot make matrix. Matrix is empty, check input!\n")
		*m = nil
		return
	} else if q == MissingRowData || q == MissingRows || q == MissingRowsAndData {
		m.fixMatrix()
	}
}

// Print the matrix in the console with tabs as spacers.
// If the matrix is nill log nil.
func (m *Matrix) Print() {
	if *m == nil {
		log.Printf("[nil]")
		return
	}
	var s string
	for r := 0; r < len(*m); r++ {
		for c := 0; c < len((*m)[r]); c++ {
			s += fmt.Sprintf("%.2f\t", (*m)[r][c])
		}
		s += "\n"
	}
	fmt.Printf("%s", s)
}

func getMaxRowLength(m Matrix) int {
	var ml int
	for _, row := range m {
		if len(row) > ml {
			ml = len(row)
		}
	}
	return ml
}

// Check for empty rows or values.
func checkIntegrity(m Matrix) int {
	if m == nil {
		return MatrixEmpty
	}
	var ml int = getMaxRowLength(m)
	if ml == 0 {
		return MatrixEmpty
	}
	var ce int // Counter empty rows
	var fm int // flag missing data in row
	for _, row := range m {
		if len(row) < 1 {
			ce++
		}
	}
	for _, row := range m {
		if len(row) < ml && len(row) > 0 {
			fm++
			break
		}
	}
	switch ce + fm {
	case 0:
		return DataIsGood
	case ce:
		return MissingRows
	case fm:
		return MissingRowData
	default:
		return MissingRowsAndData
	}
}

// If data is missing and the matrix is not empty
// fill the missing entries with zeros.
func (m *Matrix) fixMatrix() {
	var wg sync.WaitGroup
	var ml = getMaxRowLength(*m)
	for r := range *m {
		wg.Add(1)
		go func(r int) {
			defer wg.Done()
			if len((*m)[r]) < ml {
				for v := len((*m)[r]); v < ml; v++ {
					(*m)[r] = append((*m)[r], 0)
				}
			}
		}(r)
	}
	wg.Wait()
}

// Calculate the product of 2 matrices concurrently.
// Retrun nil if dimensions are not valid.
func ProductMM(a, b *Matrix) (Matrix, error) {
	if len((*a)[0]) != len(*b) {
		err := fmt.Sprintf("critical invalid dimensions!\n"+
			"number of colums of the first matrix (%d) has "+
			"to be equal to the number of rows of the second matrix (%d)\n", len((*a)[0]), len(*b))
		return nil, errors.New(err)
	}
	var p Matrix
	var wg sync.WaitGroup
	p.allocateMatrix(len((*a)), len((*b)[0]))
	for r := range *a {
		for c := 0; c < len((*b)[0]); c++ {
			wg.Add(1)
			go func(r, c int) {
				defer wg.Done()
				for k := 0; k < len(*b); k++ {
					p[r][c] += (*a)[r][k] * (*b)[k][c]
				}
			}(r, c)
		}
	}
	wg.Wait()
	return p, nil
}

// Calculate the product of a matrix with a vector concurrently.
// Retrun nil if dimensions are not valid.
func productVM(a *Matrix, v *Vector) (Vector, error) {
	if len((*a)[0]) != len(*v) {
		err := fmt.Sprintf("critical invalid dimensions!\n"+
			"number of colums of the first matrix (%d) has "+
			"to be equal to the dimenston of the vector (%d)\n", len((*a)[0]), len(*v))
		return nil, errors.New(err)
	}
	p := make(Vector, len(*v))
	var wg sync.WaitGroup
	for r := range *a {
		for k := 0; k < len(*v); k++ {
			wg.Add(1)
			go func(r, k int) {
				defer wg.Done()
				p[r] += (*a)[r][k] * (*v)[k]
			}(r, k)
		}
	}
	wg.Wait()
	return p, nil
}

// Wrapper around the multiplications of matrix matrix or matrix vector
func Product[T linearObject](a *Matrix, b interface{}) (T, error) {
	var p interface{}
	var err error
	switch bt := b.(type) {
	case *Matrix:
		p, err = ProductMM(a, bt)
	case *Vector:
		p, err = productVM(a, bt)
	default:
		err = fmt.Errorf("invalid second argument type! expected matrix or vector. reieved: %v", reflect.TypeOf(b))
		return nil, err
	}
	return p.(T), err
}
