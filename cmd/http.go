package cmd

import (
	"strings"

	"github.com/honwen/wait-for/poller"
	"github.com/honwen/wait-for/poller/http"
	"github.com/urfave/cli"
)

var HTTPCommand = cli.Command{
	Name: "http",
	Action: func(c *cli.Context) error {
		url := c.String("url")
		if len(url) < 1 {
			cli.ShowAppHelp(c)
			cli.OsExiter(1)
		}
		if !(strings.HasPrefix(url, `https://`) || strings.HasPrefix(url, `http://`)) {
			url = `http://` + url
		}
		h, err := http.New(c.String("method"), url, c.String("body"))
		if err != nil {
			return err
		}

		p := poller.Poller{
			Timeout:    c.GlobalDuration("timeout"),
			Interval:   c.GlobalDuration("poll-interval"),
			CheckReady: h.CheckReady,
		}

		return p.WaitReady()
	},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "method, m",
			Value: "GET",
			Usage: "http request method to use for polling",
		},
		cli.StringFlag{
			Name:  "url",
			Value: "",
			Usage: "http uri to poll status of",
		},
		cli.StringFlag{
			Name:  "body",
			Value: "",
			Usage: "optional body to send",
		},
	},
}
