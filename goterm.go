package goterm

import (
	"golang.org/x/sys/unix"
)

func IsTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), ioctlReadTermios)
	return err == nil
}
