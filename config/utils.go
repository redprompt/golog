package config

import (
	"bytes"
	"io"
	"log"
	"os"
)

// DisableDefaultLogger disables the default logger.
func DisableDefaultLogger() {
	log.SetFlags(0)
	log.SetOutput(io.Discard)
}

// EnableDefaultLogger enables the default logger.
func EnableDefaultLogger() {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(os.Stderr)
}

func SetBuffer(message string) ([]byte, error) {
	buffer := &bytes.Buffer{}
	buffer.Grow(len(message))
	buffer.WriteString(message)
	return buffer.Bytes(), nil
}
