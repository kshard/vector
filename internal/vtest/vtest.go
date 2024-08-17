//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

// Package vtest is test suite for vector algebra re-usable across
// provided implementation variants
package vtest

import (
	"math/rand/v2"
	"testing"

	"github.com/chewxy/math32"
)

type F32 = []float32

// SUT Vector dimension
const N = 300

// Number of iterations
const I = N * 100

func ID(x []float32) []float32 { return x }

func TestEqual[T any](t *testing.T, fmap func([]float32) T, sut interface{ Equal(T, T) bool }) {
	t.Helper()

	t.Run("ZeqZ", func(t *testing.T) {
		a := fmap(zero())
		if !sut.Equal(a, a) {
			t.Errorf("zero vectors must be equal")
		}
	})

	t.Run("ZeqV", func(t *testing.T) {
		for i := 0; i < I; i++ {
			a := fmap(zero())
			b := fmap(Vector())
			if sut.Equal(a, b) {
				t.Errorf("zero vector must not be equal to non-zero")
			}
		}
	})

	t.Run("VeqV", func(t *testing.T) {
		for i := 0; i < I; i++ {
			a := fmap(Vector())
			b := fmap(Vector())
			if !sut.Equal(a, a) {
				t.Errorf("non zero vector must be equal to itself")
			}

			if sut.Equal(a, b) {
				t.Errorf("non zero vector must not be equal")
			}
		}
	})
}

func TestDistance[T any](t *testing.T, fmap func([]float32) T, ref, sut interface{ Distance(T, T) float32 }) {
	t.Helper()

	t.Run("ZeqZ", func(t *testing.T) {
		a := fmap(zero())

		got := sut.Distance(a, a)
		expect := ref.Distance(a, a)
		if !equal(expect, got) {
			t.Errorf("unexpected distance: expect = %f, got = %f", expect, got)
		}
	})

	t.Run("ZeqV", func(t *testing.T) {
		for i := 0; i < I; i++ {
			a := fmap(zero())
			b := fmap(Vector())

			got := sut.Distance(a, b)
			expect := ref.Distance(a, b)
			if !equal(expect, got) {
				t.Errorf("unexpected distance: expect = %f, got = %f", expect, got)
			}
		}
	})

	t.Run("VeqV", func(t *testing.T) {
		for i := 0; i < I; i++ {
			a := fmap(Vector())
			b := fmap(Vector())

			got := sut.Distance(a, b)
			expect := ref.Distance(a, b)
			if !equal(expect, got) {
				t.Errorf("unexpected distance: expect = %f, got = %f", expect, got)
			}
		}
	})
}

//------------------------------------------------------------------------------

func zero() F32 {
	v := make(F32, N)
	for i := 0; i < N; i++ {
		v[i] = 0.0
	}
	return v
}

func Vector() F32 {
	v := make(F32, N)
	for i := 0; i < N; i++ {
		v[i] = 2.0*rand.Float32() - 1.0
	}
	return v
}

func equal(a, b float32) bool {
	if math32.IsNaN(a) && math32.IsNaN(b) {
		return true
	}

	if math32.IsNaN(a) || math32.IsNaN(b) {
		return false
	}

	e := float32(1e-5)
	if a > 10.0 {
		e = 1e-4
	}
	if a > 80.0 {
		e = 1e-3
	}

	d := a - b
	return -1*e < d && d < e
}
