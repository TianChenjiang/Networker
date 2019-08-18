package schema

//type UserRoles struct {
//	RoleId     string    `json:"role_id"`
//}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`

	//CreatedAt  time.Time `json:"created_at"`
}

type CreateUserForm struct {
	ID       uint   `form:"id" ` //valid:"Required;MaxSize(100)
	Username string `form:"username"`
	Password string `form:"password"`
	Role     string `form:"role"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
}












