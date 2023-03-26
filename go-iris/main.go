package main

import (
	"fmt"
	"os"

	"go-iris/cmd"
)

func main() {
	app := cmd.New()
	if err := app.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
