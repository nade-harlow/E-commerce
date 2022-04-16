package requests

type UserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Country     string `json:"country"`
	DateOfBirth string `json:"dateofbirth"`
	Avatar      string `json:"avatar"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `json:"password,omitempty"`
}
