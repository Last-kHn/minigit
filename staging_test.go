package main

import (
	"os"
	"testing"
)

func TestAddFile(t *testing.T) {

	InitLogger()

	initRepository()

	os.WriteFile("prueba.txt", []byte("Hola MiniGit"), 0644)

	addFile("prueba.txt")

	if _, err := os.Stat(".minigit/staging/prueba.txt"); os.IsNotExist(err) {
		t.Error("El archivo no fue agregado a staging")
	}
}

func TestAddNonExistingFile(t *testing.T) {

	InitLogger()

	addFile("archivo_inexistente.txt")

	// La prueba pasa si el programa no se cae
}
