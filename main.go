/*
Copyright © 2025 NAME HERE FrakenboK@cR4.sh
*/
package main

import (
	"github.com/FrakenboK/asnx/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
