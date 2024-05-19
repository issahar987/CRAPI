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

type VideoResponse struct {
	ID               int    `json:"id"`
	VideoName        string `json:"video_name"`
	ConversionParams string `json:"conversion_params"`
	ProfileVideo     string `json:"profileVideo"`
}
