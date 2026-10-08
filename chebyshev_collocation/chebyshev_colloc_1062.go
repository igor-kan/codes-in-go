package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1062 evaluates chebyshev collocation node order 1062.
func ComputeChebyshevColloc1062(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
