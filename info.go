// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

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
		cmd.Usage("\nUsage:\n\n\t")
		os.Exit(1)
	}

	a, err := SqliteStorage.SelectActivityRunning(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TODO Implement `info` command!")
	fmt.Println(a)
}

func helpInfo(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}
