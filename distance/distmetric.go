package distance

import (
	"errors"
	"math"
)

type DistMetric func(x []float64, y []float64) (dist float64, err error)

func Binary(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	denominator := float64(0)
	numerator := float64(0)
	for i := range x {
		if x[i] > 0 && y[i] > 0 {
			numerator++
		}
		// Ignore i when both x[i] and y[i] are zero.
		if x[i] > 0 || y[i] > 0 {
			denominator++
		}
	}
	dist = 1 - (numerator / denominator)
	return
}

func Canberra(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	dist = 0
	for i := range x {
		// Ignore i when both x[i] and y[i] are zero.
		if x[i] > 0 || y[i] > 0 {
			dist += math.Abs(x[i]-y[i]) / (math.Abs(x[i]) + math.Abs(y[i]))
		}
	}
	return
}

func Jaccard(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	denominator := float64(0)
	numerator := float64(0)
	for i := range x {
		// Ignore i when both x[i] and y[i] are zero.
		if x[i] > 0 || y[i] > 0 {
			numerator += math.Min(x[i], y[i])
			denominator += math.Max(x[i], y[i])
		}
	}
	dist = 1 - (numerator / denominator)
	return
}

func Manhattan(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	dist = 0
	for i := range x {
		dist += math.Abs(x[i] - y[i])
	}
	return
}

func Maximum(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	dist = 0
	for i := range x {
		diff := math.Abs(x[i] - y[i])
		if diff > dist {
			dist = diff
		}
	}
	return
}

func Euclidean(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}
	dist = 0
	for i := range x {
		diff := x[i] - y[i]
		dist += diff * diff
	}
	dist = math.Sqrt(dist)
	return
}

func Cosine(x []float64, y []float64) (dist float64, err error) {
	if len(x) != len(y) {
		err = errors.New("Vectors for calculating distance must have equal length")
		return
	}

	var ab float64
	var a2 float64
	var b2 float64

	for i := range x {
		ab += x[i] * y[i]
		a2 += math.Pow(x[i], 2)
		b2 += math.Pow(y[i], 2)
	}

	denominator := math.Sqrt(a2) * math.Sqrt(b2)
	if denominator == 0 {
		err = errors.New("Zero vector detected")
		return
	}

	dist = 1 - (ab / denominator)
	return
}
