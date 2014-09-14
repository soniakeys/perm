package perm

// An Alphabet is an ordered set of symbols.
type Alphabet []rune

// Count computes permutations with repetitions.
//
// Count calls c for each permutation with repetitions of length n
// of alphabet a.  That is, it counts in base len(a), enumerating
// len(a)**n words of length n.
//
// The element at index 0 of slice 'word' is most significant;
// it varies slowest.  Thus, words are returned in lexical order with respect
// to alphabet a.
//
// The slice 'word' passed to callback c is reused on each call.  You must copy
// it if you need to keep the value after c returns.
//
// Callback c should normally return true.
// If c returns false, Count terminates immediately.
func (a Alphabet) Count(n int, c func(word []rune) (more bool)) {
	w := make([]rune, n)
	var r func(int) bool
	r = func(p int) bool {
		if p == n {
			return c(w)
		}
		for _, sym := range a {
			w[p] = sym
			if !r(p + 1) {
				return false
			}
		}
		return true
	}
	r(0)
}
