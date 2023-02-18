package database

// -- Tasks
type Task struct {
	UUID   string `json:"uuid" example:"uuidv4" gorm:"primary_key"`
	UserID string `json:"user_id" example:"uuidv4" gorm:""` // Foreign key

	Title   string `json:"title" example:"My first task" gorm:"not null"`
	Content string `json:"content" example:"This is my first task" gorm:"not null"`

	CreatedAt int64 `json:"created_at" example:"1676546709" gorm:"not null"`
	UpdatedAt int64 `json:"updated_at" example:"1676546709" gorm:"not null"`
}

// -- Users
// TODO add 'plan' - free, premium, etc.
type User struct {
	UUID      string `json:"uuid" example:"uuidv4" gorm:"primary_key"` // Primary key
	Email     string `json:"email" example:"example@togo.dev" gorm:"unique;not null"`
	Password  string `json:"-" example:"my_super_secret_password" gorm:"not null"`
	CreatedAt int64  `json:"created_at" example:"1676546709" gorm:"not null"`
	UpdatedAt int64  `json:"updated_at" example:"1676546709" gorm:"not null"`
	Role      string `json:"role" example:"user" gorm:"not null"`
	Plan      string `json:"plan" example:"free" gorm:"not null"`

	Tasks []Task `json:"tasks" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Plan string

const (
	PlanFree    Plan = "free"
	PlanPremium Plan = "premium"
)

// -- Misc
type UserCredentials struct {
	Email    string `json:"email" example:"example@togo.dev"`
	Password string `json:"password" example:"my_super_secret_password"`
}
