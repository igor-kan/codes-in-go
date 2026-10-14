package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7202 evaluates chebyshev collocation node order 7202.
func ComputeChebyshevColloc7202(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
