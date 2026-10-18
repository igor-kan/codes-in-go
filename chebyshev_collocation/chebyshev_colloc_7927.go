package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7927 evaluates chebyshev collocation node order 7927.
func ComputeChebyshevColloc7927(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
