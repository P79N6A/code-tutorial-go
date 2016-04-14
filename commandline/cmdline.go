package main
import (
	"github.com/codegangsta/cli"
	"os"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) {
		println("Hello friend!")
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:"debug",
			Usage:"xxx",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "add",
			Aliases:     []string{"a"},
			Flags:[]cli.Flag {
				cli.StringFlag{
					Name: "lang",
					Value: "english",
					Usage: "language for the greeting",
				},
			},
			Usage:     "add a task to the list",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args().First())
				fmt.Println(c.FlagNames())
				fmt.Println(c.NumFlags())
				fmt.Println(c.String("lang"))
				fmt.Println(c.GlobalFlagNames())
				fmt.Println(c.GlobalBool("debug"))
			},
		},
		{
			Name:      "complete",
			Aliases:     []string{"c"},
			Usage:     "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
		},
		{
			Name:      "template",
			Aliases:     []string{"r"},
			Usage:     "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) {
						println("new task template: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) {
						println("removed task template: ", c.Args().First())
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
