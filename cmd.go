package main

import (
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"goreadmongo/api"
	"goreadmongo/suites"
)

func GetCli() *cli.App {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "list-rooms",
				Aliases: []string{"lr"},
				Usage:   "List all available rooms",
				Action: func(c *cli.Context) error {
					_ = suites.GetAllRooms()
					return nil
				},
			},
			{
				Name:    "start-api",
				Aliases: []string{"api"},
				Usage:   "Act like an API server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Usage:   "Port to start API on",
						Value: 8080,
						DefaultText: "8080",
					},
					&cli.IntFlag{
						Name:    "cors-port",
						Usage:   "Where is your frontend? Needed to passthrough CORS requests",
						Value: 3000,
						DefaultText: "3000",
					},
					&cli.StringFlag{
						Name:    "base",
						Usage:   "Base url of your RC instance",
						Value: "rc.iotfox.ru",
						Required: true,
						DefaultText: "rc.iotfox.ru",
					},
				},

				Action: func(c *cli.Context) error {
					suites.BASEURL = c.String("base")
					api.StartServer(c.Int("port"),c.Int("cors-port"))
					return nil
				},
			},
			{
				Name:    "penetrate",
				Aliases: []string{"dap"},
				Usage:   "reveal message history",
				Subcommands: []*cli.Command{
					{
						Name:  "room",
						Usage: "reveal room's history",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "roomid",
								Usage:   "ID of the room you want to see history from",
								Value: "GENERAL",
								DefaultText: "None",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "base",
								Usage:   "Base url of your RC instance",
								Value: "rc.iotfox.ru",
								DefaultText: "rc.iotfox.ru",
							},
						},
						Action: func(c *cli.Context) error {
							bs := bson.D{{"rid", c.String("roomid")}}
							suites.BASEURL = c.String("base")
							_ = suites.GetAllMessagesByFilter(bs, c.String("base"))
							return nil
						},
					},
					{
						Name:  "whole",
						Usage: "Download all the existing messages (whole dump)",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "base",
								Usage:   "Base url of your RC instance",
								Value: "rc.iotfox.ru",
								DefaultText: "rc.iotfox.ru",
							},
						},
						Action: func(c *cli.Context) error {
							suites.BASEURL = c.String("base")
							bs := bson.D{{}}
							suites.GetAllMessagesByFilter(bs, c.String("base"))
							return nil
						},
					},
				},
			},
		},
	}
	return app
}