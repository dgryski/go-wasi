// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall_test

import (
	"syscall"
	"testing"
)

func TestJoinPath(t *testing.T) {
	tests := []struct {
		dir, file, path string
	}{
		{".", ".", "."},
		{"./", "./", "./"},
		{"././././", ".", "."},
		{".", "./././", "./"},
		{".", "a", "a"},
		{".", "a/b", "a/b"},
		{".", "..", ".."},
		{".", "../", "../"},
		{".", "../..", "../.."},
		{".", "../..//..///", "../../../"},
		{"/", "/", "/"},
		{"/", "a", "/a"},
		{"/", "a/b", "/a/b"},
		{"/a", "b", "/a/b"},
		{"/", ".", "/"},
		{"/", "..", "/"},
		{"/", "../../", "/"},
		{"/a", "../", "/"},
		{"/a/b/c", "..", "/a/b"},
		{"/a/b/c", "..///..///", "/a/"},
		{"/a/b/c", "..///..///..", "/"},
		{"/a/b/c", "..///..///..///..", "/"},
		{"/a/b/c", "..///..///..///..///..", "/"},
		{"/a/b/c/", "/d/e/f/", "/a/b/c/d/e/f/"},
		{"a/b/c/", ".", "a/b/c"},
		{"a/b/c/", "./d", "a/b/c/d"},
		{"a/b/c/", "./d", "a/b/c/d"},
		{"../", "..", "../.."},
		{"a/b/c/d", "e/../..", "a/b/c"},
		{"a/b/c/d", "./e/../..", "a/b/c"},
		{"a/b/c/d", "./e/..//../../f/g//", "a/b/f/g/"},
		{"../../../", "a/../../b/c", "../../b/c"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			path := syscall.JoinPath(test.dir, test.file)
			if path != test.path {
				t.Errorf("join(%q,%q): want=%q got=%q", test.dir, test.file, test.path, path)
			}
		})
	}
}
