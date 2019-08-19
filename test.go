package main

import (
	"fmt"
	"go.uber.org/dig"
)

//======controller==========
type c_user struct {
	BlUser bl_user
}

//controller 构造器
func newCUser(user bl_user) *c_user {
	return &c_user{
		BlUser: user,
	}
}

//======bl=================
type bl_user struct {
	MUser m_user
}

// bl 构造器
func NewBlUser(user m_user) *bl_user {
	return &bl_user{
		MUser: user,
	}
}

//=======model=========
type m_user struct {
	name string
}

func NewMUser() *m_user {
	return &m_user{name: "nameaaa"}
}

func provide(container *dig.Container) {
	container.Provide(NewBlUser)
	container.Provide(NewMUser)
	container.Provide(newCUser)
	fmt.Print(container)
	err := container.Provide(func(user *c_user) error {
		fmt.Printf(user.BlUser.MUser.name)
		return nil
	})
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	fmt.Print("hello")
	contrainer := dig.New()
	provide(contrainer)
}
