// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package perm

import (
	"math/big"
)

// MRPerm creates the permutation of n integers with rank i relative to
// the order described by Myrvold and Ruskey in "Ranking and Unranking
// Permutations in Linear Time."
//
// While M&R claim time complexity of O(n), that would only seem true with
// constant time arithmetic on i.  Allowing time for big.Int arithmetic,
// I suspect time complexity is O(n log n).
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

// MRRank returns rank of p relative to the order described by Myrvold and
// Ruskey in "Ranking and Unranking Permutations in Linear Time."
//
// While M&R claim time complexity of O(n), that would only seem true with
// constant time arithmetic on i.  Allowing time for big.Int arithmetic,
// I suspect time complexity is O(n log n).
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
