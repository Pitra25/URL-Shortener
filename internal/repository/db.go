package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type ConfigDB struct {
	TypeDB   string
	DBName   string
	UserName string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

func NewDB(cfg *ConfigDB) *ConfigDB {
	return &ConfigDB{
		TypeDB:   cfg.TypeDB,
		DBName:   cfg.DBName,
		UserName: cfg.UserName,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		SSLMode:  cfg.SSLMode,
	}
}

func (db *ConfigDB) Connect() (*pgx.Conn, error) {
	if db.validator() {
		return nil, fmt.Errorf("database structure has no values")
	}

	connectionStr := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		db.TypeDB,
		db.UserName, db.Password,
		db.Host, db.Port, db.DBName,
		db.SSLMode,
	)

	conn, err := pgx.Connect(
		context.Background(),
		connectionStr,
	)
	if err != nil {
		logrus.Error("Error connection DB: ", err)
		return nil, err
	}

	return conn, nil
}

func (db *ConfigDB) Ping(conn *pgx.Conn) error {
	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func Close(conn *pgx.Conn) error {
	if err := conn.Close(context.Background()); err != nil {
		return err
	}

	return nil
}

func (db *ConfigDB) validator() bool {
	if db.TypeDB == "" || db.DBName == "" ||
		db.UserName == "" || db.Password == "" ||
		db.Host == "" || db.Port == "" || db.SSLMode == "" {
		return true
	}
	return false
}
