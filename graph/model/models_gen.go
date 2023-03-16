package model

type Link struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	User    *User  `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewLink struct {
	Title   string `json:"title"`
	Address string `json:"address"`
}

type NewUser struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID   string `json:"id"`
	Fullname string `json:"fullname"`
	Name string `json:"name"`
	Status string `json:"status"`
}
