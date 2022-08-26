package reciver

import "context"

type LogSaver interface {
	SaveLog(ctx context.Context, sLog SendLog) error
}
