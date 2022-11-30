package dbapi

import (
	"context"

	"github.com/EfimoffN/receivingLogs/reciver"
	"github.com/uptrace/go-clickhouse/ch"
)

type CLCAPI struct {
	ch *ch.DB
}

func NewCLCAPI(ch *ch.DB) *CLCAPI {
	return &CLCAPI{
		ch: ch,
	}
}

func ConnectCLC(cfg string, ctx context.Context) (*ch.DB, error) {
	// clickhouse://<user>:<password>@<host>:<port>/<database>?sslmode=disable
	db := ch.Connect(
		ch.WithDSN(cfg),
	)

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func (api *CLCAPI) SaveLog(ctx context.Context, sLog reciver.SendLog) error {
	const query = `INSERT INTO prj_log(loguui, ip, useruuid, timestamp, url, datarequest, dataresponse) VALUES (:loguui, :ip, :useruuid, :timestamp, :url, :datarequest, :dataresponse) ON CONFLICT DO NOTHING;`

	if _, err := api.ch.ExecContext(ctx, query, sLog); err != nil {
		return err
	}

	return nil
}

func (api *CLCAPI) Close() {
	api.ch.Close()
}
