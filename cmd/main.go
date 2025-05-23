package main

import (
	"fmt"

	"github.com/joaobaronii/to-do-list-go/configs"
)

func main() {
	config := configs.LoadConfig("cmd")	
	fmt.Printf("Config completa: %+v\n", *config)
}