package main

import (
	"fmt"
	"os"
	"time"

	"github.com/navaz-alani/clubhouse/client"
	"github.com/urfave/cli/v2"
)

func main() {
	config, err := client.LoadConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("client configuration could not be loaded")
		os.Exit(1)
	}
	config.Init()
	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:  "new-story",
			Usage: "Create a new story in a project",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: true},
				&cli.StringFlag{Name: "project", Aliases: []string{"proj"}, Required: true},
				&cli.StringFlag{Name: "description", Aliases: []string{"desc"}, Required: true},
				&cli.IntFlag{Name: "deadline", Aliases: []string{"end"}, Required: true},
			},
			Action: func(c *cli.Context) error {
				name := c.String("name")
				project := c.String("project")
				description := c.String("description")
				daysToDeadline := c.Int("deadline")
				dl := time.Now().AddDate(0, 0, daysToDeadline).Format(time.RFC3339)
				err := config.StoryCreate(name, project, dl, description)
				if err != nil {
					fmt.Println(err)
					fmt.Println("failed to create new story")
					os.Exit(1)
				}
				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
