package write

import (
	"os"
	"sync"

	"github.com/redprompt/golog/config"
)

// Concurrent system output to terminal
type Command struct {
	mutex sync.Mutex
}

func NewCommand() *Command {
	return &Command{mutex: sync.Mutex{}}
}

func (c *Command) Write(data []byte, level config.Level) {
	// Make sure output is not concurrent
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Adds new line
	data = append(data, '\n')

	// Stdout and Stderr accordingly
	switch level {
	case config.Silent, config.Info, config.Verbose, config.Debug:
		os.Stdout.Write(data)
	case config.Fatal, config.Error, config.Warning:
		os.Stderr.Write(data)
	}
}
