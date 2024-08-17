//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package pure

import (
	"github.com/chewxy/math32"
)

const ENABLED_COSINE = false

// Cosine distance between two vectors
func CosineF32(a, b []float32) (d float32) {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	ab := float32(0.0)
	aa := float32(0.0)
	bb := float32(0.0)

	for i := 0; i < len(a); i++ {
		ab += a[i] * b[i]
		aa += a[i] * a[i]
		bb += b[i] * b[i]
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

func (Cosine) Equal(a, b []float32) bool {
	return EqualF32(a, b)
}
