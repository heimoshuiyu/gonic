package ffprobe

import "errors"

var (
	ErrNoMediaStreams        = errors.New("no media streams")
	ErrNoMediaDuration       = errors.New("no media duration")
	ErrFFProbeScroeNotEnough = errors.New("ffprobe score not enough")
)
