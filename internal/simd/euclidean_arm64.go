//go:build arm64

//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package simd

import (
	"github.com/kshard/vector/internal/noasm"
	"golang.org/x/sys/cpu"
)

// Note: https://github.com/kshard/vector/issues/6
//
// Assembly function causes NaN exception on some input
// SIMD disable until fixed (not a high priority, see noasm)
var ENABLED_EUCLIDEAN = cpu.ARM64.HasASIMD && false

// Squared Euclidean distance between two vectors.
//
//go:noescape
func EuclideanF32(a, b, c []float32)

// Type Class for Euclidean distance
type Euclidean int

func (Euclidean) Distance(a, b []float32) float32 {
	var d [4]float32
	EuclideanF32(a, b, d[:])
	return d[3] + d[2] + d[1] + d[0]
}

func (Euclidean) Equal(a, b []float32) bool {
	return noasm.EqualF32(a, b)
}
