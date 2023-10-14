package multi

import (
	"io/fs"
	"path/filepath"

	"github.com/synopkg/devbox"
	"github.com/synopkg/devbox/internal/debug"
	"github.com/synopkg/devbox/internal/impl/devopt"
)

func Open(opts *devopt.Opts) ([]devbox.Devbox, error) {
	defer debug.FunctionTimer().End()

	var boxes []devbox.Devbox
	err := filepath.WalkDir(
		".",
		func(path string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !dirEntry.IsDir() && filepath.Base(path) == "devbox.json" {
				optsCopy := *opts
				optsCopy.Dir = path
				box, err := devbox.Open(&optsCopy)
				if err != nil {
					return err
				}
				boxes = append(boxes, box)
			}

			return nil
		},
	)

	return boxes, err
}
