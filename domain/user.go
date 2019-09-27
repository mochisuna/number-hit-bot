package domain

type EventStatus int

const (
	CLEAR EventStatus = iota
	NODATA
	FAIL
	GAMEOVER
	FAIL_TOO_SMALL
	FAIL_TOO_LARGE
)
const MAXIMUM_MISSCOUNT = 10

type UserID string
type AnswerNumber int
type User struct {
	ID        UserID       `json:"id"`
	MissCount int          `json:"miss_count"`
	Answer    AnswerNumber `json:"answer"`
}
