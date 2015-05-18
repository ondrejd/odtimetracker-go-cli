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

// Here is implementation of the `start` command.
var cmdStart = &Command{
	Name:      "start",
	Desc:      "Start new activity.",
	UsageDesc: "ACTIVITY_STRING",
	Run:       runStart,
	Help:      helpStart,
}

// Template for help of "start" command.
const cmdStartHelp = `
Usage:

	%[1]s %[2]s %[3]s

%[4]s

ACTIVITY_STRING should be in this format:

	[Activity name]@[Project name];[comma-separated tags]#[Description]

Order is important but all parts are optional except the name of activity of course. For more details see examples below.

If project name is not found in database than new project with given name is created.

Examples:

	%[1]s start "New activity@Project name;tag1,tag2#Some activity description."
	%[1]s start "Another activity@Project name"
	%[1]s start "Yet other activity;tag1,tag2"
	%[1]s start "Other activity@Project name#Some activity description."

`

// Execute "start" command. Called from function "main()".
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

// Render help for "start" command.
func helpStart(cmd *Command) {
	fmt.Printf(cmdStartHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

