// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package permute has functions to generate permutations.  And other related
// functions
package perm

import (
	"math/big"
)

func MRPerm(n int, i *big.Int) ZPerm {
	p := NewZPerm(n)
	var quo, rem, b big.Int
	for quo.Set(i); n > 0; {
		quo.QuoRem(&quo, b.SetInt64(int64(n)), &rem)
		n--
		m := rem.Int64()
		p[n], p[m] = p[m], p[n]
	}
	return p
}

func (p ZPerm) MRRank() *big.Int {
	p = append(ZPerm{}, p...)
	inv := p.Inverse()
	for i := len(p) - 1; i > 0; i-- {
		s := p[i]
		p[inv[i]] = s
		inv[s] = inv[i]
	}
	var r, b big.Int
	for i := 1; i < len(p); i++ {
		r.Mul(&r, b.SetInt64(int64(i+1)))
		r.Add(&r, b.SetInt64(int64(p[i])))
	}
	return &r
}
