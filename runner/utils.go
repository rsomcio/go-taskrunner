package runner

import (
	"io"
	"log"
	"os/exec"
)

func (e *Engine) startCmd(cmd string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {

	c := exec.Command("/bin/sh", "-c", cmd)
	if err := c.Run(); err != nil {
		log.Fatalf("error called %s %s", cmd, err)
	}
	// c := exec.Command(cmd)
	// f, err := pty.Start(c)
	// return c, f, f, err

	return nil, nil, nil, nil
}
