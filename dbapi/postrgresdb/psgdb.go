package dbapi

import (
	"context"

	"github.com/EfimoffN/receivingLogs/reciver"
	_ "github.com/jackc/pgx/v5/stdlib"
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

func ConnectPSG(cfg string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", cfg)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func (api *PSGAPI) SaveLog(ctx context.Context, sLog reciver.SendLog) error {
	const query = `INSERT INTO prj_log(loguui, ip, useruuid, tstamp, logurl, datarequest, dataresponse) VALUES (:loguui, :ip, :useruuid, :timestamp, :url, :datarequest, :dataresponse) ON CONFLICT DO NOTHING;`

	if _, err := api.db.NamedExecContext(ctx, query, sLog); err != nil {
		return err
	}

	return nil
}

func (api *PSGAPI) Close() {
	api.db.Close()
}
