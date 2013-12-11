// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package permute has functions related to permutations.
package perm

import (
	"math/big"
	"strconv"
)

// ZPerm represents a permutation of integers starting with 0.
//
// A ZPerm of length n will contain exactly the integers from 0 to n-1,
// with no omissions and no repeats.
type ZPerm []int

// NewZPerm returns a new ZPerm with integers in order from 0 to n-1.
func NewZPerm(n int) ZPerm {
	p := make(ZPerm, n)
	for i := range p {
		p[i] = i
	}
	return p
}

// Inverse returns the inverse of a permutation.
//
// It returns the permutation where eaach element is its position in
// the original permutation p.
func (p ZPerm) Inverse() ZPerm {
	v := make(ZPerm, len(p))
	for i, x := range p {
		v[x] = i
	}
	return v
}

// Fact represents a non-negative integer as a "factoradic" or Lehmer code.
//
// Slice index i corresponds to place value (i+1)! and the i-th element
// must be in the range 0..(i+2).  For a Fact of length n, the largest
// representable integer is (n+1)!-1.
type Fact []int

// String formats f with the most significant digit first and with a trailing
// zero (for the informationless and unrepresented 0! place value.)
//
// If any digit is out of range of the place value, "invalid" appears after
// the digit list.
func (f Fact) String() string {
	if len(f) == 0 {
		return "Fact(0)"
	}
	s := "0)"
	invalid := false
	for i, d := range f {
		s = strconv.Itoa(d) + " " + s
		if d < 0 || d > i+1 {
			invalid = true
		}
	}
	s = "Fact(" + s
	if invalid {
		s += "invalid"
	}
	return s
}

// NewFact creates a new Fact and initializes it to x.
//
// The Fact object is created of length to allow numbers up to n!, the number
// of permutations of n items.
//
// NewFact returns false if n < 0, x < 0, or x >= n!.
func NewFact(x *big.Int, n int) (Fact, bool) {
	if n <= 1 {
		return nil, n >= 0
	}
	f := make(Fact, n-1)
	return f, f.Set(x)
}

// Set sets the value of f to n.
//
// Set returns false if n cannot be represented, that is, if n is < 0 or
// n >= (len(f)+1)!.
func (f Fact) Set(n *big.Int) bool {
	var x, y, r big.Int
	x.Set(n)
	var i int
	for ; i < len(f) && x.Sign() == 1; i++ {
		x.QuoRem(&x, y.SetInt64(int64(i+2)), &r)
		f[i] = int(r.Int64())
	}
	for ; i < len(f); i++ {
		f[i] = 0
	}
	return x.Sign() == 0
}

// Int returns the value of f as a big.Int.
func (f Fact) Int() *big.Int {
	s := &big.Int{}
	if len(f) == 0 {
		return s
	}
	var m big.Int
	s.SetInt64(int64(f[len(f)-1]))
	for i := len(f); i > 1; {
		s.Mul(s, m.SetInt64(int64(i)))
		i--
		s.Add(s, m.SetInt64(int64(f[i-1])))
	}
	return s
}

func (f Fact) Perm() ZPerm {
	p := make(ZPerm, len(f) + 1)
	setzf(p, f)
	return p
}

func (p ZPerm) SetFact(f Fact) bool {
	if len(p) != len(f) + 1 {
		return false
	}
	setzf(p, f)
	return true
}

// set p to f.  panics if p not big enough.
func setzf(p ZPerm, f Fact) {
	for i := range p {
		p[i] = i
	}
	last := len(f)-1
	for i := range f {
		dx := i + f[last-i]
		item := p[dx]
		copy(p[i+1:dx+1], p[i:dx])
		p[i] = item
	}
}
