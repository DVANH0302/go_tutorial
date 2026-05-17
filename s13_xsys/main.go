package main

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	unix.Syscall(unix.SYS_WRITE, 1,
		uintptr(unsafe.Pointer(&[]byte("Hello, World!")[0])),
		uintptr(len("Hello, World!")),
	)
	syscall.Sysinfo(&syscall.Sysinfo_t{})
}
