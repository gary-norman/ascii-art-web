package ascii_art_web

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16 // unused
	Ypixel uint16 // unused
}

func GetWinSize() Winsize {
	// Get the file descriptor for stdout
	fd := syscall.Stdout

	// Create an instance of Winsize
	var ws Winsize

	// Use the TIOCGWINSZ ioctl system call to get the window size
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if err != 0 {
		fmt.Println("Error getting terminal size:", err)
	}
	return ws
}

func GetArtWidth(s string) int {
	length := len(s)
	return length
}
