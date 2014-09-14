package perm_test

import (
	"fmt"

	"github.com/soniakeys/perm"
)

func ExampleAlphabet_Count() {
	a := perm.Alphabet("01")
	a.Count(3, func(w []rune) bool {
		s := string(w)
		fmt.Println(s)
		return s != "011"
	})
	// Output:
	// 000
	// 001
	// 010
	// 011
}
