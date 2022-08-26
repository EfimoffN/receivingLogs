package reciver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type LogSaver interface {
	SaveLog(ctx context.Context, sLog SendLog) error
	Close()
}

type Reciver struct {
	IP  string
	DB  LogSaver
	ctx context.Context
}

func NewReciver(ip string, db LogSaver, ctx context.Context) *Reciver {
	return &Reciver{
		IP:  ip,
		DB:  db,
		ctx: ctx,
	}
}

func (rc *Reciver) StartServerLog() error {
	http.HandleFunc("/log", rc.logReception)
	err := http.ListenAndServe(":"+rc.IP, nil)
	if err != nil {
		return err
	}

	return nil
}

func (rc *Reciver) logReception(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var acceptLog AcceptLog
	err = json.Unmarshal([]byte(body), &acceptLog)
	if err != nil {
		fmt.Println(err)
	}
}

func (rc *Reciver) logProcessing(acceptLog AcceptLog) {
	var sendLog []SendLog = make([]SendLog, 0)

	for _, ac := range acceptLog.Events {
		uuidEvent := strings.Replace(uuid.New().String(), "-", "", -1)

		s := SendLog{
			LogUUI:       uuidEvent,
			IP:           rc.IP,
			UserUUID:     acceptLog.UserUUID,
			Timestamp:    acceptLog.Timestamp,
			URL:          ac.URL,
			DataRequest:  ac.DataRequest,
			DataResponse: ac.DataResponse,
		}

		sendLog = append(sendLog, s)
	}

	rc.logWrite(sendLog)
}

func (rc *Reciver) logWrite(sendLog []SendLog) {
	for _, l := range sendLog {
		rc.DB.SaveLog(rc.ctx, l)
	}
}
