
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func addFile(filename string) {

	// Verificar archivo
	source, err := os.Open(filename)

	if err != nil {
		fmt.Println("Archivo no encontrado")
		return
	}

	defer source.Close()

	// Ruta destino
	destPath := filepath.Join(
		".minigit/staging",
		filepath.Base(filename),
	)

	// Crear archivo staging
	dest, err := os.Create(destPath)

	if err != nil {
		fmt.Println("Error agregando archivo")
		return
	}

	defer dest.Close()

	// Copiar contenido
	_, err = io.Copy(dest, source)

	if err != nil {
		fmt.Println("Error copiando archivo")
		return
	}

	fmt.Println("Archivo agregado correctamente")
}

