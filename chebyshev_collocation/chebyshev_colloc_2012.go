package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2012 evaluates chebyshev collocation node order 2012.
func ComputeChebyshevColloc2012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
