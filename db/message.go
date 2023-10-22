package db

type Message struct {
	ID      uint   `json:"id"gorm:"primaryKey"`
	Type    string `json:"name"`
	Content string `json:"phone"`
	Unread  byte   `json:"read_status"` //1 unread 0 read
}

func (c Message) Add() error {
	return DB.Create(&c).Error
}

func (c Message) Update() error {
	return DB.Model(&c).Where("id=?", c.ID).Updates(c).Error
}
