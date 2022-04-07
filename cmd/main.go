package cmd

import (
	"fmt"
	"github.com/fades-io/reg/internal/logs"
	"log"

	"github.com/fades-io/reg/internal"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf(logs.FileAccessFailed, err)
	} else {
		fmt.Println(logs.FileAccessSuccess)
	}

	internal.Run()
}
