package model

import "time"

type User struct {
    Id        int    `gorm:"type:int;primaryKey"`
    Firstname string `gorm:"type:varchar(255);not null"`
    Lastname  string `gorm:"type:varchar(255);not null"`
    Email     string `gorm:"type:varchar(255);unique"`
    Password  string `gorm:"type:varchar(255);not null"`
	IsAdmin   bool   `gorm:"default:false" json:"isadmin"`
}


func (User) TableName() string {
    return "user"
}

type Book struct {
    Id        int    `gorm:"type:int;primaryKey"`
    Title string `gorm:"type:varchar(255);not null"`
    Author  string `gorm:"type:varchar(255);not null"`
    Release *time.Time  `gorm:"type:time"`
    Resume  string `gorm:"type:varchar(255);not null"`
	Stock   bool   `gorm:"type:bool"`
    Price   float64 
}


func (Book) TableName() string {
    return "book"
}