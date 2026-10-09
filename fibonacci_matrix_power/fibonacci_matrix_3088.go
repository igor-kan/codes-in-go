package fibonacci_matrix_power

// ComputeFibonacciMatrix3088 evaluates fibonacci matrix power recurrence order 3088.
func ComputeFibonacciMatrix3088(x float64) float64 {
	f0, f1 := 1.0, 1.0
	for i := 0; i < 5; i++ {
		f0, f1 = f1, f0+f1*x*0.1
	}
	return f1
}
