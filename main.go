package main

import (
	"fmt"
	"movies-api/config"
)

func main() {
	fmt.Printf("Configs loaded: API: %v, DB: %v\n", config.GetAPI(), config.GetDB())
}
