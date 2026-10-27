package continued_fractions

// ComputeContinuedFrac9110 evaluates continued fraction approximant order 9110.
func ComputeContinuedFrac9110(x float64) float64 {
	a := 1.0
	for k := 3; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
