package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7237 evaluates chebyshev collocation node order 7237.
func ComputeChebyshevColloc7237(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
