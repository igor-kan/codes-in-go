package continued_fractions

// ComputeContinuedFrac1070 evaluates continued fraction approximant order 1070.
func ComputeContinuedFrac1070(x float64) float64 {
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
