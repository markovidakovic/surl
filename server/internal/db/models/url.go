package models

import "time"

type Url struct {
	ID        uint `gorm:primaryKey`
	CreatedAt time.Time
	UpdatedAt time.Time
	Original  string `gorm:not null`
	Short     string `gorm:unique;not null`
}
