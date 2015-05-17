// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
)

// Common storage definition.
type Storage struct {
	// Initialize storage.
	Init func() (db *sql.DB, err error)

	// Insert new activity.
	InsertActivity func(db *sql.DB, pid int64, name string, desc string, tags string) (Activity, error)

	// Insert new project.
	InsertProject func(db *sql.DB, name string, desc string) (Project, error)

	// Remove activity(-ies) with given Id(s) form the database.
	RemoveActivity func(db *sql.DB, Id ...int64) (int, error)

	// Remove project(s) with given Id(s) form the database.
	RemoveProject func(db *sql.DB, Id ...int64) (int, error)

	// Return activities.
	SelectActivities func(db *sql.DB) ([]Activity, error)

	// Return activity(-ies) by given ID(s).
	SelectActivityById func(db *sql.DB, Id ...int64) ([]Activity, error)

	// Return currently running activity.
	SelectActivityRunning func(db *sql.DB) (Activity, error)

	// Return projects.
	SelectProjects func(db *sql.DB) ([]Project, error)

	// Return project(s) by given ID(s).
	SelectProjectById func(db *sql.DB, Id ...int64) ([]Project, error)

	// Return single project by given name(s).
	SelectProjectByName func(db *sql.DB, Name ...string) ([]Project, error)

	// Update activity in the database
	UpdateActivity func(db *sql.DB, a Activity) (int64, error)

	// Update project in the database
	UpdateProject func(db *sql.DB, p Project) (int64, error)
}
