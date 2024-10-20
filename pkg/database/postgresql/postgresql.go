package postgresql

import (
	"database/sql"
	"fmt"
	"products/internal/config"

	_ "github.com/lib/pq"
)

func ConnectToDB(conf *config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s",
		conf.User,
		conf.Database,
		conf.Password,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
