package single_responsibility

import "time"

type Gorm struct{}

type GormModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (g *Gorm) Create(t interface{}) error {
	return nil
}
