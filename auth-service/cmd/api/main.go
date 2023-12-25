package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/terajari/go-microservices/auth-service/data"
)

const webPort = "80"

var counts int

type Application struct {
	Db     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	db := connectDb()
	if db == nil {
		log.Panic("Failed to connect to database")
	}

	app := &Application{
		Db:     db,
		Models: data.NewModels(db),
	}

	if err := app.Routes().Run(":" + webPort); err != nil {
		log.Fatalf("Failed to start auth service: %s", err)
	}

}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func connectDb() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDb(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready: %s\n", err)
			counts++
		} else {
			log.Printf("Connected to Postgres with count: %d", counts)
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
