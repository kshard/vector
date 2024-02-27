//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package noasm

import "github.com/chewxy/math32"

const ENABLED_COSINE = true

// Cosine distance between two vectors
func CosineF32(a, b []float32) (d float32) {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	if len(a)%4 != 0 {
		panic("vector length must be multiple of 4")
	}

	ab := float32(0.0)
	aa := float32(0.0)
	bb := float32(0.0)

	for i := 0; i < len(a); i += 4 {
		asl := a[i : i+4 : i+4]
		bsl := b[i : i+4 : i+4]

		ax0, ax1, ax2, ax3 := asl[0], asl[1], asl[2], asl[3]
		bx0, bx1, bx2, bx3 := bsl[0], bsl[1], bsl[2], bsl[3]

		ab0 := ax0 * bx0
		ab1 := ax1 * bx1
		ab2 := ax2 * bx2
		ab3 := ax3 * bx3
		ab += ab0 + ab1 + ab2 + ab3

		aa0 := ax0 * ax0
		aa1 := ax1 * ax1
		aa2 := ax2 * ax2
		aa3 := ax3 * ax3
		aa += aa0 + aa1 + aa2 + aa3

		bb0 := bx0 * bx0
		bb1 := bx1 * bx1
		bb2 := bx2 * bx2
		bb3 := bx3 * bx3
		bb += bb0 + bb1 + bb2 + bb3
	}

	s := math32.Sqrt(aa) * math32.Sqrt(bb)

	d = (1 - ab/s) / 2

	return
}

// Type Class for Cosine distance
type Cosine int

func (Cosine) Distance(a, b []float32) float32 {
	return CosineF32(a, b)
}
