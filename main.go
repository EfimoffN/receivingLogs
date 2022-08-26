package main

import (
	"errors"
	"log"
	"os"

	"github.com/EfimoffN/receivingLogs/config"
	clcapi "github.com/EfimoffN/receivingLogs/dbapi/clickdb"
	kfkapi "github.com/EfimoffN/receivingLogs/dbapi/kafkadb"
	mngapi "github.com/EfimoffN/receivingLogs/dbapi/mongodb"
	psgapi "github.com/EfimoffN/receivingLogs/dbapi/postrgresdb"
	"github.com/EfimoffN/receivingLogs/reciver"
	"github.com/urfave/cli/v2"
)

var (
	ErrNoDatabaseParameters = errors.New("there are no database connection parameters")
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
			cfg, err := config.CreateConfig(cCtx.String("typedb"))
			if err != nil {
				log.Fatal("service failed on create config", err)
			}

			db, err := createConnectDB(*cfg)
			if err != nil {
				log.Fatal("service failed on create connect to db", err)
			}
			defer db.Close()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createConnectDB(cfg config.ConfigApp) (reciver.LogSaver, error) {
	conectString := ""

	if cfg.PsgConfig.DBname != "" &&
		cfg.PsgConfig.Password != "" &&
		cfg.PsgConfig.SSLmode != "" &&
		cfg.PsgConfig.User != "" {
		db, err := psgapi.ConnectPSG(conectString)
		if err != nil {
			return nil, err
		}

		psg := psgapi.NewPSGAPI(db)
		return psg, nil
	}

	if cfg.MongoConfig.LocalHost != "" {
		db, err := mngapi.ConnectMNG(conectString)
		if err != nil {
			return nil, err
		}

		psg := mngapi.NewMNGAPI(db)
		return psg, nil
	}

	if cfg.ClickConfig.User != "" &&
		cfg.ClickConfig.Password != "" &&
		cfg.ClickConfig.DBname != "" &&
		cfg.ClickConfig.Host != "" &&
		cfg.ClickConfig.SSLmode != "" {
		db, err := clcapi.ConnectCLC(conectString)
		if err != nil {
			return nil, err
		}

		psg := clcapi.NewCLCAPI(db)
		return psg, nil
	}

	if cfg.KafkaConfig.LocalHost != "" &&
		cfg.KafkaConfig.Topic != "" {
		kfk, err := kfkapi.ConnectKafka(cfg.KafkaConfig)
		if err != nil {
			return nil, err
		}

		psg := kfkapi.NewKafkaAPI(kfk, cfg.KafkaConfig.Topic)
		return psg, nil
	}

	return nil, ErrNoDatabaseParameters
}
