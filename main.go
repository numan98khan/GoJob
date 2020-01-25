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
      fmt.Println("Go Job Submission System")
      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}