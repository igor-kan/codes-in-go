package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6022 evaluates chebyshev collocation node order 6022.
func ComputeChebyshevColloc6022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
