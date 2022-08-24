package config

import "errors"

var (
	ErrUnknownRecivDatabase = errors.New("unknown database")
	ErrNoAllParametersPSG   = errors.New("there are no all parameters for Postgres")
	ErrNoAllParametersMNG   = errors.New("there are no all parameters for MongoDB")
	ErrNoAllParametersCLC   = errors.New("there are no all parameters for ClickHous")
	ErrNoAllParametersKFK   = errors.New("there are no all parameters for Kafka")
)
