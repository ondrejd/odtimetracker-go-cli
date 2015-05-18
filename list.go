// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/odTimeTracker/odtimetracker-go-lib/database"
	odstrings "github.com/odTimeTracker/odtimetracker-go-lib/strings"
	"github.com/apcera/termtables"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Here is implementation of the "list" command.
var cmdList = &Command{
	Name:      "list",
	Desc:      "List activities or projects.",
	UsageDesc: "WHAT [--all|--limit=VAL] [--full|--short]",
	Run:       runList,
	Help:      helpList,
}

// Template for help of "list" command.
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

	%[1]s %[2]s activities
	%[1]s %[2]s activities --all --full
	%[1]s %[2]s projects --full
	%[1]s %[2]s activities --limit=10

`

var (
	// If TRUE all records are shown (this suppress limit option).
	cmdList_flagShowAll   bool = false

	// If TRUE flag '--short' is used.
	cmdList_flagShowFull  bool = false

	// Default value for number of rows in list.
	cmdList_flagShowLimit int  = 5
)

// Execute "list" command. Called from function "main()".
func runList(cmd *Command, db *sql.DB, args []string) {
	if len(args) < 1 || len(args) > 3 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

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
		msg := "Wrong argument given - '%s' is not recognized keyword for 'list' command!"
		err := errors.New(fmt.Sprintf(msg, what))
		fmt.Println(err)
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}
}

// Render help for "list" command.
func helpList(cmd *Command) {
	fmt.Printf(cmdListHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

// Render list of activities.
func listActivities(db *sql.DB) {
	limit := cmdList_flagShowLimit
	if cmdList_flagShowAll == true {
		limit = -1
	}

	activities, err := database.SelectActivities(db, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	table := termtables.CreateTable()
	if cmdList_flagShowFull == true {
		table.AddHeaders("AID", "Project", "Name", /*"Description", */"Tags", "Started", "Stopped", "Duration")
	} else {
		table.AddHeaders("AID", "Project", "Name", "Started", "Duration")
	}

	for _, a := range activities {
		// TODO This needs to be rewritten!!!
		var p database.Project
		projects, _ := database.SelectProjectById(db, a.ProjectId)
		if len(projects) == 1 {
			p = projects[0]
		}

		starttime, err := a.StartedTime()
		if err != nil {
			starttime = time.Now()
		}
		started := odstrings.FormatTime(starttime)

		if cmdList_flagShowFull == true {
			var stopped string
			stoptime, err := a.StoppedTime()
			if err != nil {
				stopped = ""
			} else {
				stopped = odstrings.FormatTime(stoptime)
			}

			table.AddRow(a.ActivityId, p.Name, a.Name, /*a.Description, */a.Tags, started, stopped, a.Duration())
		} else {
			table.AddRow(a.ActivityId, p.Name, a.Name, started, a.Duration())
		}
	}

	fmt.Println(table.Render())
	fmt.Println()
}

// Render list of projects.
func listProjects(db *sql.DB) {
	limit := cmdList_flagShowLimit
	if cmdList_flagShowAll == true {
		limit = -1
	}

	projects, err := database.SelectProjects(db, limit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	table := termtables.CreateTable()
	if cmdList_flagShowFull == true {
		table.AddHeaders("PID", "Name", "Description", "Created")
	} else {
		table.AddHeaders("PID", "Name", "Created")
	}

	for _, p := range projects {
		ctime, err := p.CreatedTime()
		if err != nil {
			ctime = time.Now()
		}
		created := odstrings.FormatTime(ctime)

		if cmdList_flagShowFull == true {
			table.AddRow(p.ProjectId, p.Name, p.Description, created)
		} else {
			table.AddRow(p.ProjectId, p.Name, created)
		}
	}

	fmt.Println(table.Render())
	fmt.Println()
}

