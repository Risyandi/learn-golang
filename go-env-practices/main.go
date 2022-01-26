// go setup environment variables using OS packages
package main

import (
	"fmt"
	"os"
)

func envVariable(key string) string {
	// set environment variable
	os.Setenv(key, "gopher")

	// get environment variable
	return os.Getenv(key)
}

func main() {
	value := envVariable("name")
	fmt.Printf("os packege: name = %s \n", value)
	fmt.Printf("environment variable = %s \n", os.Getenv("APP_ENV"))
}
