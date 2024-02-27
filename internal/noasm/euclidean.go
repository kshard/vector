//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package noasm

const ENABLED_EUCLIDEAN = true

// Squared Euclidean distance between two vectors
func EuclideanF32(a, b []float32) (d float32) {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	if len(a)%4 != 0 {
		panic("vectors length must be multiple of 4")
	}

	for i := 0; i < len(a); i += 4 {
		av := a[i : i+4 : i+4]
		bv := b[i : i+4 : i+4]
		x1 := av[0] - bv[0]
		x2 := av[1] - bv[1]
		x3 := av[2] - bv[2]
		x4 := av[3] - bv[3]

		d1 := x1 * x1
		d2 := x2 * x2
		d3 := x3 * x3
		d4 := x4 * x4

		d += d1 + d2 + d3 + d4
	}

	return
}

// Type Class for Euclidean distance
type Euclidean int

func (Euclidean) Distance(a, b []float32) float32 {
	return EuclideanF32(a, b)
}
