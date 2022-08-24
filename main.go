package main

import (
	"context"
	"log"
	"os"

	"github.com/EfimoffN/receivingLogs/config"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name: "recivLog",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "typedb",
				Aliases: []string{"db"},
				Usage:   "specify the database",
			},
		},
		Action: func(cCtx *cli.Context) error {
			ctx := context.Background()

			cfg, err := config.CreateConfig(cCtx.String("typedb"))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
