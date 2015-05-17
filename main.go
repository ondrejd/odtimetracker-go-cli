// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.
package main

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	AppName      = "odTimeTracker"                                                          // Application's name
	AppShortName = "odtimetracker"                                                          // Application's short name (system name)
	AppVersion   = "0.1"                                                                    // Application's version
	AppInfo      = AppName + " " + AppVersion                                               // Application's info line
	AppDesc      = "Simple tool for tracking time you have spent working on your projects." // Application's description
)

// Simple struct representing command
type Command struct {
	Name      string // Name of a command
	Desc      string // Description of a command
	UsageDesc string // Usage description (arguments)

	// Runs the command self.
	Run func(cmd *Command, db *sql.DB, args []string)

	// Prints help on the command.
	Help func(cmd *Command)
}

// Prints usage information for the command.
func (cmd *Command) Usage(prefix string, suffix string) {
	fmt.Printf("%s%s %s\t%s\n%s", prefix,
		cmd.Name, cmd.UsageDesc, cmd.Desc, suffix)
}

// All commands supported by this tool
var commands = []*Command{
	cmdInfo,
	cmdList,
	cmdStart,
	cmdStop,
}

// Main (entry) function.
func main() {
	fmt.Println(AppInfo)

	if len(os.Args) <= 1 {
		usage()
		return
	}

	if os.Args[1] == "help" {
		help(os.Args[2:])
		return
	}

	db, err := SqliteStorage.Init()
	if err != nil {
		fmt.Printf("Error occured during initializing database connection:\n\n%s\n\n", err.Error())
		return
	}
	defer db.Close()

	for _, cmd := range commands {
		if os.Args[1] == cmd.Name {
			cmd.Run(cmd, db, os.Args[2:])
			return
		}
	}

	fmt.Printf("Unknown command '%s'.\n\nRun '%s help' for usage.\n", os.Args[1], AppShortName)
}

// Prints usage informations.
func usage() {
	fmt.Printf("\n%s\n\n", AppDesc)
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\t%s command [arguments]\n\n", AppShortName)
	fmt.Printf("Available commands:\n\n")
	for _, cmd := range commands {
		fmt.Printf("\t%s\t%s\n", cmd.Name, cmd.Desc)
	}
	fmt.Printf("\nUse \"%s help [command]\" for more information about a command.\n\n", AppShortName)
	//fmt.Printf("Additional help topics:\n\n")
	//fmt.Printf("\tactivityString\tHelp on creating strings describing new activity.\n\n")
	//fmt.Printf("Use \"%s help [topic]\" for more information about a topic.\n\n", AppShortName)
}

// Implements the 'help' command.
// TODO Add help for topic `activityString`!
func help(args []string) {
	if len(args) == 0 {
		usage()
		os.Exit(0)
	}

	if len(args) > 1 {
		fmt.Printf("\nUsage:\n\n\t%s help [command]\n\nToo many arguments given!\n", AppShortName)
		os.Exit(1)
	}

	for _, cmd := range commands {
		if args[0] == cmd.Name {
			cmd.Help(cmd)
			os.Exit(0)
		}
	}

	fmt.Printf("Unknown command name %s given\nRun '%s help' for usage.\n", args[0], AppShortName)
	os.Exit(1)
}
