package entities

import (
	"time"

	_models "github.com/MarkTBSS/082_Item_Creating/pkg/itemShop/models"
)

type Item struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	AdminID     *string   `gorm:"type:varchar(64);"`
	Name        string    `gorm:"type:varchar(64);unique;not null;"`
	Description string    `gorm:"type:varchar(128);not null;"`
	Picture     string    `gorm:"type:varchar(256);not null;"`
	Price       uint      `gorm:"not null;"`
	IsArchive   bool      `gorm:"not null;default:false;"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
}

func (i *Item) ChangeToItemModel() *_models.Item {
	return &_models.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
		Price:       i.Price,
	}
}
