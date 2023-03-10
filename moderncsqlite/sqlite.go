package moderncsqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"time"

	"go.bbkane.com/taggedmarks2/taggedmarks"
	_ "modernc.org/sqlite"
)

//go:embed migration/*.sql
var migrationFS embed.FS

type TaggedmarkService struct {
	db *sql.DB
}

func NewTaggedmarkService(dsn string) (*TaggedmarkService, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("db open error: %s: %w", dsn, err)
	}

	// TODO: Re-enable this once I need the files I guess :)
	// // Enable WAL. SQLite performs better with the WAL  because it allows
	// // multiple readers to operate while data is being written.
	// if _, err := db.Exec(`PRAGMA journal_mode = wal;`); err != nil {
	// 	return nil, fmt.Errorf("enable wal: %w", err)
	// }

	// Enable foreign key checks. For historical reasons, SQLite does not check
	// foreign key constraints by default... which is kinda insane. There's some
	// overhead on inserts to verify foreign key integrity but it's definitely
	// worth it.
	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, fmt.Errorf("foreign keys pragma: %w", err)
	}

	// Busy timeouts set how long write transactions will wait to start.
	// If unset, writes will fail immediately if another write is running
	if _, err := db.Exec(`PRAGMA busy_timeout = 5000;`); err != nil {
		return nil, fmt.Errorf("busy timeout: %w", err)
	}

	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	return &TaggedmarkService{db: db}, nil

}

func (ts *TaggedmarkService) CreateTaggedmark(ctx context.Context, tm *taggedmarks.Taggedmark) error {
	err := withTx(ts.db, func(tx *sql.Tx) error {

		// NOTE: if nothing is modified,
		// then no IDs are returned with RETURNING. This means
		// we always want to modify something so we can use this appoach :)

		tmCreateTime := (*NullTime)(&tm.CreateTime)
		tmUpdateTime := (*NullTime)(&tm.UpdateTime)
		var scannedTmCreateTime NullTime
		var scannedTmUpdateTime NullTime
		// Insert basic information
		{
			err := tx.QueryRowContext(
				ctx,
				`
				INSERT INTO taggedmark (
					url,
					create_time,
					update_time
				)
				VALUES (?, ?, ?)
				ON CONFLICT(url)
				DO UPDATE SET update_time = ?
				RETURNING id, create_time, update_time
				`,
				tm.URL,
				tmCreateTime,
				tmUpdateTime,
				tmUpdateTime,
			).Scan(&tm.ID, &scannedTmCreateTime, &scannedTmUpdateTime)
			tm.CreateTime = (time.Time)(scannedTmCreateTime)
			tm.UpdateTime = (time.Time)(scannedTmUpdateTime)
			if err != nil {
				return fmt.Errorf("taggedmark initial insert err: %w", err)
			}
		}

		// Loop through tags and INSERT or UPDATE each one
		for i := 0; i < len(tm.Tags); i++ {
			tagCreateTime := (*NullTime)(&tm.Tags[i].CreateTime)
			tagUpdateTime := (*NullTime)(&tm.Tags[i].UpdateTime)
			var scannedTagCreateTime NullTime
			var scannedTagUpdateTime NullTime
			err := tx.QueryRowContext(
				ctx,
				`
				INSERT INTO tag (
					name,
					create_time,
					update_time
				)
				VALUES(?, ?, ?)
				ON CONFLICT (name)
				DO UPDATE SET update_time = ?
				RETURNING id, create_time, update_time
				`,
				tm.Tags[i].Name,
				tagCreateTime,
				tagUpdateTime,
				tagUpdateTime,
			).Scan(&tm.Tags[i].ID, &scannedTagCreateTime, &scannedTagUpdateTime)
			tm.Tags[i].CreateTime = time.Time(scannedTagCreateTime)
			tm.Tags[i].UpdateTime = time.Time(scannedTagUpdateTime)
			if err != nil {
				return fmt.Errorf("taggedmark tag upsert err: %w", err)
			}

			_, err = tx.ExecContext(
				ctx,
				`
				INSERT INTO taggedmark_tag (
					taggedmark_id,
					tag_id,
					update_time
				)
				VALUES (?, ?, ?)
				ON CONFLICT(tag_id, taggedmark_id)
				DO UPDATE SET update_time = ?
				`,
				tm.ID,
				tm.Tags[i].ID,
				// Use the bookmark create time
				tmCreateTime,
				tmUpdateTime,
			)
			if err != nil {
				return fmt.Errorf("taggedmark taggedmark_tag upsert err: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("CreateTaggedmark err: %w", err)
	}
	return err
}
