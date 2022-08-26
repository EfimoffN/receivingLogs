package dbapi

import (
	"context"

	"github.com/EfimoffN/receivingLogs/reciver"
	"gopkg.in/mgo.v2"
)

type MNGAAPI struct {
	session *mgo.Session
}

type SendLogMNG struct {
	LogUUI       string `bson:"loguui"`
	IP           string `bson:"ip"`
	UserUUID     string `bson:"useruuid"`
	Timestamp    int64  `bson:"timestamp"`
	URL          string `bson:"url"`
	DataRequest  string `bson:"datarequest"`
	DataResponse string `bson:"dataresponse"`
}

func NewMNGAPI(session *mgo.Session) *MNGAAPI {
	return &MNGAAPI{
		session: session,
	}
}

func ConnectMNG(cfg string) (*mgo.Session, error) {
	session, err := mgo.Dial(cfg)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (api *MNGAAPI) SaveLog(ctx context.Context, sLog reciver.SendLog) error {
	logMNG := SendLogMNG{
		LogUUI:       sLog.LogUUI,
		IP:           sLog.IP,
		UserUUID:     sLog.UserUUID,
		Timestamp:    sLog.Timestamp,
		URL:          sLog.URL,
		DataRequest:  sLog.DataRequest,
		DataResponse: sLog.DataResponse,
	}

	productCollection := api.session.DB("logDB").C("logs")
	err := productCollection.Insert(logMNG)
	if err != nil {
		return err
	}

	return nil
}

func (api *MNGAAPI) Close() {
	api.session.Close()
}
