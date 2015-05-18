// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"fmt"
//	"github.com/odTimeTracker/odtimetracker-go-lib/database"
	"log"
	"os"
	"strings"
)

// Here is implementation of the "report" command.
var cmdReport = &Command{
	Name:      "report",
	Desc:      "Render HTML report from the collected data.",
	UsageDesc: "[--today|--week|--month] [--file=FILE] [--project=PROJECT_NAME] [--tag=TAG]",
	Run:       runReport,
	Help:      helpReport,
}

// Template for help of "report" command.
const cmdReportHelp = `
Usage:

	%[1]s %[2]s %[3]s

%[4]s

Report will be written into the specified file. If file doesn't exist will be created if exist will be overwritten.

If FILE is not set the default "report.html" is used.

Report can be created from data for the specified period, currently are supported these:

	--today                 Report for today activities.
	--week                  Report for activities tracked this week.
	--month                 Report for activities tracked this month (default).

Data for report can be also filtered:

	--project=PROJECT_NAME  Filter activities by the project.
	--tag=TAG               Filter activites by the tag.

Report is written to the specified file (or is used default file name "report.html"):

	--flag=FILE             File name for the generated report.

Examples:

	%[1]s %[2]s
	%[1]s %[2]s --project="odTimeTracker"
	%[1]s %[2]s --month --file="test.html"
	%[1]s %[2]s --month --project="odTimeTracker" --tag="tag1" --file="test.html"

`

var (
	cmdReport_FlagToday   bool = false // Report for today activities.
	cmdReport_FlagWeek    bool = false // Report for activities tracked this week.
	cmdReport_FlagMonth   bool = true  // Report for activities tracked this month.
	
	cmdReport_FlagProject string = "" // Project name filter
	cmdReport_FlagTag     string = "" // Tag filter
	
	cmdReport_FlagFile    string = "report.html"
)

// Execute "report" command. Called from function "main()".
func runReport(cmd *Command, db *sql.DB, args []string) {
	if len(args) > 4 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	log.Println(args)
	for _, a := range args {
		switch a {
		case "--today":
			cmdReport_FlagToday = true
		case "--week":
			cmdReport_FlagWeek = true
		case "--month":
			cmdReport_FlagMonth = true
		}

		if strings.HasPrefix(a, "--project=") == true {
			cmdReport_FlagProject = strings.Replace(a, "--project=", "", 1)
		}

		if strings.HasPrefix(a, "--tag=") == true {
			cmdReport_FlagTag = strings.Replace(a, "--tag=", "", 1)
		}

		if strings.HasPrefix(a, "--file=") == true {
			cmdReport_FlagFile = strings.Replace(a, "--file=", "", 1)
		}
	}

	log.Println("TODO Implement 'report' command!")
}

// Render help for "info" command.
func helpReport(cmd *Command) {
	fmt.Printf(cmdReportHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}

