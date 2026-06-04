package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func generateHash(text string) string {

	hash := sha1.New()

	hash.Write([]byte(text))

	return fmt.Sprintf("%x", hash.Sum(nil))[:8]
}

func createCommit(message string) {

	// Leer staging
	files, err := os.ReadDir(".minigit/staging")

	if err != nil || len(files) == 0 {
		fmt.Println("No hay archivos en staging")
		return
	}

	// Datos commit
	date := time.Now().Format("2006-01-02 15:04:05")

	content := message + date

	commitID := generateHash(content)

	commitPath := filepath.Join(
		".minigit/commits",
		commitID+".txt",
	)

	file, err := os.Create(commitPath)

	if err != nil {
		fmt.Println("Error creando commit")
		return
	}

	defer file.Close()

	// Escribir información
	file.WriteString("Commit: " + commitID + "\n")
	file.WriteString("Fecha: " + date + "\n")
	file.WriteString("Mensaje: " + message + "\n\n")

	file.WriteString("Archivos:\n")

	for _, stagedFile := range files {

	file.WriteString("- " + stagedFile.Name() + "\n")
	}

	// Limpiar staging
	for _, stagedFile := range files {

	filePath := filepath.Join(
		".minigit/staging",
		stagedFile.Name(),
	)

	os.Remove(filePath)
	}	

	fmt.Println("Commit realizado:", commitID)
	fmt.Println("Staging limpiado correctamente.")
	}

func showLog() {

	files, err := os.ReadDir(".minigit/commits")

	if err != nil || len(files) == 0 {
		fmt.Println("No hay commits")
		return
	}

	fmt.Println("=== HISTORIAL ===")

	for _, file := range files {

		fmt.Println(file.Name())
	}
}
