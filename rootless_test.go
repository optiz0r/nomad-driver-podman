// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"path/filepath"
	"testing"

	"github.com/shoenig/test/must"
)

func TestRootlessTaskDirRewritePath(t *testing.T) {
	allocDir := filepath.Join("data", "nomad", "alloc-id")
	mountDir := filepath.Join("run", "user", "12080", "nomad", "alloc-id")
	rootlessDir := &rootlessTaskDir{
		mountDir: mountDir,
		allocDir: allocDir,
	}

	logPath := filepath.Join(allocDir, "alloc", "logs", ".whoami.stdout.fifo")
	must.Eq(t,
		filepath.Join(mountDir, "alloc", "logs", ".whoami.stdout.fifo"),
		rootlessDir.rewritePath(logPath),
	)

	outsidePath := filepath.Join("var", "log", "containers", "whoami.log")
	must.Eq(t, outsidePath, rootlessDir.rewritePath(outsidePath))
}
