package models

import "time"

type Todo struct {
	ID        uint  		`json:"id"`
	UserID    uint			`json:"user_id"`
	Title     string    	`json:"title"`
	Description string   	`json:"description"`
	Completed bool      	`json:"completed"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}