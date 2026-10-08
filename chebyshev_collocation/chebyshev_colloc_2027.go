package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2027 evaluates chebyshev collocation node order 2027.
func ComputeChebyshevColloc2027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
