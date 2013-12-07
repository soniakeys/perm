// Copyright 2013 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package permute_test

import (
	"fmt"
	"testing"

	"github.com/soniakeys/permute"
)

func ExampleSJT() {
	p := []int{11, 22, 33}
	it := permute.SJT(p)
	for it() {
		fmt.Println(p)
	}
	fmt.Println("final contents:", p)
	// Output:
	// [11 22 33]
	// [11 33 22]
	// [33 11 22]
	// [33 22 11]
	// [22 33 11]
	// [22 11 33]
	// final contents: [11 22 33]
}

func TestSJT(t *testing.T) {
	fact := 1
	for i := 0; i < 5; i++ {
		// accumulate factorial for checking number of permutations
		if i > 1 {
			fact *= i
		}
		p := permute.Ints(i)   // create a slice to permute
		it := permute.SJT(p)   // create iterator to test
		n := 0                 // count for checking number of permutations
		m := map[string]bool{} // map for checking uniqueness

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

func BenchmarkSJT_5(b *testing.B) {
	bsjt(b, permute.Ints(5))
}

func BenchmarkSJT_10(b *testing.B) {
	bsjt(b, permute.Ints(10))
}

func bsjt(b *testing.B, p []int) {
	it := permute.SJT(p)
	ok := true
	for i := 0; i < b.N; i++ {
		if !ok {
			it = permute.SJT(p)
		}
		ok = it()
	}
}

func ExampleLexNext() {
	p := permute.Ints(3)
	fmt.Println(p)
	for permute.LexNext(p) {
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

func BenchmarkLexNext_5(b *testing.B) {
    ok := false
    var p []int
    for i := 0; i < b.N; i++ {
        if !ok {
            p = permute.Ints(5)
        }
        ok = permute.LexNext(p)
    }
}

func BenchmarkLexNext_10(b *testing.B) {
    ok := false
    var p []int
    for i := 0; i < b.N; i++ {
        if !ok {
            p = permute.Ints(10)
        }
        ok = permute.LexNext(p)
    }
}

func ExampleSJTE() {
	p, f := permute.SJTE(3)
	fmt.Println(p)
	for f() {
		fmt.Println(p)
	}
	// Output:
	// [0 1 2]
	// [0 2 1]
	// [2 0 1]
	// [2 1 0]
	// [1 2 0]
	// [1 0 2]
}

func BenchmarkSJTE_5(b *testing.B) {
	_, f := permute.SJTE(5)
    for i := 0; i < b.N; i++ {
		f()
    }
}


func BenchmarkSJTE_10(b *testing.B) {
	_, f := permute.SJTE(10)
    for i := 0; i < b.N; i++ {
		f()
    }
}

