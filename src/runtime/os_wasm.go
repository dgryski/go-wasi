// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

package runtime

func osinit() {
	// https://webassembly.github.io/spec/core/exec/runtime.html#memory-instances
	physPageSize = 64 * 1024
	initBloc()
	ncpu = 1
	getg().m.procid = 2
}
