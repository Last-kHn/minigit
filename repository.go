package main

import (
	"fmt"
	"os"
)

func initRepository() {

	dirs := []string{
		".minigit",
		".minigit/staging",
		".minigit/commits",
    ".minigit/branches",
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
	}

	fmt.Println("Repositorio inicializado")
  Logger.Println("Repositorio inicializado")
}
