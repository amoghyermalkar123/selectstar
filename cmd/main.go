package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "server",
				Subcommands: []*cli.Command{
					{
						Name:  "start",
						Usage: "start the server",
						Action: func(ctx *cli.Context) error {
							cmd := exec.Command("/bin/bash", "-c", "selectstar-daemon")
							_, err := cmd.Output()
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "stop",
						Usage: "stop the server",
						Action: func(ctx *cli.Context) error {
							cmd := exec.Command("/bin/bash", "-c", "kill `cat /opt/selectstar/selectstar.pid`")
							_, err := cmd.Output()
							if err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
