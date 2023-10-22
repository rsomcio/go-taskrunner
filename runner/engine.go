package runner

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

func (e *Engine) RunCommand(command string) error {
	e.startCmd(command)
	// cmd, stdout, stderr, err := e.startCmd(command)
	// if err != nil {
	// 	return err
	// }

	// defer func() {
	// 	stdout.Close()
	// 	stderr.Close()
	// }()
	// _, _ = io.Copy(os.Stdout, stdout)
	// _, _ = io.Copy(os.Stderr, stderr)
	// // wait for command to finish
	// err = cmd.Wait()
	// if err != nil {
	// 	return err
	// }
	return nil
}
