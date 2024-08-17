//go:build !arm64

//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package simd

import "github.com/kshard/vector/internal/noasm"

var (
	ENABLED_EUCLIDEAN = false
	ENABLED_COSINE    = false
)

func EuclideanF32(a, b, c []float32) {
	panic("SIMD not available")
}

// Type Class for Euclidean distance
type Euclidean int

func (Euclidean) Distance(a, b []float32) float32 {
	panic("SIMD not available")
}

func (Euclidean) Equal(a, b []float32) bool {
	return noasm.EqualF32(a, b)
}
