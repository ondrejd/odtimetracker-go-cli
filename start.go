// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

// Here is implementation of the `start` command.
import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var cmdStart = &Command{
	Name:      "start",
	Desc:      "Start new activity.",
	UsageDesc: "[activityString]",
	Run:       runStart,
	Help:      helpStart,
}

func runStart(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 1 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	activityString := args[0]
	log.Println("Start activity with string: %s\n", activityString)

	fmt.Println("TODO Implement `start` command!")
	fmt.Println()
}

func helpStart(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
