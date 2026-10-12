package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6017 evaluates chebyshev collocation node order 6017.
func ComputeChebyshevColloc6017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
