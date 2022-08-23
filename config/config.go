package config

import "os"

func SetConfiguration() {
	err := os.Setenv("SECRET_KEY", "")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("DB_HOST", "")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("DB_PORT", "")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("DB_USER", "")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("DB_PASSKEY", "")
	if err != nil {
		panic(err)
	}

	err = os.Setenv("DB_NAME", "")
	if err != nil {
		panic(err)
	}
}
