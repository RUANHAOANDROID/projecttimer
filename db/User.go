package db

type User struct {
	ID     uint   `json:"id"gorm:"primaryKey"`
	Number string `json:"number"`
	Pwd    string `json:"pwd"`
}

func (u User) Add() error {
	return DB.Create(&u).Error
}

func (u User) Update() error {
	return DB.Model(&u).Where("id=?", u.ID).Updates(u).Error
}
