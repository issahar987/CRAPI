package main

type Config struct {
	TargetURL string `json:"target_url"`
	LoginURL  string `json:"login_url"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Post struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Author    Author        `json:"author"`
	Comments  []interface{} `json:"comments"`
	AuthorID  int           `json:"authorid"`
	CreatedAt string        `json:"CreatedAt"`
}

type Author struct {
	Nickname      string `json:"nickname"`
	Email         string `json:"email"`
	VehicleID     string `json:"vehicleid"`
	ProfilePicURL string `json:"profile_pic_url"`
	CreatedAt     string `json:"created_at"`
}
