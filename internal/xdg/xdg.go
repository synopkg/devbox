// Copyright 2023 Jetpack Technologies Inc and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package xdg

import (
	"os"
	"path/filepath"

	"github.com/synopkg/devbox/internal/envir"
)

func DataSubpath(subpath string) string {
	return filepath.Join(dataDir(), subpath)
}

func ConfigSubpath(subpath string) string {
	return filepath.Join(configDir(), subpath)
}

func CacheSubpath(subpath string) string {
	return filepath.Join(cacheDir(), subpath)
}

func StateSubpath(subpath string) string {
	return filepath.Join(stateDir(), subpath)
}

func dataDir() string   { return resolveDir(envir.XDGDataHome, ".local/share") }
func configDir() string { return resolveDir(envir.XDGConfigHome, ".config") }
func cacheDir() string  { return resolveDir(envir.XDGCacheHome, ".cache") }
func stateDir() string  { return resolveDir(envir.XDGStateHome, ".local/state") }

func resolveDir(envvar, defaultPath string) string {
	dir := os.Getenv(envvar)
	if dir != "" {
		return dir
	}

	home, err := os.UserHomeDir()
	if err != nil {
		home = "~"
	}

	return filepath.Join(home, defaultPath)
}