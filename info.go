// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"fmt"
	"github.com/odTimeTracker/odtimetracker-go-lib/database"
	"log"
	"os"
)

// Here is implementation of the `info` command.
var cmdInfo = &Command{
	Name:      "info",
	Desc:      "Print info about current status.",
	UsageDesc: "",
	Run:       runInfo,
	Help:      helpInfo,
}

// Template for help of "list" command.
const cmdInfoHelp = `
Usage:

	%s %s

%s

If there is a running activity print its name and duration otherwise print 
message that there is no running activity.

`

// Execute "info" command. Called from function "main()".
func runInfo(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 0 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	a, err := database.SelectActivityRunning(db)
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

// Render help for "info" command.
func helpInfo(cmd *Command) {
	fmt.Printf(cmdInfoHelp, AppShortName, cmd.Name, cmd.Desc)
}

