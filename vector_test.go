//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package vector_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kshard/vector"
	"github.com/kshard/vector/internal/noasm"
	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/simd"
)

//
//
//

type Node struct {
	ID     int
	Vector vector.F32
}

func ntov(n Node) []float32 { return n.Vector }

var (
	n  = 300
	d  float32
	a  vector.F32
	b  vector.F32
	n1 Node
	n2 Node
)

func init() {
	a = randF32()
	b = randF32()
	n1 = Node{ID: 1, Vector: a}
	n2 = Node{ID: 2, Vector: b}
}

func zeroF32() vector.F32 {
	v := make(vector.F32, n)
	for i := 0; i < n; i++ {
		v[i] = 0.0
	}
	return v
}

func randF32() vector.F32 {
	v := make(vector.F32, n)
	for i := 0; i < n; i++ {
		v[i] = rand.Float32()
	}
	return v
}

func equal(a, b float32) bool {
	d := a - b
	return -1e-4 < d && d < 1e-4
}

//
//
//

func TestNoAsmEuclideanF32(t *testing.T) {
	euc := pure.Euclidean(0)
	sut := noasm.Euclidean(0)

	for i := 0; i < n*100; i++ {
		a := randF32()
		b := randF32()

		g := euc.Distance(a, b)
		d := sut.Distance(a, b)
		if !equal(d, g) {
			t.Errorf("failed distance")
		}
	}
}

func TestSIMDEuclideanF32(t *testing.T) {
	if !simd.ENABLED_EUCLIDEAN {
		return
	}

	euc := pure.Euclidean(0)
	sut := simd.Euclidean{}

	for i := 0; i < n*100; i++ {
		a := randF32()
		b := randF32()

		g := euc.Distance(a, b)
		d := sut.Distance(a, b)
		if !equal(d, g) {
			t.Errorf("failed distance")
		}
	}
}

func TestContraMapEuclidean(t *testing.T) {
	euc := pure.Euclidean(0)
	sut := vector.ContraMap[vector.F32, Node]{
		Surface:   vector.Euclidean(),
		ContraMap: func(n Node) []float32 { return n.Vector },
	}

	for i := 0; i < n*100; i++ {
		a := randF32()
		b := randF32()

		g := euc.Distance(a, b)
		d := sut.Distance(Node{Vector: a}, Node{Vector: b})
		if !equal(d, g) {
			t.Errorf("failed distance")
		}
	}
}

func TestNoAsmCosineF32(t *testing.T) {
	euc := pure.Cosine(0)
	sut := noasm.Cosine(0)

	for i := 0; i < n*100; i++ {
		a := randF32()
		b := randF32()

		g := euc.Distance(a, b)
		d := sut.Distance(a, b)
		if !equal(d, g) {
			t.Errorf("failed distance")
		}
	}
}

func TestPureCosineZeroF32(t *testing.T) {
	sut := pure.Cosine(0)

	for i := 0; i < n*100; i++ {
		a := zeroF32()
		b := randF32()

		d := sut.Distance(a, b)
		if math.IsNaN(float64(d)) {
			t.Errorf("failed distance")
		}
	}
}

func TestNoAsmCosineZeroF32(t *testing.T) {
	sut := noasm.Cosine(0)

	for i := 0; i < n*100; i++ {
		a := zeroF32()
		b := randF32()

		d := sut.Distance(a, b)
		if math.IsNaN(float64(d)) {
			t.Errorf("failed distance")
		}
	}
}

//
// Benchmark
//

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

func BenchmarkNoAsmEuclideanUn(t *testing.B) {
	euc := noasm.EuclideanU(0)

	for i := t.N; i > 0; i-- {
		d = euc.Distance(a, b)
	}
}

func BenchmarkSIMDEuclideanF32(t *testing.B) {
	euc := simd.Euclidean{}

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
