package main

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Value     string `json:"value"`
	ExpiresIn int    `json:"expiresIn"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}