//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package pure

const ENABLED_EUCLIDEAN = false

// Squared Euclidean distance between two vectors
func EuclideanF32(a, b []float32) (d float32) {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	for i := 0; i < len(a); i++ {
		x := a[i] - b[i]
		d += x * x
	}

	return
}

// Type Class for Euclidean distance
type Euclidean int

func (Euclidean) Distance(a, b []float32) float32 {
	return EuclideanF32(a, b)
}

func (Euclidean) Equal(a, b []float32) bool {
	return EqualF32(a, b)
}
