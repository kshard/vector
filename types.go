//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package vector

import (
	"github.com/kshard/vector/internal/noasm"
	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/simd"
)

// Vector of float32
type F32 = []float32

const (
	EUCLIDEAN_WITH_PURE = iota
	EUCLIDEAN_WITH_NOASM
	EUCLIDEAN_WITH_SIMD
)

// Squared Euclidean distance between two vectors
func Euclidean() interface{ Distance(F32, F32) float32 } {
	switch euclideanConfig() {
	case EUCLIDEAN_WITH_PURE:
		return pure.Euclidean(0)
	case EUCLIDEAN_WITH_NOASM:
		return noasm.Euclidean(0)
	case EUCLIDEAN_WITH_SIMD:
		return simd.Euclidean{}
	}

	return nil
}

func euclideanConfig() int {
	if simd.ENABLED_EUCLIDEAN {
		return EUCLIDEAN_WITH_SIMD
	}

	if noasm.ENABLED_EUCLIDEAN {
		return EUCLIDEAN_WITH_NOASM
	}

	return EUCLIDEAN_WITH_PURE
}

const (
	COSINE_WITH_PURE = iota
	COSINE_WITH_NOASM
	COSINE_WITH_SIMD
)

// Cosine Distance
func Cosine() interface{ Distance(F32, F32) float32 } {
	switch cosineConfig() {
	case COSINE_WITH_PURE:
		return pure.Cosine(0)
	case COSINE_WITH_NOASM:
		return noasm.Cosine(0)
		// case COSINE_WITH_SIMD:
		// 	return simd.Euclidean{}
	}

	return nil
}

func cosineConfig() int {
	if simd.ENABLED_COSINE {
		return COSINE_WITH_SIMD
	}

	if noasm.ENABLED_COSINE {
		return COSINE_WITH_NOASM
	}

	return COSINE_WITH_PURE
}

const (
	CONFIG_EUCLIDEAN = "euclidean"
	CONFIG_COSINE    = "cosine"
)

// Info about config
func Info() map[string]int {
	info := map[string]int{}

	info[CONFIG_EUCLIDEAN] = euclideanConfig()
	info[CONFIG_COSINE] = cosineConfig()

	return info
}
