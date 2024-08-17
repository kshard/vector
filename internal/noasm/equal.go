//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package noasm

func EqualF32(a, b []float32) bool {
	if len(a) != len(b) {
		panic("vectors must have equal lengths")
	}

	if len(a)%4 != 0 {
		panic("vectors length must be multiple of 4")
	}

	for i := 0; i < len(a); i += 4 {
		av := a[i : i+4 : i+4]
		bv := b[i : i+4 : i+4]

		if av[0] != bv[0] || av[1] != bv[1] || av[2] != bv[2] || av[3] != bv[3] {
			return false
		}
	}

	return true
}
