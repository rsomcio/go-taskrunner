package runner

import (
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

type Engine struct {
	running bool
	runArgs []string
}

func NewEngine() (*Engine, error) {
	defaultArgs := []string{"arg"}
	return &Engine{
		running: false,
		runArgs: defaultArgs,
	}, nil
}

// RunCommand executes the command and assumes output directory exists
func (e *Engine) RunCommand(command string) error {

	cmd, stdout, stderr, err := e.startCmd(command)
	if err != nil {
		return err
	}

	defer func() {
		stdout.Close()
		stderr.Close()
	}()

	// To-Do Ray Somcio - move to configuration
	// Save Output
	destDir := fmt.Sprintf("%s/%s", DestDir(), uuid.New())
	os.MkdirAll(destDir, os.ModePerm)
	outFile := fmt.Sprintf("%s/%s", destDir, "filename.txt")
	f, err := os.Create(outFile)
	defer f.Close()
	_, _ = io.Copy(f, stdout)
	_, _ = io.Copy(f, stderr)

	// wait for command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
