package model

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