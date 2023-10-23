package runner

import (
	"io"
	"os/exec"

	"github.com/creack/pty"
)

func (e *Engine) startCmd(cmd string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	c := exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(c)
	return c, f, f, err
}
