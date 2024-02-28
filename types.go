//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package vector

import "github.com/fogfish/golem/pure"

// Vector of float32
type F32 = []float32

// Generic trait to  estimate "distance" between two vectors
type Surface[Vector any] interface {
	Distance(Vector, Vector) float32
}

// From is a combinator that lifts V ⟼ V ⟼ float32 function to
// an instance of Distance type trait
type From[Vector any] func(Vector, Vector) float32

func (f From[Vector]) Distance(a, b Vector) float32 { return f(a, b) }

// ContraMap is a combinator that build a new instance of type trait Surface[V] using
// existing instance of Distance[A] and f: b ⟼ a
type ContraMap[A, B any] struct {
	Surface[A]
	pure.ContraMap[A, B]
}

// Distance implementation of contra variant functor
func (f ContraMap[A, B]) Distance(a, b B) float32 {
	return f.Surface.Distance(
		f.ContraMap(a),
		f.ContraMap(b),
	)
}
