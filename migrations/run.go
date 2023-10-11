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

	tx.MustExec("INSERT INTO albums (title, upc, cover, artist_id) VALUES ($1, $2, $3, $4)", "Tyranny", "654184", "https://www.sentireascoltare.com/wp-content/uploads/2014/06/julian-casablancas-tyranny-artwork-650x650.jpeg", 2)
	tx.MustExec("INSERT INTO albums (title, upc, cover, artist_id) VALUES ($1, $2, $3, $4)", "Virtue", "135842", "https://ksassets.timeincuk.net/wp/uploads/sites/55/2018/03/Virtue-1024x1024-920x920.jpg", 2)
	tx.MustExec("INSERT INTO albums (title, upc, cover, artist_id) VALUES ($1, $2, $3, $4)", "Korn", "463123", "https://dvdhaven.s3.amazonaws.com/3078/8718469536375.jpg", 1)
	tx.MustExec("INSERT INTO albums (title, upc, cover, artist_id) VALUES ($1, $2, $3, $4)", "Life is peachy", "985135", "https://i.pinimg.com/originals/5c/5f/3e/5c5f3e1989415d7d5c10cf5bd0938344.jpg", 1)
	tx.MustExec("INSERT INTO albums (title, upc, cover, artist_id) VALUES ($1, $2, $3, $4)", "Follow the leader", "233561", "http://www.metalmusicarchives.com/images/covers/korn-follow-the-leader-20141121221611.jpg", 1)

	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "Take Me In Your Army", 254, 12, 1)
	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "Human sadness", 1054, 5, 1)
	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "QYURRYUS", 314, 4, 2)
	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "Blind", 223, 1, 3)
	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "Twist", 131, 1, 4)
	tx.MustExec("INSERT INTO tracks (title, duration, rank, album_id) VALUES ($1, $2, $3, $4)", "Freak on a leash", 236, 1, 5)
	return tx.Commit()
}
