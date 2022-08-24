package reciver

import "context"

type SavedLogI interface {
	SaveLog(ctx context.Context, sLog SendLog) error
}
