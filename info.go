//
// Copyright (C) 2024 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/kshard/vector
//

package vector

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
