// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"fmt"
)

var cmdStop = &Command{
	Name:	"stop",
	Desc:	"Stop currently running activity.",
	Run:	runStop,
	Help:	helpStop,
}

func runStop(cmd *Command, args []string) {
	//if len(args) != 0 {
	//	cmd.Usage()
	//}

	fmt.Println("TODO Implement `stop` command!")
	fmt.Println()
}

func helpStop(cmd *Command) {
	fmt.Printf("\nTODO Finish help on `%s`!\n\n", cmd.Name)
}