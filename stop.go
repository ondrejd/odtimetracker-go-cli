// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"fmt"
	"github.com/ondrejd/odtimetracker/database"
	"log"
	"os"
	"time"
)

// Here is implementation of the `stop` command.
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

func helpStop(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
