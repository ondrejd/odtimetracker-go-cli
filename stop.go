// Copyright 2015 Ondřej Doněk. All rights reserved.
// See LICENSE file for more details about licensing.

package main

import (
	"database/sql"
	"fmt"
	"github.com/odTimeTracker/odtimetracker-go-lib/database"
	"log"
	"os"
	"time"
)

// Here is implementation of the "stop" command.
var cmdStop = &Command{
	Name:      "stop",
	Desc:      "Stop currently running activity.",
	UsageDesc: "",
	Run:       runStop,
	Help:      helpStop,
}

// Template for help of "stop" command.
const cmdStopHelp = `
Usage:

	%[1]s %[2]s %[3]s

%[4]s

`

// Execute "stop" command. Called from function "main()".
func runStop(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 0 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	ra, err := database.SelectActivityRunning(db)
	if err != nil {
		fmt.Printf("\nThere is no running activity!\n\n")
		os.Exit(1)
	}

	ra.Stopped = time.Now().Format(time.RFC3339)
	_, err = database.UpdateActivity(db, ra)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nActivity '%s' was successfully stopped.\n\n", ra.Name)
}

// Render help for "stop" command.
func helpStop(cmd *Command) {
	fmt.Printf(cmdStopHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

