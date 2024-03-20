package main

import (
	"github.com/redprompt/golog"
	"github.com/redprompt/golog/config"
	"github.com/redprompt/golog/examples/persist"
)

func main() {
	config.DisableDefaultLogger()

	golog.Configuration.SetLevel(config.Debug)
	// golog.Configuration.SetColor(false)
	// golog.Configuration.SetLabel(false)

	// Show Configuration for logger
	golog.Debug("%+v\n", golog.Configuration)

	// Shows all level options
	golog.Fatal("This is a %s log!", "Fatal")
	golog.Silent("This is a %s log!", "Silent")
	golog.Error("This is a %s log!", "Error")
	golog.Info("This is a %s log!", "Info")
	golog.Warning("This is a %s log!", "Warning")
	golog.Verbose("This is a %s log!", "Verbose")
	golog.Debug("This is a %s log!", "Debug")

	// Call function from different file maintaining golog configuration
	persist.Persistency()
}
