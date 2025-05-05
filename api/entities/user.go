package entities

type (
	User struct {
		name     string
		age      int
		password string
	}
)

func (u *User) AuthUser() string {
	return u.name
}
func NewUser(name, password string, age int) *User {
	return &User{
		name:     name,
		age:      age,
		password: "123123",
	}
}
