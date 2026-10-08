package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1002 evaluates chebyshev collocation node order 1002.
func ComputeChebyshevColloc1002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
