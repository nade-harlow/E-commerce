package requests

type UserSignUpRequest struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name"`
	Role            string `json:"role,omitempty"`
	Email           string `gorm:"unique" json:"email" validate:"required,email"`
	Password        string `json:"password,omitempty" validate:"required,min=8,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password,omitempty" gorm:"-" validate:"required,min=8,eqfield=Password"`
	Telephone       string `json:"telephone" validate:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAddressRequest struct {
	UserID       string `json:"user_id"`
	AddressLine1 string `json:"address_line_1" validate:"required"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	Country      string `json:"country" validate:"required"`
	Mobile       string `json:"mobile"`
}
