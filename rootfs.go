package main

import (
	"os"

	"github.com/k1LoW/errors"
	"golang.org/x/sys/unix"
)

// rootfs設定
type RootfsConfig struct {
	// ルートファイルシステムのパス
	RootfsPath string `json:"rootfs_path"`
}

func SetupRootfs(c RootfsConfig) error {
	if err := unix.Chroot(c.RootfsPath); err != nil {
		return errors.WithStack(err)
	}

	if err := os.Chdir("/"); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
