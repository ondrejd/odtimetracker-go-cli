// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

// Here is implementation of the `info` command.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var cmdInfo = &Command{
	Name:      "info",
	Desc:      "Print info about current status.",
	UsageDesc: "",
	Run:       runInfo,
	Help:      helpInfo,
}

func runInfo(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 0 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	a, err := SqliteStorage.SelectActivityRunning(db)
	if err == sql.ErrNoRows {
		fmt.Printf("\nThere is no running activity.\n\n")
	} else if err != nil {
		fmt.Printf("\nFatal error occured!\n\n")
		log.Fatal(err)
	} else {
		fmt.Printf("\nThere is running activity '%s'.\nTime spent up to now: %s\n\n",
			a.Name, a.Duration())
	}
}

func helpInfo(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
