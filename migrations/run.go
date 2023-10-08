package migrations

import (
	"database/sql"
	"embed"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed *
var migrationsFiles embed.FS

func Run(postgresConn string) error {
	db, err := sql.Open("postgres", postgresConn)
	if err != nil {
		return err
	}

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFiles,
		Root:       ".",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return err
	}

	return nil
}

func PopulateDB(db *sqlx.DB) error {
	count := 0

	err := db.Get(&count, "SELECT COUNT(*) FROM artists")
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO artists (name, picture, fans) VALUES ($1, $2, $3)", "Korn", "https://www.rollingstone.com/wp-content/uploads/2018/06/rs-korn-988f948a-92c1-487c-85dd-e11957f337c1.jpg", 1204085)
	tx.MustExec("INSERT INTO artists (name, picture, fans) VALUES ($1, $2, $3)", "The voidz", "https://wallpapercave.com/wp/wp9103941.jpg", 604354)
	return tx.Commit()
}
