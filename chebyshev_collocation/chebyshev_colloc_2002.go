package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2002 evaluates chebyshev collocation node order 2002.
func ComputeChebyshevColloc2002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
