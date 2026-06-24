package main

import (
	"os"
	"testing"
)

func TestInitRepository(t *testing.T) {

	InitLogger()

	initRepository()

	if _, err := os.Stat(".minigit"); os.IsNotExist(err) {
		t.Error("No se creó .minigit")
	}
}

func TestStagingDirectory(t *testing.T) {

	InitLogger()

	initRepository()

	if _, err := os.Stat(".minigit/staging"); os.IsNotExist(err) {
		t.Error("No se creó el directorio staging")
	}
}
