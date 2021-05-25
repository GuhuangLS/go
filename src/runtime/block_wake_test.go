// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Futex is only available on DragonFly BSD, FreeBSD and Linux.
// The race detector emits calls to split stack functions so it breaks
// the test.

// +build dragonfly freebsd linux
// +build !race

package runtime_test

import (
	"runtime"
	"testing"
	"time"
)

var testG []uintptr

func TestBlockWakeG(t *testing.T) {
	for i := 0; i < 1; i++ {
		go func() {
			testG = append(testG, runtime.GetG())
			runtime.BlockG()
			runtime.ClearGStatus()
		}()
	}
	<-time.After(time.Second)

	for _, g := range testG {
		runtime.WakeG(g)
	}
	<-time.After(time.Second)
}
