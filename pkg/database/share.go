package database

type MessageStruct struct {
	ErrorMessage string `json:"error_message" example:"User already exists"`
	CreatedAt    int64  `json:"unix" example:"1620000000"`
}
