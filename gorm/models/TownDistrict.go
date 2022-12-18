package models

import (
	"encoding/json"
	"time"
	// "gorm.io/gorm"
)

type TownDistrict struct {
	Id           string     `gorm:"column:ID;not null; primaryKey"`
	LocName      string     `gorm:"column:LOC_NAME"`
	Parent       string     `gorm:"column:PARENT"`
	CreatedBy    string     `gorm:"column:CREATED_BY"`
	CreatedDate  *time.Time `gorm:"column:CREATED_DATE"`
	ModifiedBy   string     `gorm:"column:MODIFIED_BY"`
	ModifiedDate *time.Time `gorm:"column:MODIFIED_DATE"`
	IsLastest    *int       `gorm:"column:IS_LASTEST"`
}

func ToJson(val []byte) TownDistrict {
	town := TownDistrict{}
	err := json.Unmarshal(val, &town)
	if err != nil {
		panic(err)
	}
	return town
}

func (TownDistrict) TableName() string { return "CL_MB_TOWN_DISTRICT" }
