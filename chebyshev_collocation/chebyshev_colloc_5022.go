package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5022 evaluates chebyshev collocation node order 5022.
func ComputeChebyshevColloc5022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
