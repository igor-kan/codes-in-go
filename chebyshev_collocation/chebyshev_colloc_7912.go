package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7912 evaluates chebyshev collocation node order 7912.
func ComputeChebyshevColloc7912(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
