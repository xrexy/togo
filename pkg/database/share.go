package database

import "time"

type MessageStruct struct {
	ErrorCode    int       `json:"error_code" example:"400"`
	ErrorMessage string    `json:"error_message" example:"User already exists"`
	CreatedAt    time.Time `json:"created_at" example:"1620000000"`
}
