package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createBranch(name string) {

	branchPath := filepath.Join(
		".minigit/branches",
		name+".txt",
	)

	// Verificar si ya existe
	if _, err := os.Stat(branchPath); err == nil {
		fmt.Println("La rama ya existe.")
		return
	}

	file, err := os.Create(branchPath)

	if err != nil {
		fmt.Println("Error creando rama.")
		return
	}

	defer file.Close()

	file.WriteString("Rama: " + name)

	fmt.Println("Rama creada:", name)

	Logger.Println("Rama creada: " + name)
}

func listBranches() {

	files, err := os.ReadDir(".minigit/branches")

	if err != nil {
		fmt.Println("No existen ramas.")
		return
	}

	fmt.Println("=== RAMAS ===")

	for _, file := range files {
		fmt.Println("- " + file.Name())
	}
}
