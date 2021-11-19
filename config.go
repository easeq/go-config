package config

import (
	"errors"
	"os"
	"sync"

	"github.com/Netflix/go-env"
)

var (
	// ErrEnvConfigLoad returned when env config load results in an error
	ErrEnvConfigLoad = errors.New("error loading env config")
)

// EnvConfig interface for managing env config
type EnvConfig interface {
	UnmarshalEnv(env.EnvSet) error
}

var (
	// envOnce guards initialization by loadEnvironSet
	envOnce sync.Once
	// envSet holds the environment variables
	envSet env.EnvSet
)

// NewEnvConfig adds env config to struct
func NewEnvConfig(cfg EnvConfig) EnvConfig {
	es := EnvSet()
	err := cfg.UnmarshalEnv(es)
	if err != nil {
		panic(ErrEnvConfigLoad)
	}

	return cfg
}

// EnvSet creates and returns a set of environment variables
func EnvSet() env.EnvSet {
	envOnce.Do(loadEnvironSet)
	return envSet
}

// loadEnvironSet loads the env.EnvSet into envSet
func loadEnvironSet() {
	es, err := env.EnvironToEnvSet(os.Environ())
	if err != nil {
		panic(err)
	}

	envSet = es
}
