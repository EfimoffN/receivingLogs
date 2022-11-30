package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

			rc := reciver.NewReciver(cfg.Port, cfg.Host, db, cCtx.Context)
			srv, err := rc.CreateServerLog()
			if err != nil {
				log.Fatal("service failed on start server", err)
			}

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			log.Print("Server Started")

			<-done
			log.Print("Server Stopped")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer func() {
				cancel()
				db.Close()
			}()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatalf("Server Shutdown Failed:%+v", err)
			}
			log.Print("Server Exited Properly")

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
		cfg.PsgConfig.User != "" &&
		cfg.PsgConfig.Port != "" {
		wp := "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
		connectionString := fmt.Sprintf(
			wp,
			"localhost",
			cfg.PsgConfig.Port,
			cfg.PsgConfig.User,
			cfg.PsgConfig.Password,
			cfg.PsgConfig.DBname,
			cfg.PsgConfig.SSLmode,
		)

		db, err := psgapi.ConnectPSG(connectionString)
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
