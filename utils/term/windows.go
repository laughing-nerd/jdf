//go:build windows
// +build windows

package term

import (
	"os"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")
var getConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")

type short int16
type word uint16

type coord struct {
	X short
	Y short
}

type smallRect struct {
	Left   short
	Top    short
	Right  short
	Bottom short
}

type consoleScreenBufferInfo struct {
	Size              coord
	CursorPosition    coord
	Attributes        word
	Window            smallRect
	MaximumWindowSize coord
}

func getTerminalWidth() (int, error) {
	handle := syscall.Handle(os.Stdout.Fd())
	var info consoleScreenBufferInfo

	ret, _, err := getConsoleScreenBufferInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&info)),
	)

	if ret == 0 {
		return 0, err
	}
	return int(info.Window.Right - info.Window.Left + 1), nil
}
