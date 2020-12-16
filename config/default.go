package config

import (
	"log"
	"strconv"
)

var Default = &config{
	Database: Database{
		Driver:          "postgres",
		IP:              "localhost",
		Port:            getPort(),
		User:            "postgres",
		Password:        "25102510",
		Name:            "datahero2",
		ConnMaxIdle:     96,
		ConnMaxOpen:     144,
		ConnMaxLifetime: 10,
		Debug:           false,
		SSLMode:         "disable",
	},
	Server: Server{
		IP:   "127.0.0.1",
		Port: 4510,
	},
}

func getPort() int {
	sPort := "5432"
	port, err := strconv.ParseInt(sPort, 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	return int(port)
}
