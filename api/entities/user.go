package entities

type (
	User struct {
		Email     string
		Age      int
		Password string
	}
	UserDTO struct {
		Email     string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func (u *User) AuthUser() string {
	return u.Email
}
