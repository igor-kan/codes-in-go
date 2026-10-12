package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6027 evaluates chebyshev collocation node order 6027.
func ComputeChebyshevColloc6027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
