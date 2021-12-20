package cmd

import (
	"github.com/honwen/wait-for/poller"
	"github.com/honwen/wait-for/poller/postgres"
	"github.com/urfave/cli"
)

var PostgresCommand = cli.Command{
	Name: "postgres",
	Action: func(c *cli.Context) error {
		postgres, err := postgres.New(
			c.String("connection-string"),
		)
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: postgres.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "connection-string, cs",
			Value:  "",
			Usage:  "psql connection string",
			EnvVar: "WAIT_FOR_POSTGRES_CONNECTION_STRING",
		},
	},
}
