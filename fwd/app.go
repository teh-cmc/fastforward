package main

import "github.com/codegangsta/cli"

// -----------------------------------------------------------------------------

var app *cli.App

func init() {
	app = cli.NewApp()

	app.Name = "Forward"
	app.HelpName = app.Name
	app.Version = "0.0.1"
	app.Usage = "Kanban-like tool built on git, designed to move fast."

	//app.HideHelp = true
	app.HideVersion = true

	app.Commands = []cli.Command{}

	app.Authors = []cli.Author{
		{Name: "Clement 'cmc' Rey", Email: "cr.rey.clement@gmail.com"},
	}

	app.Copyright = `The MIT License (MIT) - see LICENSE for more details
   Copyright (c) 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>`
}
