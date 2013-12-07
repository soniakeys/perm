// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

// Package permute generates permutations.
package permute

//import "fmt"

// Ints returns a slice of ints containing sequential integers 0..n-1
//
// Simply a utility function.
func Ints(n int) []int {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return p
}

// SJT takes a slice p and returns an iterator function.  The iterator
// permutes p in place and returns true for each permutation.  After all
// permutations have been generated, the iterator returns false and p is left
// in its initial order.
func SJT(p []int) func() bool {
	f := sjt(len(p))
	return func() bool {
		return f(p)
	}
}

// Recursive function used by perm, returns a chain of closures that
// implement a loopless recursive SJT.
func sjt(n int) func([]int) bool {
	perm := true
	switch n {
	case 0, 1:
		return func([]int) (r bool) {
			r = perm
			perm = false
			return
		}
	default:
		p0 := sjt(n - 1)
		i := n
		var d int
		return func(p []int) bool {
			switch {
			case !perm:
			case i == n:
				i--
				perm = p0(p[:i])
				d = -1
			case i == 0:
				i++
				perm = p0(p[1:])
				d = 1
				if !perm {
					p[0], p[1] = p[1], p[0]
				}
			default:
				p[i], p[i-1] = p[i-1], p[i]
				i += d
			}
			return perm
		}
	}
}

func LexNext(a []int) bool {
	if len(a) <= 1 {
		return false
	}
	last := len(a) - 1
	k := last - 1
	for ; a[k] >= a[k+1]; k-- {
		if k == 0 {
			return false
		}
	}
	l := last
	for ; a[k] >= a[l]; l-- {
	}
	a[k], a[l] = a[l], a[k]
	for l, r := k+1, last; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return true
}

func SJTE(n int) ([]int, func() bool) {
	p := make([]int, n+2)
	d := make([]int, n+2)
	p[0] = n
	for i := range p[1:] {
		p[i+1] = i
		d[i] = -1
	}
	return p[1:n+1:n+1], func() bool {
		var k int
		max := -1
		for i := 1; i <= n; i++ {
			if v := p[i]; v > max && v > p[i+d[i]] {
				max = v
				k = i
			}
		}
		if k == 0 {
			p[1], p[2] = 0, 1
			for i := 3; i <= n; i++ {
				d[i] = -1
			}
			return false
		}
		nx := k+d[k]
		p[k], p[nx] = p[nx], max
		d[k], d[nx] = d[nx], d[k]
		for i := 1; i <= n; i++ {
			if p[i] > max {
				d[i] = -d[i]
			}
		}
		return true
	}
}
