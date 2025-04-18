package main

import (
	"fmt"
	"os"

	"github.com/sa-giga/saga-cli/cmd/saga/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %s\n", err)
		os.Exit(1)
	}
}
