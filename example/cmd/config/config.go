package main

import (
	"fmt"
	"github.com/abtinokhovat/gox/example/config"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg)
}
