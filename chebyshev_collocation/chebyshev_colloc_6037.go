package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6037 evaluates chebyshev collocation node order 6037.
func ComputeChebyshevColloc6037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
