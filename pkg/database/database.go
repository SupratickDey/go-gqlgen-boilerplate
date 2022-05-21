package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/joho/godotenv"
)

type Conn struct {
	Db     *pgxpool.Pool
	Schema string
}

var DbConn Conn

func DBInit() *Conn {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	schema := os.Getenv("DB_SCHEMA")
	dbURI := "host=" + host + " port=" + port + " user=" + username + " password=" + password + " dbname=" + dbName + " sslmode=" + sslMode
	//nolint:gosec // configuration
	maxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	if err != nil {
		log.Panic("cannot parse DB_MAX_CONN")
	}
	//nolint:gocritic // configuration
	maxIdleTime, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME"))
	if err != nil {
		log.Panic("cannot parse DB_MAX_IDLE_TIME")
	}
	//nolint:gocritic // configuration
	maxConnectionLifetime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME"))
	if err != nil {
		log.Panic("cannot parse DB_MAX_LIFETIME")
	}

	//nolint:gocritic // configuration
	healthCheckPeriod, err := strconv.Atoi(os.Getenv("DB_HEALTHCHECK_PERIOD"))
	if err != nil {
		log.Panic("cannot parse DB_HEALTHCHECK_PERIOD")
	}

	config, err := pgxpool.ParseConfig(dbURI)
	if err != nil {
		log.Println(err)
		log.Panic("cannot parse database URI")
	}
	config.MaxConns = int32(maxOpenConnection)
	config.MaxConnLifetime = time.Duration(maxConnectionLifetime) * time.Minute
	config.HealthCheckPeriod = time.Duration(healthCheckPeriod) * time.Minute
	config.MaxConnIdleTime = time.Duration(maxIdleTime) * time.Minute

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println(err)
		log.Panic("cannot connect to database")
	} else {
		log.Println("database connected")
	}

	DbConn.Db = pool
	DbConn.Schema = schema
	return &DbConn
}

func GetPool() *Conn {
	if DbConn.Db == nil {
		DBInit()
	}

	return &DbConn
}
