// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package perm_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/soniakeys/perm"
)

func ExampleSJTRecursive() {
	// Notice that first call to it() leaves p in its initial order.
	// Subsequent calls permute p.
	p := []int{11, 22, 33}
	it := perm.SJTRecursive(p)
	for it() {
		fmt.Println(p)
	}
	fmt.Println("Final contents:", p)
	fmt.Println("Another call returns", it())
	fmt.Println("Contents remain", p)
	// Output:
	// [11 22 33]
	// [11 33 22]
	// [33 11 22]
	// [33 22 11]
	// [22 33 11]
	// [22 11 33]
	// Final contents: [11 22 33]
	// Another call returns false
	// Contents remain [11 22 33]
}

func TestSJTRecursive(t *testing.T) {
	fact := 1
	for i := 0; i < 5; i++ {
		// accumulate factorial for checking number of permutations
		if i > 1 {
			fact *= i
		}
		p := perm.NewZPerm(i)      // create a slice to permute
		it := perm.SJTRecursive(p) // create iterator to test
		n := 0                     // count for checking number of permutations
		m := map[string]bool{}     // map for checking uniqueness

		// run through permutations
		for it() {
			n++                            // count all results
			m[fmt.Sprintf("%v", p)] = true // accumulate unique results
		}

		if n != fact {
			t.Log("i:", i, "n:", n, "fact:", fact)
			t.Fatal("missing permutations")
		}
		if len(m) != fact {
			t.Fatal("duplicate permutations")
		}
	}
}

func BenchmarkSJTR_5(b *testing.B) {
	bsjt(b, perm.NewZPerm(5))
}

func BenchmarkSJTR_10(b *testing.B) {
	bsjt(b, perm.NewZPerm(10))
}

func bsjt(b *testing.B, p []int) {
	it := perm.SJTRecursive(p)
	ok := true
	for i := 0; i < b.N; i++ {
		if !ok {
			it = perm.SJTRecursive(p)
		}
		ok = it()
	}
}

func ExampleLexNextInt() {
	p := perm.NewZPerm(3)
	fmt.Println(p)
	for perm.LexNextInt(p) {
		fmt.Println(p)
	}
	// Output:
	// [0 1 2]
	// [0 2 1]
	// [1 0 2]
	// [1 2 0]
	// [2 0 1]
	// [2 1 0]
}

func ExampleLexNextInt_multiset() {
	p := []int{0, 3, 3, 7}
	fmt.Println(p)
	for perm.LexNextInt(p) {
		fmt.Println(p)
	}
	// Output:
	// [0 3 3 7]
	// [0 3 7 3]
	// [0 7 3 3]
	// [3 0 3 7]
	// [3 0 7 3]
	// [3 3 0 7]
	// [3 3 7 0]
	// [3 7 0 3]
	// [3 7 3 0]
	// [7 0 3 3]
	// [7 3 0 3]
	// [7 3 3 0]
}

func ExampleLexNextSort() {
	p := sort.StringSlice{"blue", "green", "red"}
	fmt.Println(p)
	for perm.LexNextSort(p) {
		fmt.Println(p)
	}
	// Output:
	// [blue green red]
	// [blue red green]
	// [green blue red]
	// [green red blue]
	// [red blue green]
	// [red green blue]
}

func BenchmarkLexNext_5(b *testing.B) {
	ok := false
	var p []int
	for i := 0; i < b.N; i++ {
		if !ok {
			p = perm.NewZPerm(5)
		}
		ok = perm.LexNextInt(p)
	}
}

func BenchmarkLexNext_10(b *testing.B) {
	ok := false
	var p []int
	for i := 0; i < b.N; i++ {
		if !ok {
			p = perm.NewZPerm(10)
		}
		ok = perm.LexNextInt(p)
	}
}

func ExampleSJTE() {
	p, f := perm.SJTE(3)
	fmt.Println(p)
	for f() {
		fmt.Println(p)
	}
	fmt.Println("Next call to f() returns", f())
	fmt.Println("p:", p)
	// Output:
	// [0 1 2]
	// [0 2 1]
	// [2 0 1]
	// [2 1 0]
	// [1 2 0]
	// [1 0 2]
	// Next call to f() returns true
	// p: [0 2 1]
}

func BenchmarkSJTE_5(b *testing.B) {
	_, f := perm.SJTE(5)
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkSJTE_10(b *testing.B) {
	_, f := perm.SJTE(10)
	for i := 0; i < b.N; i++ {
		f()
	}
}
