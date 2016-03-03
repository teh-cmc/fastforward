package main

import (
	"github.com/codegangsta/cli"
	"github.com/teh-cmc/fastforward"
)

// -----------------------------------------------------------------------------

var app *cli.App

func init() {
	app = cli.NewApp()

	app.Name = "FastForward"
	app.HelpName = app.Name
	app.Version = "0.0.1"
	app.Usage = "Decentralized Kanban tool built upon git, designed for small teams that need to move fast."

	app.Commands = []cli.Command{
		{
			Name:            "init",
			Usage:           "initializes the FastForward branch",
			ArgsUsage:       "<path> (defaults to current directory)",
			SkipFlagParsing: true,
			Action:          func(c *cli.Context) { forward.Init(c) },
		},
		{
			Name:            "task",
			Usage:           "task related sub-commands",
			SkipFlagParsing: true,
			Subcommands: []cli.Command{
				{
					Name:            "new",
					Usage:           "creates a new task",
					SkipFlagParsing: true,
					Action:          func(c *cli.Context) { forward.TaskNew(c) },
				},
				{
					Name:   "list",
					Usage:  "lists tasks",
					Action: func(c *cli.Context) { forward.TaskList(c) },
				},
			},
		},
	}

	app.Authors = []cli.Author{
		{Name: "Clement 'cmc' Rey", Email: "cr.rey.clement@gmail.com"},
	}

	app.Copyright = `The MIT License (MIT) - see LICENSE for more details
   Copyright (c) 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>`
}
