// Package perm contains various permutation generators.
//
// Permutation generators currently include lexicographic,
// Steinhaus-Johnson-Trotter, and Myrvold-Ruskey.
//
// The type ZPerm is defined for holding permutations of sequential integers
// starting with zero.
//
// Also there are methods for ranking permutations.
//
// Lexicographic order
//
// Functions:
//
//   LexNextInt    generates successive permutations of integer sets or multisets.
//   LexNextSort   similar, but for any type satisfying sort.Interface.
//   LexPerm       generates a ZPerm for a given rank.
//   ZPerm.LexRank computes the lexicographic rank for a ZPerm.
//
// Related, type Fact represents the Factoradic, or Lehmer Code of a
// number, which corresponds to the lexicographic rank of a permutation.
//
// Related to lexicographic order anyway, Alphabet.Count "counts" in a given
// alphabet of digits, thus generating permutations with repetitions in
// lexicographic order.
//
// Steinhaus-Johnson-Trotter order
//
// Also called "plain changes" order:
//
//   SJTRecursive permutes an integer set or multiset.
//   SJTE         permutes a ZPerm using an iterative algorithm.
//
// Myrvold-Ruskey order
//
// MR order allows fast ranking and unranking of ZPerms.
//
//   MRPerm
//   ZPerm.MRRank
package perm
