package utils

import (
	"os"
	"testing"
)

func TestSetEnvFromFile(t *testing.T) {
	SetEnvFromFile()

	env := os.Getenv("ENV")

	if env == "" {
		t.Errorf("env is empty")
	}
}

func TestGetChanges(t *testing.T) {
	out := GetChanges()

	if out == "" {
		t.Error("the output is empty")
	}
}
