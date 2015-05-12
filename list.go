// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

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

func runList(cmd *Command, db *sql.DB, args []string) {
	if len(args) != 1 {
		cmd.Usage("\nUsage:\n\n\t")
		os.Exit(1)
	}
	log.Println(args)

	what := args[0]
	if what == "activities" {
		activities, err := SqliteStorage.SelectActivities(db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(activities)
	} else if what == "projects" {
		projects, err := SqliteStorage.SelectProjects(db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(projects)
	} else {
		err := errors.New(fmt.Sprintf("Wrong argument given - '%s' is not recognized keyword for 'list' command!", what))
		fmt.Println(err)
		cmd.Usage("\nUsage:\n\n\t")
		os.Exit(1)
	}
}

func helpList(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
