package main

import (
	"fmt"
	"os"

	"github.com/vladesco/e-commerce/internal/monolith"
	"github.com/vladesco/e-commerce/store"
)

func main() {
	monolith := monolith.Monolith{}

	monolith.AddModule(&store.StoreModule{})

	err := monolith.Bootstrap()

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}
