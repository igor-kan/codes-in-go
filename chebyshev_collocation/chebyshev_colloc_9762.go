package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9762 evaluates chebyshev collocation node order 9762.
func ComputeChebyshevColloc9762(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
