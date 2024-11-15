package models

type Profile struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type User struct {
	ID           string  `json:"id"`
	UserName     string  `json:"user_name"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"password_hash"`
	Profile      Profile `json:"profile"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"deleted_at"`
}

type LoginRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type ListRequest struct{}

type ListResponse struct {
	List []User `json:"list"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Response struct {
	Token  Tokens `json:"token"`
	Status string `json:"status"`
}

type UniverResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
