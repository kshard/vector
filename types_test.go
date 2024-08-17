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
	"github.com/kshard/vector/internal/pure"
	"github.com/kshard/vector/internal/vtest"
)

type Node struct {
	ID     int
	Vector vector.F32
}

func ntov(n Node) []float32 { return n.Vector }

func TestContraMap(t *testing.T) {
	node := func(f []float32) Node { return Node{Vector: f} }

	ref := vector.ContraMap[vector.F32, Node]{
		Surface:   pure.Euclidean(0),
		ContraMap: func(n Node) []float32 { return n.Vector },
	}

	sut := vector.ContraMap[vector.F32, Node]{
		Surface:   vector.Euclidean(),
		ContraMap: func(n Node) []float32 { return n.Vector },
	}

	vtest.TestEqual(t, node, sut)
	vtest.TestDistance(t, node, ref, sut)
}
