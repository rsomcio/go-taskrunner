package runner

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/pkg/errors"
)

const (
	OutputDir = "output"
)

type config struct {
	Port    int
	DestDir string `toml:"dest_dir"`
}

type task struct {
	Name  string `toml:"name"`
	Cmd   string `toml:"cmd"`
	Flags string `toml:"flags"`
}

type tomlConfig struct {
	Version int64  `toml:"version"`
	Config  config `toml:"config"`
	Tasks   []task `toml:"tasks"`
}

func Now() string {
	return time.Now().UTC().Format(YYYYMMDD)
}

func DestDir() string {
	return fmt.Sprintf("%s/%s", OutputDir, Now())
}

func readConfig(path string) (*tomlConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	cfg := new(tomlConfig)
	if err = toml.Unmarshal(data, cfg); err != nil {
		// return nil, errors.New(fmt.Sprintf("failed to unmarshal user setting locale value (%s)", err))
		return nil, errors.WithStack(err)
	}
	return cfg, nil
}

func InitConfig() (cfg *tomlConfig, err error) {
	config, err := readConfig("config.toml")
	if err != nil {
		return nil, err
	}
	return config, nil
}
