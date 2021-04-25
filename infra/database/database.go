package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase() (*Database, error) {

	uri := getURI()

	pool, err := pgxpool.Connect(context.Background(), uri)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &Database{pool: pool}, nil
}

func (d *Database) Close() {
	d.pool.Close()
}

func getURI() string {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	if err != nil {
		log.Println("error on load db port from env:", err.Error())
		dbPort = 5432
	}
	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass)
}
