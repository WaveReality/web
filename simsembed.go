// Copyright (c) 2026, The WaveReality Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !generatehtml

package main

import "github.com/WaveReality/web/sims"

func init() {
	embedSim = sims.Embed
}
