//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

#include "textflag.h"

// func EuclideanF32(a, b, c []float32) float32
TEXT ·EuclideanF32(SB), NOSPLIT, $0
  MOVD a_base + 0(FP), R0
  MOVD b_base + 24(FP), R1
  MOVD c_base + 48(FP), R2

  MOVD a_size + (8)(FP), R20
  MOVD $4, R8
  MUL R20, R8
  ADD R0, R8, R8

	// Drop one pointer since we compare after the loop
  MOVD $16, R7
  SUB R7, R8, R8

	WORD $0x4ea0d400;    // 	fsub.4s	v0, v0, v0

loop:

  WORD $0x4cdf7801;    // 	ld1.4s	{ v1 }, [x0], #16
  WORD $0x4cdf7822;    // 	ld1.4s	{ v2 }, [x1], #16
  WORD $0x4ea2d423;    // 	fsub.4s	v3, v1, v2
  WORD $0x6e23dc63;    // 	fmul.4s	v3, v3, v3
  WORD $0x4e23d400;    // 	fadd.4s	v0, v0, v3

  CMP R0, R8
  BGE loop

  WORD $0x4c9f7840;    // 	st1.4s	{ v0 }, [x2], #16

  RET

