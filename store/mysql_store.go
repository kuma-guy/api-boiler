package store

import (
	"api-boiler/store/schema"
	"database/sql"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

type mySQLStore struct {
	db *sql.DB
}

// NewMySQLStore creates a new store that uses MySQL as a backend.
func NewMySQLStore(username, password, host, port, dbname string) (Store, error) {
	creds := username
	if password != "" {
		creds += ":" + password
	}

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", creds, host, port, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return &mySQLStore{db}, nil
}

func (s *mySQLStore) Ping() error {
	var err error

	for i := 0; i < 10; i++ {
		err = s.db.Ping()
		if err == nil {
			return nil
		}

		log.Warn("Failed to ping the database. Retry in 1s.")
		time.Sleep(time.Second)
	}

	return err
}

func (s *mySQLStore) Migrate() error {
	migrations := &migrate.AssetMigrationSource{
		Asset:    schema.Asset,
		AssetDir: schema.AssetDir,
		Dir:      "store/schema",
	}

	_, err := migrate.Exec(s.db, "mysql", migrations, migrate.Up)
	return err
}
