// Copyright 2015 Ondřej Doněk. All rights reserved.
// See LICENSE file for more details about licensing.

package main

import (
	"database/sql"
	"fmt"
	//	db "github.com/odTimeTracker/odtimetracker-go-lib/database"
	"github.com/odTimeTracker/odtimetracker-go-lib/reports"
    "io/ioutil"
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

// Execute "report" command. Called from function "main()".
func runReport(cmd *Command, db *sql.DB, args []string) {
	if len(args) > 4 {
		cmd.Usage("\nUsage:\n\n\t", "\n")
		os.Exit(1)
	}

	// These are default values (will be updated according to given arguments):
	var rtype = reports.ReportTypeMonthly
	var pid int64 = 0 // Project's ID (report just activities of the project)
	var tags string = ""
	var format = reports.ReportFormatHtml

	for _, a := range args {
		switch a {
		case "--today":
			rtype = reports.ReportTypeDaily
		case "--week":
			rtype = reports.ReportTypeWeekly
		case "--month":
			rtype = reports.ReportTypeMonthly
		}

		if strings.HasPrefix(a, "--project=") == true {
			//pname = strings.Replace(a, "--file=", "", 1)
			// ...
			log.Println("TODO Implement --project=[..] flag!")
		}

		if strings.HasPrefix(a, "--tag=") == true {
			//tags = strings.Replace(a, "--tag=", "", 1)
			// ...
			log.Println("TODO Implement --tag=[..] flag!")
		}

		if strings.HasPrefix(a, "--file=") == true {
			//filename = strings.Replace(a, "--file=", "", 1)
			// ...
		}
	}

	log.Println("TODO Implement 'report' command!")
	log.Println(rtype)
	log.Println(pid)
	log.Println(tags)
	log.Println(format)

	r := reports.NewReport(db, rtype, format, pid, tags)
	//log.Println(r.Render())

	// TODO Name of output file should be either set by user or generated using type...
	err := ioutil.WriteFile("report.html", []byte(r.Render()), 0644)
	if err != nil {
		panic(err)
	}
}

// Render help for "info" command.
func helpReport(cmd *Command) {
	fmt.Printf(cmdReportHelp, AppShortName, cmd.Name, cmd.UsageDesc, cmd.Desc)
}
