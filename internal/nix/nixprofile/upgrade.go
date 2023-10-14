// Copyright 2023 Jetpack Technologies Inc and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package nixprofile

import (
	"os"

	"github.com/synopkg/devbox/internal/devpkg"
	"github.com/synopkg/devbox/internal/lock"
	"github.com/synopkg/devbox/internal/nix"
)

func ProfileUpgrade(ProfileDir string, pkg *devpkg.Package, lock *lock.File) error {
	idx, err := ProfileListIndex(
		&ProfileListIndexArgs{
			Lockfile:   lock,
			Writer:     os.Stderr,
			Package:    pkg,
			ProfileDir: ProfileDir,
		},
	)
	if err != nil {
		return err
	}

	return nix.ProfileUpgrade(ProfileDir, idx)
}
