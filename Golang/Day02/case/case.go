package _case

import "errors"

var g_val int
var G_val int

func SumCase(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("a b can not be zero")
		return 0, err
	}
	sum = a + b
	return sum, nil
}

func ReferenceCase(a *int, b *int) {
	*a += 1
	*b += 3
}

func ScopeCase(a, b int) {
	c := 100
	g_val = a + b + c
	G_val = g_val
}

type User struct {
	Name string
	Age  uint
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() uint {
	return u.Age
}

func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}
