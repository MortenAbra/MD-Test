package dbmodels

type RocketModel struct {
	Id      string `gorm:"primaryKey"`
	Name    string `gorm:"column:name;type:varchar;size:255;not null"`
	Mission string `gorm:"column:mission;type:varchar;size:255;not null"`
	Speed   int    `gorm:"column:speed;type:int;not null"`
}

func (RocketModel) TableName() string {
	return "rockets"
}
