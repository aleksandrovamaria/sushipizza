package models

import "time"

// District is the district model
type District struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Category is the category of food model
type Category struct {
	ID        int
	Name      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Item is the item model
type Item struct {
	ID          int
	Category    Category
	Name        string
	Weight      string
	Description string
	Price       int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
