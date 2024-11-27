package model

type User struct {
    Id        int    `gorm:"type:int;primaryKey"`
    Firstname string `gorm:"type:varchar(255);not null"`
    Lastname  string `gorm:"type:varchar(255);not null"`
    Email     string `gorm:"type:varchar(255);unique"`
    Password  string `gorm:"type:varchar(255);not null"`
	IsAdmin   bool   `gorm:"default:false" json:"isadmin"`
}

// TableName permet de définir explicitement le nom de la table.
func (User) TableName() string {
    return "user"
}