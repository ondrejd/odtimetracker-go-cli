// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.
package main 

import (
	"fmt"
	"os"
)

const (
	AppName = "odTimeTracker CLI Tool" // Application's name
	AppShortName = "odtimetracker" // Application's short name (system name)
	AppVersion = "0.1" // Application's version
	AppInfo = AppName+" "+AppVersion // Application's info line
	AppDesc = "Simple tool for tracking time you have spent working on your projects." // Application's description
)

// Simple struct representing command
type Command struct {
	// Name of a command
	Name string

	// Description of a command
	Desc string

	// Runs the command self.
	Run func(cmd *Command, args []string)
	
	// Prints help on the command.
	Help func(cmd *Command)
}

// Prints usage information for the command.
func (cmd *Command) Usage(Prefix string) {
	fmt.Printf("%s%s\t%s\n", Prefix, cmd.Name, cmd.Desc)
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
    }
    
    if os.Args[1] == "help" {
    	help(os.Args[2:])
    }
    
	os.Exit(0)
}

// Prints usage informations.
func usage() {
   	fmt.Println()
   	fmt.Println(AppDesc)
   	fmt.Println()
   	fmt.Println("Usage:")
   	fmt.Println()
   	fmt.Printf("\t"+AppShortName+" command [arguments]\n")
   	fmt.Println()
   	fmt.Println("Available commands:")
   	fmt.Println()
   	for _, cmd := range commands {
   		cmd.Usage("\t")
   	}
   	fmt.Println()
   	fmt.Printf("Use \""+AppShortName+" help [command]\" for more information about a command.\n")
   	fmt.Println()
   	fmt.Println("Additional help topics:")
   	fmt.Println()
   	fmt.Printf("\tactivityString\tHelp on creating strings describing new activity.\n")
   	fmt.Println()
   	fmt.Printf("Use \""+AppShortName+" help [topic]\" for more information about a topic.\n")
   	fmt.Println()
   	os.Exit(1)
}

// Implements the 'help' command.
// TODO Add help for topic `activityString`!
func help(args []string) {
	if len(args) == 0 {
		usage()
		os.Exit(0)
	}
	
	if len(args) != 1 {
		fmt.Printf("\nUsage:\n\n\t"+AppShortName+" help [command]\n\nToo many arguments given!\n")
		os.Exit(1)
	}
	
	for _, cmd := range commands {
		if args[0] == cmd.Name {
			cmd.Help(cmd)
		}
	}
}
