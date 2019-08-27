package schema

//用户资源 VO
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type Users []User


type Auth struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordSchema struct {
	Password string  `json:"password"`
}



