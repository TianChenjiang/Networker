package schema

import (
	"time"
)


type UserRoles struct {
	RoleId     string    `json:"role_id"`
}

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string	 `json:"password"`
	Role       UserRoles `json:"role"`
	Phone      string    `json:"phone"`
	Email      string	 `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}











