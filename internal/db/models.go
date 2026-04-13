package db

import (
	"database/sql"
	"time"
)

type Answer string

const (
	AnswerYes   Answer = "YES"
	AnswerNo    Answer = "NO"
	AnswerMaybe Answer = "MAYBE"
)

type Guest struct {
	ID            int
	Name, Surname string
	Answer        sql.Null[Answer]
	AnsweredAt    sql.Null[time.Time]
	CreatedAt     time.Time
}
