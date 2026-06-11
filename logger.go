package main

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {

	file, err := os.OpenFile(
		"minigit.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(
		file,
		"[MINIGIT] ",
		log.Ldate|log.Ltime,
	)
}
