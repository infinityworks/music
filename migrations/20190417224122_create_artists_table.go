package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20190417224122, Down20190417224122)
}

func Up20190417224122(tx *sql.Tx) error {
	return exec(`CREATE TABLE artists
(
    id          INTEGER PRIMARY KEY,
    external_id TEXT NOT NULL,
    name        TEXT NOT NULL
);
CREATE UNIQUE INDEX artists_external_id ON artists (external_id);
`, tx)
}

func Down20190417224122(tx *sql.Tx) error {
	return exec("DROP TABLE artists;", tx)
}

func exec(query string, tx *sql.Tx) error {
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
