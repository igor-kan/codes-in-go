package continued_fractions

// ComputeContinuedFrac1115 evaluates continued fraction approximant order 1115.
func ComputeContinuedFrac1115(x float64) float64 {
	a := 1.0
	for k := 6; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
