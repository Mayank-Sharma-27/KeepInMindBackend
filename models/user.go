package models

import (
	"time"
) 

type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
	
}