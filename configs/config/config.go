package config

import (
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigDB struct{
	config *pgxpool.Config
}

func (cdb *ConfigDB) EnvConfig() (config *pgxpool.Config, err error){
	cdb.config, err = pgxpool.ParseConfig("")
	cdb.config.ConnConfig.Host = os.Getenv("DB_HOST")
	cdb.config.ConnConfig.User = os.Getenv("DB_USER")
	cdb.config.ConnConfig.Password = os.Getenv("DB_PASSWORD")
	cdb.config.ConnConfig.Database = os.Getenv("DB_NAME")
	return cdb.config, err
}