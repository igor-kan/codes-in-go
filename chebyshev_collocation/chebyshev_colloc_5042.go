package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5042 evaluates chebyshev collocation node order 5042.
func ComputeChebyshevColloc5042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
