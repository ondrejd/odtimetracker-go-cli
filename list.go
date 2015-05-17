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
	"strconv"
	"strings"
	"time"
)

var cmdList = &Command{
	Name:      "list",
	Desc:      "List activities or projects.",
	UsageDesc: "WHAT [--all|--limit=VAL] [--full|--short]",
	Run:       runList,
	Help:      helpList,
}

const cmdListHelp = `
Usage:

	%[1]s %[2]s %[3]s

%[4]s

There are several flags that have influence on output:

Either you can use one of these to set count of results:

	--all         List all possible results.
	--limit=VAL   Set limit for results (default 5).

Or these to set format of the results:

	--full        Full format of the output list.
	--short       Short format of the output list (default).

Examples:

	%[1]s list activities
	%[1]s list activities --all --full
	%[1]s list projects --full
	%[1]s list activities --limit=10

`

var (
	cmdList_flagShowAll   bool = false
	cmdList_flagShowFull  bool = false // If TRUE flag '--short' is used
	cmdList_flagShowLimit int  = 5     // Default value for number of rows in list
)

func runList(cmd *Command, db *sql.DB, args []string) {
	if len(args) < 1 || len(args) > 3 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}
	//log.Println(args)
	for _, a := range args {
		switch a {
		case "--all":
			cmdList_flagShowAll = true
		case "--full":
			cmdList_flagShowFull = true
		}

		if strings.HasPrefix(a, "--limit=") == true {
			lstr := strings.Replace(a, "--limit=", "", 1)
			limit, err := strconv.Atoi(lstr)
			if err != nil {
				limit = 5
			}
			cmdList_flagShowLimit = limit
		}
	}

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
	fmt.Printf(cmdListHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

const (
	ActivityFormatFull  = "%d\t%s-%s\t%s\n"
	ActivityFormatShort = "%d\t%s-%s\t%s\n"
	ProjectFormatFull   = "%d\t%s\t%s\n"
	ProjectFormatShort  = "%d\t%s\n"
)

func listActivities(db *sql.DB) {
	limit := cmdList_flagShowLimit
	if cmdList_flagShowAll == true {
		limit = 0
	}

	activities, err := SqliteStorage.SelectActivities(db, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, a := range activities {
		started, err := a.StartedTime()
		if err != nil {
			started = time.Now()
		}

		stopped, err := a.StoppedTime()
		if err != nil {
			stopped = time.Now()
		}

		if cmdList_flagShowFull == true {
			if a.Description == "" {
				fmt.Printf(ActivityFormatFull, a.ActivityId,
					formatDatetime(started), formatDatetime(stopped),
					a.Name)
			} else {
				fmt.Printf(ActivityFormatFull+"\t%s\n", a.ActivityId,
					formatDatetime(started), formatDatetime(stopped),
					a.Name, a.Description)
			}
		} else {
			fmt.Printf(ActivityFormatShort, a.ActivityId, formatDatetime(started),
				formatDatetime(stopped), a.Name)
		}
	}

	fmt.Println()
}

func listProjects(db *sql.DB) {
	limit := cmdList_flagShowLimit
	if cmdList_flagShowAll == true {
		limit = 0
	}

	projects, err := SqliteStorage.SelectProjects(db, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, p := range projects {
		if cmdList_flagShowFull == true {
			ctime, err := p.CreatedTime()
			if err != nil {
				ctime = time.Now()
			}

			if p.Description == "" {
				fmt.Printf(ProjectFormatFull, p.ProjectId,
					formatDatetime(ctime), p.Name)
			} else {
				fmt.Printf(ProjectFormatFull+"\t%s\n", p.ProjectId,
					formatDatetime(ctime), p.Name, p.Description)
			}
		} else {
			fmt.Printf(ProjectFormatShort, p.ProjectId, p.Name)
		}
	}

	fmt.Println()
}

func formatDatetime(t time.Time) string {
	return fmt.Sprintf("%d.%d %d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
}
