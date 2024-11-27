package tg

import "errors"

var (
	ErrSessionsPathNotSet = errors.New("sessions path not set")
	ErrNoSessions         = errors.New("no sessions")
)
