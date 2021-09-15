package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// Get dummy user
func LoadTestUser() *User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &User{Password: string(hashedPassword), Email: "test@email.com"}
}
