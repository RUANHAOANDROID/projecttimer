package db

import "projecttimer/utils"

type Register struct {
	Pos      int64 `json:"pos"`       //pos数量
	Server   int64 `json:"server"`    //后台服务
	PosDroid int64 `json:"pos_droid"` //自助机
	Other    int64 `json:"other"`     //其他
}
type Customer struct {
	ID         uint   `json:"id"gorm:"primaryKey"` //序号
	Name       string `json:"name"`                //项目名称
	Brand      string `json:"brand"`               //品牌
	Version    string `json:"version"`             //版本
	Register          //注册数量
	Salesman   string `json:"salesman"`   //业务员
	Technician string `json:"technician"` //技术员
	//Phone      string `json:"phone"`               //电话
	//Address    string `json:"address"`             //地址
	UseTime    int64  `json:"use_time"`   //使用日期
	EndTime    int64  `json:"end_time"`   //截止日期
	Purchased  byte   `json:"purchased"`  //1已购买 0 WEIGOUMAI
	Remark     string `json:"remark"`     //备注
	Customize1 string `json:"customize1"` //定制列1
	Customize2 string `json:"customize2"` //定制列2
}

func (c Customer) Add() error {
	return DB.Create(&c).Error
}

func (c Customer) Update() error {
	return DB.Model(&c).Updates(c).Error
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
