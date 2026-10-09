package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3062 evaluates chebyshev collocation node order 3062.
func ComputeChebyshevColloc3062(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
