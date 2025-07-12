//go:build !windows
// +build !windows

package term

import (
	"syscall"
	"unsafe"
)

type winsize struct {
	Row uint16
	Col uint16
	X   uint16
	Y   uint16
}

func getTerminalWidth() (int, error) {
	ws := &winsize{}
	retCode, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return 0, err
	}
	return int(ws.Col), nil
}
