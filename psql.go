package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func psqlDNS() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		envVars.PsqlHost,
		envVars.PsqlPort,
		envVars.PsqlUser,
		envVars.PsqlPassword,
		envVars.PsqlDatabase,
		envVars.PsqlSslMode,
	)
}

type psqlClient struct{ *sql.DB }

func connPsql() *psqlClient {
	db, err := sql.Open("postgres", psqlDNS())
	if err != nil {
		panic(err)
	}
	return &psqlClient{db}
}

func (c psqlClient) ping() error {
	return c.DB.Ping()
}

func (c psqlClient) watch() error {
	for {
		if err := c.ping(); err != nil {
			return err
		}
		fmt.Printf("It is success to ping: %+v\n", time.Now())
		time.Sleep(5 * time.Second)
	}
}
