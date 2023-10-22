package db

type Customer struct {
	ID      uint   `json:"id"gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	UseTime int64  `json:"use_time"`
	EndTime int64  `json:"end_time"`
}

func (c Customer) Add() error {
	return DB.Create(&c).Error
}

func (c Customer) Update() error {
	return DB.Model(&c).Where("id=?", c.ID).Updates(c).Error
}
func (c Customer) List(cs *[]Customer) error {
	return DB.Order("end_time").Find(cs).Error
}

func (c Customer) Delete(id string) error {
	return DB.Delete(&c, "id=?", id).Error
}
