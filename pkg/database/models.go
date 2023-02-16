package database

// -- Tasks
type Task struct {
	UUID    string `json:"uuid" example:"uuidv4" gorm:"primaryKey"`
	Creator string `json:"creator" example:"uuidv4" gorm:"not null"` // Foreign key

	Title   string `json:"title" example:"My first task" gorm:"not null"`
	Content string `json:"content" example:"This is my first task" gorm:"not null"`

	CreatedAt int64 `json:"created_at" example:"1676546709" gorm:"not null"`
	UpdatedAt int64 `json:"updated_at" example:"1676546709" gorm:"not null"`
}

// -- Users
// TODO add 'plan' - free, premium, etc.
// TODO add 'role' - admin, user, etc.
type User struct {
	UUID string `json:"uuid" example:"uuidv4" gorm:"primaryKey"`

	Email    string `json:"email" example:"example@togo.dev" gorm:"uniqueIndex"`
	Password string `json:"password" example:"my_super_secret_password" gorm:"not null"`

	CreatedAt int64 `json:"created_at" example:"1676546709" gorm:"not null"`
	UpdatedAt int64 `json:"updated_at" example:"1676546709" gorm:"not null"`

	Tasks []Task `json:"tasks" gorm:"foreignKey:UserUUID"` // One to many relationship
}

// -- Misc
type UserCredentials struct {
	Email    string `json:"email" example:"example@togo.dev"`
	Password string `json:"password" example:"my_super_secret_password"`
}
