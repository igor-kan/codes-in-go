package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7762 evaluates chebyshev collocation node order 7762.
func ComputeChebyshevColloc7762(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
