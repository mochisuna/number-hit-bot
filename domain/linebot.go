package domain

type EventMode int

const (
	Follow EventStatus = iota
	Message
)

type EventStatus int

const (
	CLEAR EventStatus = iota
	NODATA
	FAIL
	GAMEOVER
	FAIL_TOO_SMALL
	FAIL_TOO_LARGE
)
