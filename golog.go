package golog

import (
	"fmt"
	"os"

	"github.com/redprompt/golog/config"
	"github.com/redprompt/golog/write"
)

var Configuration *Logger

func init() {
	Configuration = &Logger{}
	Configuration.SetLevel(config.Info)
	Configuration.SetWrite(write.NewCommand())
	Configuration.SetLabel(true)
	Configuration.SetColor(true)
}

type Logger struct {
	level  config.Level
	output write.Output
	label  bool
	color  bool
}

func (l *Logger) Log(message string, msgLevel config.Level) {
	if msgLevel <= l.level {
		if l.label {
			message = config.Format(message, msgLevel, l.color)
		}
		data, err := config.SetBuffer(message)
		if err != nil {
			return
		}
		l.output.Write(data, l.level)

		if msgLevel == config.Fatal {
			os.Exit(1)
		}
	}
}

func (l *Logger) SetLevel(level config.Level) {
	l.level = level
}

func (l *Logger) SetWrite(command write.Output) {
	l.output = command
}

func (l *Logger) SetLabel(label bool) {
	l.label = label
}

func (l *Logger) SetColor(color bool) {
	l.color = color
}

/*
Functions regarding all log levels below
*/
func Fatal(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Fatal)
}

func Silent(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Silent)
}

func Error(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Error)
}

func Info(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Info)
}

func Warning(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Warning)
}

func Verbose(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Verbose)
}

func Debug(msg string, args ...interface{}) {
	message := fmt.Sprintf(msg, args...)
	Configuration.Log(message, config.Debug)
}

/*
End of log levels print
*/
