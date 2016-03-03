package main

import (
	"github.com/codegangsta/cli"
	"github.com/teh-cmc/fastforward/commands"
)

// -----------------------------------------------------------------------------

var app *cli.App

func init() {
	app = cli.NewApp()

	app.Name = "FastForward"
	app.HelpName = app.Name
	app.Version = "0.0.1"
	app.Usage = "Decentralized Kanban tool built upon git, designed for small teams that need to move fast."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "branch, b",
			Value:  "fastforward",
			Usage:  "specifies the name of the branch used by FastForward",
			EnvVar: "FF_BRANCH",
		},
		cli.BoolFlag{
			Name:   "offline, o",
			Usage:  "enables offline mode (diables auto-pulling & auto-pushing)",
			EnvVar: "FF_OFFLINE",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "initializes the FastForward branch",
			Action: func(c *cli.Context) {
				commands.Run(commands.NewCommand(commands.CommandTypeInit), c)
			},
		},
		{
			Name:  "pull",
			Usage: "pulls the FastForward branch",
			Action: func(c *cli.Context) {
				commands.Run(commands.NewCommand(commands.CommandTypePull), c)
			},
		},
		{
			Name:  "push",
			Usage: "pushes the FastForward branch",
			Action: func(c *cli.Context) {
				commands.Run(commands.NewCommand(commands.CommandTypePush), c)
			},
		},
		{
			Name:  "task",
			Usage: "task related sub-commands",
			Subcommands: []cli.Command{
				{
					Name:  "new",
					Usage: "creates a new task",
					Action: func(c *cli.Context) {
						commands.Run(commands.NewCommand(commands.CommandTypeTaskNew), c)
					},
				},
				{
					Name:  "list",
					Usage: "lists tasks",
					Action: func(c *cli.Context) {
						commands.Run(commands.NewCommand(commands.CommandTypeTaskList), c)
					},
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
