package main

import (
	"os"
	"testing"
)

func TestCreateCommit(t *testing.T) {

	InitLogger()

	initRepository()

	os.WriteFile("commit_test.txt", []byte("contenido"), 0644)

	addFile("commit_test.txt")

	createCommit("commit de prueba")

	files, _ := os.ReadDir(".minigit/commits")

	if len(files) == 0 {
		t.Error("No se creó ningún commit")
	}
}

func TestStagingCleanAfterCommit(t *testing.T) {

	InitLogger()

	initRepository()

	os.WriteFile("clean_test.txt", []byte("datos"), 0644)

	addFile("clean_test.txt")

	createCommit("limpiar staging")

	files, _ := os.ReadDir(".minigit/staging")

	if len(files) != 0 {
		t.Error("El staging no fue limpiado")
	}
}
