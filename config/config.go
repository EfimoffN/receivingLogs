package config

import (
	"os"
)

type ConfigApp struct {
	PsgConfig   PsgConfig
	MongoConfig MongoConfig
	ClickConfig ClickConfig
	KafkaConfig KafkaConfig
	Port        string
	Host        string
}

type PsgConfig struct {
	User     string
	Password string
	DBname   string
	SSLmode  string
	Port     string
}

type MongoConfig struct {
	LocalHost string
}

type ClickConfig struct {
	User     string
	Password string
	DBname   string
	Host     string
	Port     string
	SSLmode  string
}

type KafkaConfig struct {
	LocalHost string
	Topic     string
	ClientID  string
}

func CreateConfig(typedb string) (*ConfigApp, error) {
	cfg := ConfigApp{}
	cfg.Port = os.Getenv("portR")
	cfg.Host = os.Getenv("host")

	switch typedb {
	case "psg":
		cfgPSG, err := getConfigPSG()
		if err != nil {
			return nil, err
		}
		cfg.PsgConfig = *cfgPSG
	case "mng":
		cfgMNG, err := getConfigMNG()
		if err != nil {
			return nil, err
		}
		cfg.MongoConfig = *cfgMNG
	case "clc":
		cfgCLC, err := getConfigCLC()
		if err != nil {
			return nil, err
		}
		cfg.ClickConfig = *cfgCLC
	case "kfk":
		cfgCLC, err := getConfigCLC()
		if err != nil {
			return nil, err
		}

		cfgKFK, err := getConfigKFK()
		if err != nil {
			return nil, err
		}

		cfg.ClickConfig = *cfgCLC
		cfg.KafkaConfig = *cfgKFK
	default:
		return nil, ErrUnknownRecivDatabase
	}

	return &cfg, nil
}

func getConfigPSG() (*PsgConfig, error) {
	user := os.Getenv("userPSG")
	if len(user) == 0 {
		return nil, ErrNoAllParametersPSG
	}

	password := os.Getenv("passwordPSG")
	if len(password) == 0 {
		return nil, ErrNoAllParametersPSG
	}

	dbname := os.Getenv("dbnamePSG")
	if len(dbname) == 0 {
		return nil, ErrNoAllParametersPSG
	}

	sslmode := os.Getenv("sslmodePSG")
	if len(sslmode) == 0 {
		return nil, ErrNoAllParametersPSG
	}

	port := os.Getenv("portDB")
	if len(port) == 0 {
		return nil, ErrNoAllParametersPSG
	}

	psgConfig := PsgConfig{
		User:     user,
		Password: password,
		DBname:   dbname,
		SSLmode:  sslmode,
		Port:     port,
	}

	return &psgConfig, nil
}

func getConfigMNG() (*MongoConfig, error) {
	localhost := os.Getenv("localhostMNG")
	if len(localhost) == 0 {
		return nil, ErrNoAllParametersMNG
	}

	mngConfig := MongoConfig{
		LocalHost: localhost,
	}
	return &mngConfig, nil
}

func getConfigCLC() (*ClickConfig, error) {
	user := os.Getenv("userCLC")
	if len(user) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	password := os.Getenv("passwordCLC")
	if len(password) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	host := os.Getenv("hostCLC")
	if len(host) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	dbname := os.Getenv("dbnameCLC")
	if len(dbname) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	sslmode := os.Getenv("sslmodeCLC")
	if len(sslmode) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	port := os.Getenv("portCLC")
	if len(sslmode) == 0 {
		return nil, ErrNoAllParametersCLC
	}

	clcConfig := ClickConfig{
		User:     user,
		Password: password,
		Host:     host,
		DBname:   dbname,
		SSLmode:  sslmode,
		Port:     port,
	}

	return &clcConfig, nil
}

func getConfigKFK() (*KafkaConfig, error) {
	localhost := os.Getenv("localhostKFK")
	if len(localhost) == 0 {
		return nil, ErrNoAllParametersKFK
	}

	topic := os.Getenv("topicKFK")
	if len(localhost) == 0 {
		return nil, ErrNoAllParametersKFK
	}

	clientId := os.Getenv("clientKafkaID")
	if len(localhost) == 0 {
		return nil, ErrNoAllParametersKFK
	}

	kfkConfig := KafkaConfig{
		LocalHost: localhost,
		Topic:     topic,
		ClientID:  clientId,
	}
	return &kfkConfig, nil
}
