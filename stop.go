// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

// Here is implementation of the `stop` command.
import (
	"database/sql"
	"fmt"
	"os"
)

var cmdStop = &Command{
	Name:      "stop",
	Desc:      "Stop currently running activity.",
	UsageDesc: "",
	Run:       runStop,
	Help:      helpStop,
}

func runStop(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 0 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	fmt.Println("TODO Implement `stop` command!")
	fmt.Println()
}

func helpStop(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
