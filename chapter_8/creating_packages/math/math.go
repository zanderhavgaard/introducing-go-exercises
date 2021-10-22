package math

// function names that start with a capitalized letter are exposed
// outside of the module, funcitons that start with a lowercase
// letter are internal to the module
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
