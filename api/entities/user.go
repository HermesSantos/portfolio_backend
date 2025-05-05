package entities

type (
	User struct {
		Email     string `json:"email" validate:"required"`
		Age      int    `json:"age"`
		Password string `json:"password" validate:"required"`
	}
)

func (u *User) AuthUser() string {
	return u.Email
}
