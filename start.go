// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"fmt"
	"github.com/ondrejd/odtimetracker/database"
	"log"
	"os"
)

// Here is implementation of the `start` command.
var cmdStart = &Command{
	Name:      "start",
	Desc:      "Start new activity.",
	UsageDesc: "ACTIVITY_STRING",
	Run:       runStart,
	Help:      helpStart,
}

func runStart(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 1 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	ra, err := database.SelectActivityRunning(db)
	if err == nil && ra.ActivityId > 0 {
		fmt.Printf("\nCan not start new activity - another one is still running!\n\n")
		os.Exit(1)
	}

	aStr := args[0]

	var a database.Activity
	err = a.Parse(db, aStr)
	if err != nil {
		log.Fatal(err)
	}

	a, err = database.InsertActivity(db, a.ProjectId, 
		a.Name, a.Description, a.Tags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nActivity successfully started.\n\n")
}

func helpStart(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
