package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6002 evaluates chebyshev collocation node order 6002.
func ComputeChebyshevColloc6002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
