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
	"golang.org/x/sys/cpu"
)

var ENABLED_EUCLIDEAN = cpu.ARM64.HasASIMD

// Squared Euclidean distance between two vectors.
//
//go:noescape
func EuclideanF32(a, b, c []float32)

// Type Class for Euclidean distance
type Euclidean [4]float32

func (d Euclidean) Distance(a, b []float32) float32 {
	EuclideanF32(a, b, d[:])
	return d[3] + d[2] + d[1] + d[0]
}
