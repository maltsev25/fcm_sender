package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Environment struct {
	PORT            string
	NotificationTTL int
	ReleaseMode     bool
	DataPrefix      string
}

var _env *Environment

func init() {
	var count int
	var err error
	var tmp string

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	env := &Environment{}

	tmp = os.Getenv("PORT")
	if tmp == "" {
		tmp = "8000"
	}
	env.PORT = tmp

	tmp = os.Getenv("NOTIFICATION_TTL")
	if tmp == "" {
		tmp = "1"
	}
	count, err = strconv.Atoi(tmp)
	if err != nil {
		log.Fatal("NOTIFICATION_TTL NOT INT")
	}
	env.NotificationTTL = count

	tmp = os.Getenv("MODE")
	if tmp == "release" {
		env.ReleaseMode = true
	}

	tmp = os.Getenv("DATA_PREFIX")
	if tmp == "" {
		tmp = "additional"
	}
	env.DataPrefix = tmp

	_env = env
}

func Get() *Environment {
	return _env
}
