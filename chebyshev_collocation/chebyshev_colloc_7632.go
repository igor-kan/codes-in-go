package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7632 evaluates chebyshev collocation node order 7632.
func ComputeChebyshevColloc7632(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
