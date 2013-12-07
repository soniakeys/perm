## Perm

Three different algorithms for generating permutations, Lexicographic,
Steinhaus-Johnson-Trotter-Even, and a recursive algorithm.

All three are reasonably efficient and might be useful as they are but really
the package is a bit a dabbling at this point.

It started with a desire to generate permutations in SJT order.  I found the
Wikipedia explanation of the SJTE algorithim unclear and found a lot of
craziness when I looked for existing implementations on the internet.
Coding just to Wikipedia’s description of the "recursive structure," I came up
with the recursive algorithm in this package.  But then wanting to benchmark
my algorithm against others, I puzzled out the iterative SJTE and also the
lexicographic algorithms.  All three turn out to be very fast.

If you are interested in using them, note that their APIs are a little
different.  Read the doc and look at the examples.
