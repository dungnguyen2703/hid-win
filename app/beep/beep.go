package beep

import "syscall"

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	beep     = kernel32.NewProc("Beep")
)

type BEEP int

const (
	LOW_BEEP    BEEP = 400
	MEDIUM_BEEP BEEP = 1000
	HIGH_BEEP   BEEP = 1200
)

func Play(freq BEEP, duration uint32) {
	beep.Call(uintptr(freq), uintptr(duration))
}
