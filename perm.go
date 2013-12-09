// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package permute has functions to generate permutations.  And other related
// functions
package perm

import "math/big"

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

type Fact []int

func NewFact(x *big.Int, n int) Fact {
	f := make(Fact, n-1)
	f.Set(x)
	return f
}

func (f Fact) Set(n *big.Int) {
	var x, y, r big.Int
	x.Set(n)
	for i := int64(2); x.Sign() == 1; i++ {
		x.QuoRem(&x, y.SetInt64(i), &r)
		f[i-2] = int(r.Int64())
	}
}

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

func (f Fact) Perm(r []int) []int {
	p := NewZPerm(len(f) + 1)
	for i := len(f) - 1; i >= 0; i-- {
		d := f[i]
		r = append(r, p[d])
		copy(p[d:], p[d+1:])
	}
	return append(r, p[0])
}
