package main

import (
	"fmt"
	"os"

	"qshell/cmd/root_cmd"
)

func main() {
	if err := root_cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
