package main

import (
	"github.com/kelseyhightower/envconfig"
)

var envVars Variables = initVariables()

func initVariables() Variables {
	var vars Variables
	if err := envconfig.Process("", &vars); err != nil {
		panic(err)
	}
	return vars
}

// Variables .
type Variables struct {
	PsqlPort     int    `envconfig:"PSQL_PORT"`
	PsqlHost     string `envconfig:"PSQL_HOST"`
	PsqlUser     string `envconfig:"PSQL_USER"`
	PsqlPassword string `envconfig:"PSQL_PASSWORD"`
	PsqlDatabase string `envconfig:"PSQL_DATABASE"`
	PsqlSslMode  string `envconfig:"PSQL_SSLMODE"`
}
