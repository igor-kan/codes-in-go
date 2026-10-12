package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6072 evaluates chebyshev collocation node order 6072.
func ComputeChebyshevColloc6072(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
