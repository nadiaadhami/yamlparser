package models

import "time"

type States struct {
	ID           int  `orm:"column(id);pk;"`
	Name         string `orm:"column(name);size(255)"`
	Code         string `orm:"column(code);size(2)"`
	CountryAlpha string `orm:"column(country_alpha);size(2)"`

	InsertedAt time.Time `orm:"column(inserted_at);auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"column(updated_at);auto_now_add;type(datetime)"`
}

func (p *States) TableName() string {
	return "states"
}

type StateMap map[int]States
