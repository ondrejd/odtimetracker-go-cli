// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"fmt"
)

var cmdList = &Command{
	Name:	"list",
	Desc:	"List activities or projects.",
	Run:	runList,
	Help:	helpList,
}

func runList(cmd *Command, args []string) {
	//if len(args) != 0 {
	//	cmd.Usage()
	//}

	fmt.Println("TODO Implement `list` command!")
	fmt.Println()
}

func helpList(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}