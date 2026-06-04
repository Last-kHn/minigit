package main

import (
	"fmt"
	"os"
)

func ShowStatus() {

	files, err := os.ReadDir(".minigit/staging")

	if err != nil {
		fmt.Println("No existe un repositorio inicializado.")
		return
	}

	fmt.Println("=== ESTADO DEL REPOSITORIO ===")
	fmt.Println()

	if len(files) == 0 {
		fmt.Println("No hay archivos en staging.")
		return
	}

	fmt.Printf("Archivos preparados: %d\n\n", len(files))

	for _, file := range files {
		fmt.Println("- " + file.Name())
	}
}
