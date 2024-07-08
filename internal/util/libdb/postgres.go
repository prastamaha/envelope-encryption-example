package libdb

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Hostname string `env:"DATABASE_HOSTNAME"`
	Port     string `env:"DATABASE_PORT"`
	Username string `env:"DATABASE_USERNAME"`
	Password string `env:"DATABASE_PASSWORD"`
	DBName   string `env:"DATABASE_NAME"`
}

func NewPostgres(hostname, port, username, password, dbName string) *Postgres {
	return &Postgres{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
}

func (p *Postgres) InitDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", "host="+p.Hostname+" port="+p.Port+" user="+p.Username+" password="+p.Password+" dbname="+p.DBName+" sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
