package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8927 evaluates chebyshev collocation node order 8927.
func ComputeChebyshevColloc8927(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
