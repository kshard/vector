//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package vector_test

import (
	"testing"

	"github.com/kshard/vector"
	"github.com/kshard/vector/internal/noasm"
	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/simd"
	"github.com/kshard/vector/internal/vtest"
)

var (
	d  float32
	eq bool
	a  = vtest.Vector()
	b  = vtest.Vector()
	n1 = Node{ID: 1, Vector: a}
	n2 = Node{ID: 2, Vector: b}
)

func BenchmarkPureEqualF32(t *testing.B) {
	euc := pure.Euclidean(0)

	for i := t.N; i > 0; i-- {
		eq = euc.Equal(a, a)
	}
}

func BenchmarkNoAsmEqualF32(t *testing.B) {
	euc := noasm.Euclidean(0)

	for i := t.N; i > 0; i-- {
		eq = euc.Equal(a, a)
	}
}

func BenchmarkPureEuclideanF32(t *testing.B) {
	euc := pure.Euclidean(0)

	for i := t.N; i > 0; i-- {
		d = euc.Distance(a, b)
	}
}

func BenchmarkNoAsmEuclideanF32(t *testing.B) {
	euc := noasm.Euclidean(0)

	for i := t.N; i > 0; i-- {
		d = euc.Distance(a, b)
	}
}

func BenchmarkSIMDEuclideanF32(t *testing.B) {
	euc := simd.Euclidean(0)

	for i := t.N; i > 0; i-- {
		d = euc.Distance(a, b)
	}
}

func BenchmarkContraMapEuclidean(t *testing.B) {
	euc := vector.ContraMap[vector.F32, Node]{
		Surface:   vector.Euclidean(),
		ContraMap: ntov,
	}

	for i := t.N; i > 0; i-- {
		d = euc.Distance(n1, n2)
	}
}

func BenchmarkPureCosineF32(t *testing.B) {
	cos := pure.Cosine(0)

	for i := t.N; i > 0; i-- {
		d = cos.Distance(a, b)
	}
}

func BenchmarkNoAsmCosineF32(t *testing.B) {
	cos := noasm.Cosine(0)

	for i := t.N; i > 0; i-- {
		d = cos.Distance(a, b)
	}
}
