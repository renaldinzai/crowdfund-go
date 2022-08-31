package config

import (
	"fmt"
	"os"
)

func SetConfiguration() {
	err := os.Setenv("SECRET_KEY", "")
	if err != nil {
		panic(err)
	}

	dbUser := ""
	dbPasskey := ""
	dbHost := ""
	dbPort := ""
	dbName := ""

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPasskey, dbHost, dbPort, dbName)

	err = os.Setenv("DSN", dsn)
	if err != nil {
		panic(err)
	}
}
