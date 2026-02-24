package app

import (
	database "URL-Shortener/internal/repository"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func dbInit() *pgx.Conn {
	logrus.Debug("DB start initialization")

	config := database.ConfigDB{
		TypeDB:   viper.GetString("database.type"),
		DBName:   viper.GetString("database.database_name"),
		UserName: viper.GetString("database.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		SSLMode:  viper.GetString("database.sslmode"),
	}

	conn, err := config.Connect()
	if err != nil {
		logrus.Fatal("Error connection DB: ", err)
	}

	if err := config.Ping(conn); err != nil {
		logrus.Error("Error ping DB: ", err)
	}

	logrus.Debug("DB initialization")

	return conn
}
