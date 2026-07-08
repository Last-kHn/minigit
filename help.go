package main

import "fmt"

func showHelp() {

	fmt.Println("===================================")
	fmt.Println("         MiniGit v1.0")
	fmt.Println("===================================")
	fmt.Println()
	fmt.Println("Comandos disponibles:")
	fmt.Println()

	fmt.Println("init")
	fmt.Println("   Inicializa un repositorio.")

	fmt.Println()

	fmt.Println("add <archivo>")
	fmt.Println("   Agrega un archivo al staging.")

	fmt.Println()

	fmt.Println("commit <mensaje>")
	fmt.Println("   Crea un commit.")

	fmt.Println()

	fmt.Println("status")
	fmt.Println("   Muestra los archivos en staging.")

	fmt.Println()

	fmt.Println("log")
	fmt.Println("   Muestra el historial de commits.")

	fmt.Println()

	fmt.Println("branch <nombre>")
	fmt.Println("   Crea una nueva rama.")

	fmt.Println()

	fmt.Println("version")
	fmt.Println("   Muestra la versión del programa.")

	fmt.Println()

	fmt.Println("help")
	fmt.Println("   Muestra este menú.")
}
