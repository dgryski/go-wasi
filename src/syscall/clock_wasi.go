// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasi

package syscall

type clockid uint32

const (
	clock_realtime clockid = iota
	clock_monotonic
	clock_process_cputime_id
	clock_thread_cputime_id
)

type timestamp uint64

// TODO: where did we land on go:noescape?

//go:wasmimport wasi_snapshot_preview1 clock_res_get
func __wasi_clock_res_get(id clockid, precision *timestamp) Errno

//go:wasmimport wasi_snapshot_preview1 clock_time_get
func __wasi_clock_time_get(id clockid, precision timestamp, time *timestamp) Errno
