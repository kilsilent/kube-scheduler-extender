// +build !windows

package debugger

// compareSignal is the signal to trigger node info. For non-windows
// environment it's SIGUSR2.
var CompareSignal = syscall.SIGUSR2
