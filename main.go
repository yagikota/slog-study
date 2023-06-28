package main

import (
	"log"
	"os"
)

func main() {
	// デフォルトのロガーが標準エラーではなく標準出力に書き込まれるように設定
	defaultLogger := log.Default()
	defaultLogger.SetOutput(os.Stdout)
	log.Println("Hello from Go application!")
}
