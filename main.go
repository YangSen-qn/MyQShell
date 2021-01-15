package main

import (
	"fmt"
	"os"

	"./cmd/root"
)

func main() {
	if err := root.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
