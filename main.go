package main

import (
	"fmt"
	"os"

	"qshell/cmd/root"
)

func main() {
	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
