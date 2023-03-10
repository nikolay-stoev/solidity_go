// Copyright (c) 2014-2018 The Notify Authors. All rights reserved.
// Use of this source code is governed by the MIT license that can be
// found in the LICENSE file.

//go:build !darwin && !linux && !freebsd && !dragonfly && !netbsd && !openbsd && !windows && !kqueue && !solaris && !illumos
// +build !darwin,!linux,!freebsd,!dragonfly,!netbsd,!openbsd,!windows,!kqueue,!solaris,!illumos

package notify

import "errors"

// newWatcher stub.
func newWatcher(chan<- EventInfo) watcher {
	return watcherStub{errors.New("notify: not implemented")}
}
