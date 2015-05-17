// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

// Here is implementation of the `list` command.
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

var cmdList = &Command{
	Name:      "list",
	Desc:      "List activities or projects.",
	UsageDesc: "[what]",
	Run:       runList,
	Help:      helpList,
}

var cmdList_longDesc = `
Usage:

	%s %s %s

%s

There are several flags that have influence on output:

	--all			List all possible results.
	--limit=[VAL]	Set limit for results (default 5).
	--full			Full form of the output list.
	--short			Short form of the output list (default).

`

func runList(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 1 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}
	log.Println(args)

	what := args[0]
	if what == "activities" {
		listActivities(db)
	} else if what == "projects" {
		listProjects(db)
	} else {
		err := errors.New(fmt.Sprintf("Wrong argument given - '%s' is not recognized keyword for 'list' command!", what))
		fmt.Println(err)
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}
}

func helpList(cmd *Command) {
	fmt.Printf(cmdList_longDesc, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

func listActivities(db *sql.DB) {
	activities, err := SqliteStorage.SelectActivities(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, a := range activities {
		fmt.Printf("%d\t%s\n", a.ActivityId, a.Name)
	}

	fmt.Println()
}

func listProjects(db *sql.DB) {
	projects, err := SqliteStorage.SelectProjects(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, p := range projects {
		fmt.Printf("%d\t%s\n", p.ProjectId, p.Name)
	}

	fmt.Println()
}
