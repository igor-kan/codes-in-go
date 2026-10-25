package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8872 evaluates chebyshev collocation node order 8872.
func ComputeChebyshevColloc8872(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
