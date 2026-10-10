package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4112 evaluates chebyshev collocation node order 4112.
func ComputeChebyshevColloc4112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
