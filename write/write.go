package write

import (
	"github.com/redprompt/golog/config"
)

type Output interface {
	Write(data []byte, level config.Level)
}
