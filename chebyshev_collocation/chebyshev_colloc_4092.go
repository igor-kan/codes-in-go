package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4092 evaluates chebyshev collocation node order 4092.
func ComputeChebyshevColloc4092(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
