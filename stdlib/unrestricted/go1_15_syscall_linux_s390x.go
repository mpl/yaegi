// Code generated by 'yaegi extract syscall'. DO NOT EDIT.

// +build go1.15

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Exec":              reflect.ValueOf(syscall.Exec),
		"Exit":              reflect.ValueOf(syscall.Exit),
		"ForkExec":          reflect.ValueOf(syscall.ForkExec),
		"Kill":              reflect.ValueOf(syscall.Kill),
		"PtraceAttach":      reflect.ValueOf(syscall.PtraceAttach),
		"PtraceCont":        reflect.ValueOf(syscall.PtraceCont),
		"PtraceDetach":      reflect.ValueOf(syscall.PtraceDetach),
		"PtraceGetEventMsg": reflect.ValueOf(syscall.PtraceGetEventMsg),
		"PtraceGetRegs":     reflect.ValueOf(syscall.PtraceGetRegs),
		"PtracePeekData":    reflect.ValueOf(syscall.PtracePeekData),
		"PtracePeekText":    reflect.ValueOf(syscall.PtracePeekText),
		"PtracePokeData":    reflect.ValueOf(syscall.PtracePokeData),
		"PtracePokeText":    reflect.ValueOf(syscall.PtracePokeText),
		"PtraceSetOptions":  reflect.ValueOf(syscall.PtraceSetOptions),
		"PtraceSetRegs":     reflect.ValueOf(syscall.PtraceSetRegs),
		"PtraceSingleStep":  reflect.ValueOf(syscall.PtraceSingleStep),
		"PtraceSyscall":     reflect.ValueOf(syscall.PtraceSyscall),
		"RawSyscall":        reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":       reflect.ValueOf(syscall.RawSyscall6),
		"Reboot":            reflect.ValueOf(syscall.Reboot),
		"Shutdown":          reflect.ValueOf(syscall.Shutdown),
		"StartProcess":      reflect.ValueOf(syscall.StartProcess),
		"Syscall":           reflect.ValueOf(syscall.Syscall),
		"Syscall6":          reflect.ValueOf(syscall.Syscall6),

		// type definitions
		"PtraceFpregs": reflect.ValueOf((*syscall.PtraceFpregs)(nil)),
		"PtracePer":    reflect.ValueOf((*syscall.PtracePer)(nil)),
		"PtracePsw":    reflect.ValueOf((*syscall.PtracePsw)(nil)),
		"PtraceRegs":   reflect.ValueOf((*syscall.PtraceRegs)(nil)),
	}
}
