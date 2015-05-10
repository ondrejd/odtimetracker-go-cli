// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"fmt"
)

var cmdInfo = &Command{
	Name:	"info",
	Desc:	"Print info about current status.",
	Run:	runInfo,
	Help:	helpInfo,
}

func runInfo(cmd *Command, args []string) {
	//if len(args) != 0 {
	//	cmd.Usage()
	//}

	fmt.Println("TODO Implement `info` command!")
	fmt.Println()
}

func helpInfo(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}