package model

import (
    "github.com/jinzhu/gorm"
)

type Property struct {
    Name     string `sql:"type varchar(200);not null;"`
    gorm.Model
}

type Room struct {
	Property   Property
	PropertyID uint   `sql:"not null;unique"`
    gorm.Model
}

func GetAllProperties() []Property {
    var res []Property
    if err := _db.Find(&res).Error; err != nil {
        panic(err)
    }
    return res
}
