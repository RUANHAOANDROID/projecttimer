package db

import "projecttimer/utils"

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
func (c Customer) ListPage(
	count *int64, customers *[]Customer,
	offset int, limit int,
) error {
	customerTab := DB.Model(&Customer{})
	err := customerTab.Count(count).Error
	if err != nil {
		utils.Log.Error(err.Error())
	}
	//err = eventTab.Where("id BETWEEN ? AND ?", offset, limit).Find(&events).Error
	err = customerTab.Order("id DESC").Offset(offset).Limit(limit).Find(&customers).Error //查询pageindex页的数据
	return err
}
