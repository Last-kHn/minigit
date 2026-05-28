
package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: minigit <comando>")
		return
	}

	comando := os.Args[1]

	switch comando {

	case "init":
		initRepository()

	case "add":

		if len(os.Args) < 3 {
			fmt.Println("Debe indicar archivo")
			return
		}

		addFile(os.Args[2])

	case "commit":

		if len(os.Args) < 3 {
			fmt.Println("Debe escribir mensaje")
			return
		}

		createCommit(os.Args[2])

	case "log":
		showLog()

	default:
		fmt.Println("Comando no reconocido")
	}
}
