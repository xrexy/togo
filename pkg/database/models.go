package database

import "time"

// -- Tasks
type Task struct {
	UUID     string `json:"uuid" example:"uuidv4" gorm:"primaryKey"`
	UserUUID string `json:"user_uuid" example:"uuidv4" gorm:"not null"` // Foreign key

	Title   string `json:"title" example:"My first task" gorm:"not null"`
	Content string `json:"content" example:"This is my first task" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" example:"1620000000" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" example:"1620000000" gorm:"not null"`
}

// -- Users
// TODO add 'plan' - free, premium, etc.
// TODO add 'role' - admin, user, etc.
// TODO add 'created_at' and 'updated_at'
type User struct {
	UUID string `json:"uuid" example:"uuidv4" gorm:"primaryKey"`

	Email    string `json:"email" example:"example@togo.dev" gorm:"uniqueIndex"`
	Password string `json:"password" example:"my_super_secret_password" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" example:"1620000000" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" example:"1620000000" gorm:"not null"`

	Tasks []Task `json:"tasks" gorm:"foreignKey:UserUUID"` // One to many relationship
}

// -- Misc
type UserCredentials struct {
	Email    string `json:"email" example:"example@togo.dev"`
	Password string `json:"password" example:"my_super_secret_password"`
}
