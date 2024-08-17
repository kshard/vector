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

//
// Euclidean
//

const (
	EUCLIDEAN_WITH_PURE = iota
	EUCLIDEAN_WITH_NOASM
	EUCLIDEAN_WITH_SIMD
)

// Squared Euclidean distance between two vectors
func Euclidean() interface {
	Equal(F32, F32) bool
	Distance(F32, F32) float32
} {
	switch euclideanConfig() {
	case EUCLIDEAN_WITH_PURE:
		return pure.Euclidean(0)
	case EUCLIDEAN_WITH_NOASM:
		return noasm.Euclidean(0)
	case EUCLIDEAN_WITH_SIMD:
		return simd.Euclidean(0)
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

//
// Cosine
//

const (
	COSINE_WITH_PURE = iota
	COSINE_WITH_NOASM
	COSINE_WITH_SIMD
)

// Cosine Distance
func Cosine() interface {
	Equal(F32, F32) bool
	Distance(F32, F32) float32
} {
	switch cosineConfig() {
	case COSINE_WITH_PURE:
		return pure.Cosine(0)
	case COSINE_WITH_NOASM:
		return noasm.Cosine(0)
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
