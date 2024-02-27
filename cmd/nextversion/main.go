package main

import (
	"fmt"
	"os"
)

var version string = "v0.0.0-dev"

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
