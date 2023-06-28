package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(
		os.Stderr,
		"MyApplication: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Llongfile,
	)
	logger.Println("Hello from Go application!")
}
