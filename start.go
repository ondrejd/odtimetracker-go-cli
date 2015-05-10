// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"fmt"
)

var cmdStart = &Command{
	Name:	"start",
	Desc:	"Start new activity.",
	Run:	runStart,
	Help:	helpStart,
}

func runStart(cmd *Command, args []string) {
	//if len(args) != 0 {
	//	cmd.Usage()
	//}

	fmt.Println("TODO Implement `start` command!")
	fmt.Println()
}

func helpStart(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}