package main

import (
	"fmt"
	"linha-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando aplicação...")

	application := app.GenerateCommandLineApp()
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
