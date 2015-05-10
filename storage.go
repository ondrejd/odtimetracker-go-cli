// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
//	"database/sql"
//	"fmt"
//	"os"
//	_ "github.com/mattn/go-sqlite3"
)

// Common storage definition.
type Storage struct {
	// Initialize storage.
	Init func() error
	
	// Insert new activity.
	InsertActivity func(Name string, Project string, Tags string, Description string) (Activity, error)
	
	// Insert new project.
	InsertProject func(Name string, Description string) (Project, error)
	
	// Remove activity(-ies) with given Id(s) form the database.
	RemoveActivity func(Id ...int64) (int, error)
	
	// Remove project(s) with given Id(s) form the database.
	RemoveProject func(Id ...int64) (int, error)
	
	// Return activities.
	SelectActivities func() ([]Activity, error)
	
	// Return activity(-ies) by given ID(s). 
	SelectActivityById func(Id ...int64) ([]Activity, error)
	
	// Return currently running activity.
	SelectActivityRunning func() (Activity, error)
	
	// Return projects.
	SelectProjects func() ([]Project, error)
	
	// Return project(s) by given ID(s).
	SelectProjectById func(Id ...int64) ([]Project, error)
	
	// Return single project by given name(s).
	SelectProjectByName func(Name ...string) ([]Project, error)
}