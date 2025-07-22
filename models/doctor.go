package models

import "time"

type Doctor struct {
	ID            int64
	First_name    string
	Last_name     string
	Avatar        string
	Phone         string
	Rate          int64
	Job           string
	Token         string
	record        string
	Refresh_Token string
	Created_at    time.Time
	Updated_at    time.Time
	User_id       string
}
