package pythagorean

// https://www.mathsisfun.com/numbers/pythagorean-triples.html
// shows how to construt a pythagorean triple
// When m and n are any two positive integers (m < n):
// a = n^2 - m^2
// b = 2nm
// c = n^2 + m^2
// using this method can construct one such triple, but does not mean
// it can find all triples

// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var ret []Triplet

	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if isPythagoreanTriplet(i, j, k) {
					ret = append(ret, Triplet{i, j, k})
				}

			}
		}
	}
	return ret

}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	// a<=b<=c, a+b > c, a + b+ c = p, => p - c > c, p > 2c>=2b>=2c
	var ret []Triplet
	for i := 1; i <= p/2; i++ {
		for j := i; j <= p/2; j++ {
			k := p - i - j
			if k >= j && isPythagoreanTriplet(i, j, k) {
				ret = append(ret, Triplet{i, j, k})
			}
		}

	}

	return ret

}

// based on the above method, m<n, return a triplet
func constructPythagoreanTriplet(m, n int) Triplet {
	var t Triplet
	t[0] = n*n - m*m
	t[1] = 2 * n * m
	t[2] = n*n + m*m
	if t[0] > t[1] {
		t[0], t[1] = t[1], t[0]
	}
	return t
}

// a<=b<=c
func isPythagoreanTriplet(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}
