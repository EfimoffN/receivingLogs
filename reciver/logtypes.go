package reciver

type AcceptLog struct {
	UserUUID  string
	Timestamp int64
	Events    []EventLog
}

type EventLog struct {
	URL          string
	DataRequest  string
	DataResponse string
}

type SendLog struct {
	LogUUI       string
	IP           string
	UserUUID     string
	Timestamp    int64
	URL          string
	DataRequest  string
	DataResponse string
}
