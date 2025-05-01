package entities

type (
	User struct {
		name string
		age  int
	}
)

func (u *User) AuthUser() string {
	return u.name
}
func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}
