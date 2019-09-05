// +build linux darwin !windows

package runner

import (
	"os"
	"syscall"
)

var signals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGKILL,
	syscall.SIGTERM,
	syscall.SIGSTOP,
	syscall.SIGHUP,
	syscall.SIGQUIT,
}
