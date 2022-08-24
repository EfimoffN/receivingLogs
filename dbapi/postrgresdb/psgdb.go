package dbapi

import (
	"context"

	"github.com/EfimoffN/receivingLogs/reciver"
	"github.com/jmoiron/sqlx"
)

type PSGAPI struct {
	db *sqlx.DB
}

func NewPSGAPI(db *sqlx.DB) *PSGAPI {
	return &PSGAPI{
		db: db,
	}
}

func ConnectDB(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func (api *PSGAPI) SaveLog(ctx context.Context, sLog reciver.SendLog) error {
	const query = `INSERT INTO prj_log(loguui, ip, useruuid, timestamp, event) VALUES (:refid, :linkid, :userid) ON CONFLICT DO NOTHING;`

	if _, err := api.db.NamedExecContext(ctx, query, sLog); err != nil {
		return err
	}

	return nil
}
