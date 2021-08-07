package model

import (
	"fmt"
	"rdb/driver"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
	AddTime string `json:"add_time"`
}

func NewUser() *User{
	return &User{}
}

func (*User) TableName() string {
	return "user"
}

func (u *User) LoadById(id int) *User {
	fmt.Println(id)
	driver.DB.First(u,id)
	return u
}


