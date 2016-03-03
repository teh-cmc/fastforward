package main

import (
	"github.com/codegangsta/cli"
	"github.com/teh-cmc/fastforward"
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
			Name:      "init",
			Usage:     "initializes the FastForward branch",
			ArgsUsage: "<path> (defaults to current directory)",
			Action:    func(c *cli.Context) { forward.Init(c) },
		},
		{
			Name:   "pull",
			Usage:  "synchronizes the FastForward branch",
			Action: commands.Run(commands.Pull, c),
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
