package main

import "github.com/codegangsta/cli"

var SharedFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "apiKey, k",
		Value: "",
		Usage: "API key",
	},
	cli.StringFlag{
		Name:  "name, n",
		Value: "",
		Usage: "heartbeat name",
	},
}

var Commands = []cli.Command{
	{
		Name:        "start",
		Usage:       "Adds a new heartbeat and then sends a hartbeat",
		Description: "Adds a new heartbeat to OpsGenie with the configuration from the given flags. If the heartbeat with the name specified in -name exists, updates the heartbeat accordingly and enables it. It also sends a heartbeat message to activate the heartbeat.",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "description, d",
				Value: "",
				Usage: "Heartbeat description",
			},
			cli.IntFlag{
				Name:  "interval, i",
				Value: 10,
				Usage: "Amount of time OpsGenie waits for a send request before creating alert",
			},
			cli.StringFlag{
				Value: "minutes",
				Name:  "intervalUnit, u",
				Usage: "[minutes, hours or days]",
			},
		},
		Action: func(c *cli.Context) {
			startHeartbeat(extractArgs(c))
		},
	},
	{
		Name:        "stop",
		Usage:       "Disables the heartbeat",
		Description: "Disables the heartbeat specified with -name, or deletes it if -delete is true. This can be used to end the heartbeat monitoring that was previously started.",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "delete",
				Usage: "Delete the heartbeat",
			},
		},
		Action: func(c *cli.Context) {
			stopHeartbeat(extractArgs(c))
		},
	},
	{
		Name:        "send",
		Usage:       "Sends a heartbeat",
		Description: "Sends a heartbeat message to reactivate the heartbeat specified with -name.",
		Action: func(c *cli.Context) {
			sendHeartbeat(extractArgs(c))
		},
	},
}

type OpsArgs struct {
	apiKey       string
	name         string
	description  string
	interval     int
	intervalUnit string
	delete       bool
}

func extractArgs(c *cli.Context) OpsArgs {
	return OpsArgs{c.GlobalString("apiKey"), c.GlobalString("name"), c.String("description"), c.Int("interval"), c.String("intervalUnit"), c.Bool("delete")}
}
