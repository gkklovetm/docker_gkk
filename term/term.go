package term

import (
	"fmt"
	"syscall"
	"unsafe"
)

func p() {
	fmt.Println("11")
}

const (
	ICANON = 0x100
	ICRNL  = 0x100
	IXON   = 0x200
	IXOFF  = 0x400
	ISIG   = 0x80
	ISTRIP = 0x20
	ONLCR  = 0x2
	ECHO   = 0x00000008
)

type Termios struct {
	Iflag  uintptr
	Oflag  uintptr
	Cflag  uintptr
	Lflag  uintptr
	Cc     [20]byte
	Ispeed uintptr
	Ospeed uintptr
}

type State struct {
	termios Termios
}

func IsTerminal(fd int) bool {
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(getTermios), uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func Restore(fd int, state *State) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), uintptr(setTermios), uintptr(unsafe.Pointer(&state.termios)), 0, 0, 0)
	return err
}
