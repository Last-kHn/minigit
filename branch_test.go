package main

import (
	"os"
	"testing"
)

func TestCreateBranch(t *testing.T) {

	InitLogger()

	initRepository()

	createBranch("testing")

	if _, err := os.Stat(".minigit/branches/testing.txt"); os.IsNotExist(err) {
		t.Error("No se creó la rama")
	}
}

func TestDuplicateBranch(t *testing.T) {

	InitLogger()

	initRepository()

	createBranch("develop")
	createBranch("develop")

	// La prueba pasa si no ocurre panic
}
