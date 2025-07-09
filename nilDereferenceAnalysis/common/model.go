package common

import "errors"

type User struct {
	Age  int
	Name string
}

func NewUser(name string, age int) *User {
	user, err := CreateUser(name, age)
	if err != nil {
		return nil
	}
	return user
}

func CreateUser(name string, age int) (*User, error) {
	if !isValidName(name) {
		return nil, errors.New("invalid name")
	}
	if !isValidAge(age) {
		return nil, nil
	}
	return &User{Age: age, Name: name}, nil
}

func (u *User) Copy(ctx *Context) *User {
	if ctx.isDebugEnabled {
		logUserEvent("copy user", u)
	}
	if u == nil {
		return nil
	}
	return &User{Age: u.Age, Name: u.Name}
}

func isValidName(name string) bool {
	return len(name) > 0
}

func isValidAge(age int) bool {
	return age >= 0
}
