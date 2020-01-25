package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  	app := &cli.App{
    	Name: "GoJob",
    	Usage: "Jobs",
    	Action: func(c *cli.Context) error {
			fmt.Printf("Hello %q", c.Args().Get(0))
			return nil
		},
		Commands: []*cli.Command{
			{
			  	Name:    "setup",
			  	Usage:   "Sets up tej on a remote machine",
			  	Action:  func(c *cli.Context) error {
					return nil
			  	},
			},
			{
			  	Name:    "submit",
			  	Usage:   "Submits a job to a remote machine",
			  	Action:  func(c *cli.Context) error {
					return nil
			  	},
			},
			{
				Name:    "status",
				Usage:   "Gets the status of a job",
				Action:  func(c *cli.Context) error {
				  return nil
				},
			},
			{
				Name:    "download",
				Usage:   "Downloads files from finished job",
				Action:  func(c *cli.Context) error {
				  return nil
				},
			},  
			{
				Name:    "kill",
				Usage:   "Kills a running job",
				Action:  func(c *cli.Context) error {
				  return nil
				},
			},
			{
				Name:    "delete",
				Usage:   "Deletes a finished job",
				Action:  func(c *cli.Context) error {
				  return nil
				},
			}, 
			{
				Name:    "list",
				Usage:   "Lists remote jobs",
				Action:  func(c *cli.Context) error {
				  return nil
				},
		  	}, 
		},
  	}

  	err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
  	}
}