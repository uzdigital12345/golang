package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	
	app.Commands = []*cli.Command{
		{
			Name: "time",
			Action: func(c *cli.Context) error {
				time := time.Now()
				fmt.Println("Hey, Hamidulloh current time is ", time)
				return nil
			},
		},
	}
	app.Run(os.Args)

}
