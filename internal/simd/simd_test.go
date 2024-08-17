//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package simd_test

import (
	"testing"

	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/simd"
	"github.com/kshard/vector/internal/vtest"
)

func TestEuclidean(t *testing.T) {
	if !simd.ENABLED_EUCLIDEAN {
		return
	}

	sut := simd.Euclidean(0)

	vtest.TestEqual(t, vtest.ID, sut)
	vtest.TestDistance(t, vtest.ID, pure.Euclidean(0), sut)
}
