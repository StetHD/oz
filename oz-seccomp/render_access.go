package seccomp

import (
	"fmt"
	"syscall"
)

const (
	F_OK        = 0x0
	R_OK        = 0x4
	W_OK        = 0x2
	X_OK        = 0x1
	EFF_ONLY_OK = 0x08
)

func render_access(pid int, args RegisterArgs) (string, error) {

	flags := map[uint]string{
		R_OK:        "R_OK",
		W_OK:        "W_OK",
		X_OK:        "X_OK",
		EFF_ONLY_OK: "EFF_ONLY_OK",
	}

	mode := args[1]
	path, err := readStringArg(pid, uintptr(args[0]))

	if err != nil {
		return "", err
	}

	flagstr := ""
	if mode == F_OK {
		flagstr = "F_OK"
	} else {
		flagstr = renderFlags(flags, uint(mode))
	}
	callrep := fmt.Sprintf("access(\"%s\", %s)", path, flagstr)

	return fmt.Sprintf("==============================================\nseccomp hit on sandbox pid %v (%v) syscall %v (%v): \n\n%s\nI ==============================================\n\n", pid, getProcessCmdLine(pid), "access", syscall.SYS_ACCESS, callrep), nil
}