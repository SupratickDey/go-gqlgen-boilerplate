package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"strconv"
	"time"

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
	maxOpenConnection, err := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	if err != nil {
		log.Println(err)
		maxOpenConnection = 5
	}
	maxIdleTime, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME"))
	if err != nil {
		log.Println(err)
		maxIdleTime = 5
	}
	maxConnectionLifetime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME"))
	if err != nil {
		log.Println(err)
		maxConnectionLifetime = 2
	}

	healthCheckPeriod, err := strconv.Atoi(os.Getenv("DB_HEALTHCHECK_PERIOD"))
	if err != nil {
		log.Println(err)
		healthCheckPeriod = 2
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
